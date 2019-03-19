package main

import (
	"fmt"
	"strconv"

	"github.com/zhex/polkadot-go/client"
)

const MagicNumber = 0x6174656d

func IntToBytes(n int) []byte {
	return []byte(strconv.Itoa(n))
}

func main() {
	c, _ := client.New("wss://poc3-rpc.polkadot.io/")
	name, err := c.RPC.System.Peers()
	fmt.Println(name, err)
	// params := make([]interface{}, 0)
	// resp, err := p.Call("state_getMetadata", params)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data, err := utils.HexToBytes(resp.Result.(string))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // result := append(IntToBytes(MagicNumber), 0)
	// // result = append(result, data...)
	// fmt.Println(string(data))
}
