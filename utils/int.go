package utils

import (
	"encoding/binary"
)

func IntToByte(d int64) []byte {
	var b = make([]byte, binary.MaxVarintLen64)
	l := binary.PutVarint(b, d)
	return b[:l]
}

func UintToByte(d uint64) []byte {
	l := binary.MaxVarintLen64
	var b = make([]byte, l)
	binary.LittleEndian.PutUint64(b, d)
	for i := l - 1; i > 0; i-- {
		if b[i] == 0 {
			l = i
		}
	}
	return b[:l]
}
