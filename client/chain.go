package client

import (
	"fmt"
	"github.com/zhex/polkadot-go/jsonrpc"
	"github.com/zhex/polkadot-go/types/primitives"
	"github.com/zhex/polkadot-go/types/rpccall"
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
	var params []interface{} = nil
	if hash != nil {
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

func (c *chain) GetBlockHash(n *uint64) (*primitives.Hash256, error) {
	var params []interface{} = nil
	if n != nil {
		params = []interface{}{*n}
	}
	result, err := c.call("getBlockHash", params)
	if err != nil {
		return nil, err
	}
	return primitives.NewHash256(result.(string)), nil
}

func (c *chain) GetFinalisedHead() (*primitives.Hash256, error) {
	// todo
	return nil, nil
}

func (c *chain) GetHeader(hash *primitives.Hash256) (*_type.BlockHeader, error) {
	var params []interface{} = nil
	if hash != nil {
		params = []interface{}{hash.String()}
	}
	result, err := c.call("getHeader", params)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	var header _type.BlockHeader
	err = decoder.MapDecode(result, &header)
	return &header, err
}

func (c *chain) GetRuntimeVersion() (*rpccall.RuntimeVersion, error) {
	result, err := c.call("getRuntimeVersion", nil)
	if err != nil {
		return nil, err
	}
	var v rpccall.RuntimeVersion
	err = decoder.MapDecode(result, &v)
	return &v, nil
}

func (c *chain) SubscribeFinalisedHeads() {
	// todo
}

func (c *chain) SubscribeNewHead() {
	// todo
}

func (c *chain) SubscribeRuntimeVersion() {
	// todo
}
