package client

import "github.com/zhex/polkadot-go/jsonrpc"

func createChain(p *jsonrpc.WsProvider) *chain {
	c := chain{}
	c.provider = p
	c.section = "chain"
	return &c
}

type chain struct {
	rpcBase
}
