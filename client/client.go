package client

import (
	"github.com/zhex/polkadot-go/codec"
	"github.com/zhex/polkadot-go/jsonrpc"
	"github.com/zhex/polkadot-go/types/type"
)

func init() {
	codec.RegisterType("Extrinsic", _type.EncodeExtrinsic, _type.DecodeExtrinsic)
}

type Client struct {
	provider *jsonrpc.WsProvider
	RPC      *rpc
}

// New client
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
