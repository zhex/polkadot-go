package main

import (
	"fmt"
	"log"

	"github.com/zhex/polkadot-go/jsonrpc"
)

func main() {
	p := jsonrpc.NewWsProvider("wss://poc3-rpc.polkadot.io/")
	params := make([]interface{}, 0)
	err := p.Subscribe("chain_subscribeNewHead", params, func(r *jsonrpc.Response) {
		fmt.Println(r.Params.Result)
	})
	if err != nil {
		log.Fatal(err)
	}
	select {}
}
