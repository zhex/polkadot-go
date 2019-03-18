package main

import (
	"fmt"
	"log"

	"github.com/zhex/polkadot-go/jsonrpc"
)

func main() {
	p := jsonrpc.NewWsProvider("wss://poc3-rpc.polkadot.io/")
	params := make([]interface{}, 0)
	resp, err := p.Call("state_getMetadata", params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Result)
}
