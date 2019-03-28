package client

import "github.com/zhex/polkadot-go/jsonrpc"

var emptyParams []interface{}

type rpc struct {
	System *system
	State  *state
	Chain  *chain
	Author *author
}

type rpcBase struct {
	provider *jsonrpc.WsProvider
	section  string
}

func (r *rpcBase) call(method string, params []interface{}) (interface{}, error) {
	resp, err := r.provider.Call(r.section+"_"+method, params)
	return resp.Result, err
}
