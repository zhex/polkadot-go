package client

import "github.com/zhex/polkadot-go/jsonrpc"

func createAuthor(p *jsonrpc.WsProvider) *author {
	a := author{}
	a.provider = p
	a.section = "author"
	return &a
}

type author struct {
	rpcCall
}
