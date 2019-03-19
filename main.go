package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/zhex/polkadot-go/utils"

	"github.com/zhex/polkadot-go/jsonrpc"
)

const MagicNumber = 0x6174656d

func IntToBytes(n int) []byte {
	return []byte(strconv.Itoa(n))
}

func main() {
	p := jsonrpc.NewWsProvider("wss://poc3-rpc.polkadot.io/")
	params := make([]interface{}, 0)
	resp, err := p.Call("state_getMetadata", params)
	if err != nil {
		log.Fatal(err)
	}
	data, err := utils.HexToBytes(resp.Result.(string))
	if err != nil {
		log.Fatal(err)
	}
	// result := append(IntToBytes(MagicNumber), 0)
	// result = append(result, data...)
	fmt.Println(string(data))
}
