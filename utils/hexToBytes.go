package utils

import (
	"encoding/hex"
)

func HexToBytes(data string) ([]byte, error) {
	data = StripPrefix(data)
	return hex.DecodeString(data)
}

func StripPrefix(data string) string {
	if data[:2] == "0x" {
		return data[2:]
	}
	return data
}
