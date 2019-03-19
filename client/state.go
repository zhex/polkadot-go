package client

import (
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
	return string(data), nil
}
