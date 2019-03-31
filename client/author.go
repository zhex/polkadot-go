package client

import (
	"github.com/zhex/polkadot-go/codec"
	"github.com/zhex/polkadot-go/jsonrpc"
	"github.com/zhex/polkadot-go/types/type"
	"github.com/zhex/polkadot-go/utils"
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

func (a *author) PendingExtrinsics() ([]_type.Extrinsic, error) {
	result, err := a.call("pendingExtrinsics", nil)
	if err != nil {
		return nil, err
	}
	var extrinsics []_type.Extrinsic
	for _, item := range result.([]string) {
		b, _ := utils.HexToBytes(item)
		ex := _type.Extrinsic{}
		_, err := codec.Decode(b, &ex)
		if err != nil {
			return nil, err
		}
		extrinsics = append(extrinsics, ex)
	}
	return extrinsics, nil
}

func (a *author) SubmitAndWatchExtrinsic(ex _type.Extrinsic) {
	// todo
}

func (a *author) SubmitExtrinsic(ex _type.Extrinsic) {
	// todo
}
