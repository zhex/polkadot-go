package client

import (
	"fmt"
	"github.com/zhex/polkadot-go/jsonrpc"
	"github.com/zhex/polkadot-go/utils"
)

func createState(p *jsonrpc.WsProvider) *state {
	s := state{}
	s.provider = p
	s.section = "state"
	return &s
}

type state struct {
	rpcBase
}

func (s *state) GetMetadata() (interface{}, error) {
	result, err := s.call("getMetadata", emptyParams)
	if err != nil {
		return nil, err
	}
	data, err := utils.HexToBytes(result.(string))
	if err != nil {
		return nil, err
	}
	//1635018093
	data = data[4:] // strip magic number
	data = data[1:] // strip enum idx
	offset, l := utils.DecodeBytes(data)
	fmt.Println(offset)
	data = data[offset:] // strip vector prefix

	offset, l = utils.DecodeBytes(data)
	fmt.Println(offset, l, string(data[offset:offset+l]))
	data = data[l+offset:] // strip m[0].name

	offset, l = utils.DecodeBytes(data)
	fmt.Println(offset, l, string(data[offset:offset+l]))

	//offset, length := utils.DecodeBytes(data)
	return string(data), nil
}
