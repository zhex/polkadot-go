package client

import (
	"github.com/zhex/polkadot-go/jsonrpc"
	"github.com/zhex/polkadot-go/types/type"
)

func createAuthor(p *jsonrpc.WsProvider) *author {
	a := author{}
	a.provider = p
	a.section = "author"
	return &a
}

type author struct {
	rpcBase
}

func (a *author) PendingExtrinsics() {
	// todo
}

func (a *author) SubmitAndWatchExtrinsic(ex _type.Extrinsic) {
	// todo
}

func (a *author) SubmitExtrinsic(ex _type.Extrinsic) {
	// todo
}
