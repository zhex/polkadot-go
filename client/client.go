package client

import "github.com/zhex/polkadot-go/jsonrpc"

type Client struct {
	provider *jsonrpc.WsProvider
	RPC      *rpc
}

func New(url string) (*Client, error) {
	p, err := jsonrpc.NewWsProvider(url)
	if err != nil {
		return nil, err
	}
	c := &Client{
		provider: p,
		RPC: &rpc{
			System: createSystem(p),
			Author: createAuthor(p),
			State:  createState(p),
			Chain:  createChain(p),
		},
	}
	return c, nil
}
