package client

import "github.com/zhex/polkadot-go/jsonrpc"

func createState(p *jsonrpc.WsProvider) *state {
	s := state{}
	s.provider = p
	s.section = "state"
	return &s
}

type state struct {
	rpcCall
}

func (s *state) GetMetadata() (interface{}, error) {
	return s.call("getMetadata", emptyParams)
}
