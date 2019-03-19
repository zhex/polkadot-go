package client

import "github.com/zhex/polkadot-go/jsonrpc"

// examplge:
// chain := c.RPC.System.Chain()
// name := c.RPC.System.Name()

type Client struct {
	provider *jsonrpc.WsProvider
	RPC      *rpc
}

func New(url string) *Client {
	p := jsonrpc.NewWsProvider(url)
	return &Client{
		provider: p,
		RPC: &rpc{
			System: createSystem(p),
			Author: createAuthor(p),
			State:  createState(p),
			Chain:  createChain(p),
		},
	}
}
