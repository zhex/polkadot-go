package jsonrpc

import (
	"context"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func NewWsProvider(endpoint string) *WsProvider {
	client, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	p := &WsProvider{
		client:     client,
		endpoint:   endpoint,
		id:         0,
		msgChan:    make(map[int](chan *Response)),
		subscribes: make(map[int][]func(*Response)),
		ctx:        ctx,
		ctxCancel:  cancel,
	}

	go func() {
		for {
			select {
			case <-p.ctx.Done():
				return
			default:
				resp := &Response{}
				err := client.ReadJSON(resp)
				if err != nil {
					fmt.Println(err)
					continue
				}

				if resp.Params.Result == nil {
					if c, ok := p.msgChan[resp.ID]; ok {
						c <- resp
					}
				} else {
					id := resp.Params.Subscription
					if callbacks, ok := p.subscribes[id]; ok {
						for _, callback := range callbacks {
							callback(resp)
						}
					}
				}
			}
		}
	}()

	return p
}

type WsProvider struct {
	client     *websocket.Conn
	endpoint   string
	id         int
	msgChan    map[int](chan *Response)
	subscribes map[int][]func(*Response)
	ctx        context.Context
	ctxCancel  context.CancelFunc
}

func (p *WsProvider) Call(method string, params []interface{}) (*Response, error) {
	p.id++
	req := &Request{
		ID:      p.id,
		Method:  method,
		JSONRPC: "2.0",
		Params:  params,
	}
	p.msgChan[p.id] = make(chan *Response, 1)
	err := p.client.WriteJSON(req)
	if err != nil {
		return nil, err
	}

	resp := <-p.msgChan[p.id]
	delete(p.msgChan, p.id)
	return resp, nil
}

func (p *WsProvider) Subscribe(method string, params []interface{}, callback func(*Response)) error {
	resp, err := p.Call(method, params)
	if err != nil {
		return err
	}
	id := int(resp.Result.(float64))
	if _, ok := p.subscribes[id]; ok {
		p.subscribes[id] = append(p.subscribes[id], callback)
	} else {
		p.subscribes[id] = []func(*Response){callback}
	}
	return nil
}

func (p *WsProvider) Close() {
	p.ctxCancel()
	for key := range p.msgChan {
		delete(p.msgChan, key)
	}
	p.client.Close()
}
