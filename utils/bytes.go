package utils

import (
	"encoding/hex"
)

func AddBytePrefix(data []byte, l int) []byte {
	b := UintToByte(uint64(l))
	return append(b, data...)
}

func FixByteWidth(b []byte, w int, suffix bool) []byte {
	l := len(b)
	if l >= w {
		return b
	}
	bb := make([]byte, w)
	if suffix {
		for i := 0; i < l; i++ {
			bb[i] = b[i]
		}
	} else {
		for i := w - 1; i >= w-l; i-- {
			bb[i] = b[l-w+i]
		}
	}
	return bb
}

func ByteToHex(data []byte) string {
	s := hex.EncodeToString(data)
	return HexAddPrefix(s)
}

func BoolToByte(data bool) []byte {
	if data {
		return []byte{1}
	}
	return []byte{0}
}
