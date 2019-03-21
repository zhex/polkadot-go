package utils

import (
	"encoding/binary"
	"encoding/hex"
)

func DecodeBytes(data []byte) (uint64, uint64) {
	flag := data[0] & 0x03
	if flag == 0x00 {
		return 1, uint64(data[0] >> 2)
	} else if flag == 0x01 {
		d := FixByteWidth(data[:2], 8, true)
		return 2, binary.LittleEndian.Uint64(d) >> 2
	} else if flag == 0x02 {
		d := FixByteWidth(data[:4], 8, true)
		return 4, binary.LittleEndian.Uint64(d) >> 2
	}
	l := uint64(data[0]>>2) + 4
	offset := l + 1
	d := FixByteWidth(data[1:offset], 8, true)
	return offset, binary.LittleEndian.Uint64(d)
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

func BytesToHex(data []byte) string {
	s := hex.EncodeToString(data)
	return HexAddPrefix(s)
}
