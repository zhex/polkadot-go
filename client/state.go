package client

import (
	"github.com/zhex/polkadot-go/jsonrpc"
	"github.com/zhex/polkadot-go/types/rpccall"
	"github.com/zhex/polkadot-go/utils"
	"github.com/zhex/polkadot-go/utils/decoder"
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

func (s *state) GetRuntimeVersion() (*rpccall.RuntimeVersion, error) {
	result, err := s.call("getRuntimeVersion", emptyParams)
	if err != nil {
		return nil, err
	}
	var v rpccall.RuntimeVersion
	err = decoder.MapDecode(result, &v)
	return &v, nil
}

func (s *state) GetStorage() (interface{}, error) {
	// todo
	return nil, nil
}

func (s *state) GetStorageHash() (interface{}, error) {
	// todo
	return nil, nil
}

func (s *state) GetStorageSize() (interface{}, error) {
	// todo
	return nil, nil
}

func (s *state) QueryStorage() (interface{}, error) {
	// todo
	return nil, nil
}

func (s *state) SubscribeStorage() (interface{}, error) {
	// todo
	return nil, nil
}
