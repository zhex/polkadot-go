package client

import (
	"github.com/zhex/polkadot-go/jsonrpc"
	"github.com/zhex/polkadot-go/types/primitives"
	"github.com/zhex/polkadot-go/types/type"
	"github.com/zhex/polkadot-go/utils/decoder"
)

func createChain(p *jsonrpc.WsProvider) *chain {
	c := chain{}
	c.provider = p
	c.section = "chain"
	return &c
}

type chain struct {
	rpcBase
}

func (c *chain) GetBlock(hash *primitives.Hash256) (*_type.SignedBlock, error) {
	var params []interface{}
	if hash == nil {
		params = emptyParams
	} else {
		params = []interface{}{hash.String()}
	}
	result, err := c.call("getBlock", params)
	if err != nil {
		return nil, err
	}
	signedBlock := _type.SignedBlock{}
	err = decoder.MapDecode(result, &signedBlock)
	return &signedBlock, err
}
