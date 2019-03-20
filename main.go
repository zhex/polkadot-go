package main

import (
	"fmt"
	"strconv"

	"github.com/zhex/polkadot-go/utils"
)

const MagicNumber = 0x6174656d

func IntToBytes(n int) []byte {
	return []byte(strconv.Itoa(n))
}

func main() {
	// c, _ := client.New("wss://poc3-rpc.polkadot.io/")
	// name, err := c.RPC.System.Properties()
	// fmt.Println(name, err)
	offset, l := utils.DecodeBytes([]byte{253, 7})
	fmt.Println(offset, l)
	// b := utils.FixByteWidth([]byte{1, 2, 3}, 8)
	// fmt.Println(b)
}
