package client

import (
	"github.com/mitchellh/mapstructure"
	"github.com/zhex/polkadot-go/jsonrpc"
	"github.com/zhex/polkadot-go/types"
)

func createSystem(p *jsonrpc.WsProvider) *system {
	s := system{}
	s.provider = p
	s.section = "system"
	return &s
}

type system struct {
	rpcBase
}

func (s *system) Name() (string, error) {
	result, err := s.call("name", emptyParams)
	return result.(string), err
}

func (s *system) Version() (string, error) {
	result, err := s.call("version", emptyParams)
	return result.(string), err
}

func (s *system) Chain() (string, error) {
	result, err := s.call("chain", emptyParams)
	return result.(string), err
}

func (s *system) Health() (*types.Health, error) {
	result, err := s.call("health", emptyParams)
	if err != nil {
		return nil, err
	}
	health := types.Health{}
	err = mapstructure.Decode(result, &health)
	return &health, err
}

func (s *system) Peers() ([]types.PeerInfo, error) {
	result, err := s.call("peers", emptyParams)
	var peers []types.PeerInfo
	err = mapstructure.Decode(result, &peers)
	return peers, err
}

func (s *system) NetworkState() (interface{}, error) {
	result, err := s.call("networkState", emptyParams)
	return result, err
}

func (s *system) Properties() (*types.ChainProperties, error) {
	result, err := s.call("properties", emptyParams)
	if err != nil {
		return nil, err
	}
	properties := types.ChainProperties{}
	err = mapstructure.Decode(result, &properties)
	return &properties, err
}
