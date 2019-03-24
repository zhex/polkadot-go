package utils

import (
	"encoding/binary"
	"encoding/hex"
	"math"
)

func AddBytePrefix(data []byte, l int) []byte {
	b := IntToByte(l)
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

func BytesToHex(data []byte) string {
	s := hex.EncodeToString(data)
	return HexAddPrefix(s)
}

func Uint8ToByte(data uint8) []byte {
	return []byte{data}
}

func Uint16ToByte(data uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, data)
	return b
}

func Uint32ToByte(data uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, data)
	return b
}

func Uint64ToByte(data uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, data)
	return b
}

func IntToByte(d int) []byte {
	if d <= math.MaxUint8 {
		return Uint8ToByte(uint8(d))
	} else if d <= math.MaxUint16 {
		return Uint16ToByte(uint16(d))
	} else if d <= math.MaxUint32 {
		return Uint32ToByte(uint32(d))
	}
	return Uint64ToByte(uint64(d))
}

func BooltoByte(data bool) []byte {
	if data {
		return []byte{1}
	}
	return []byte{0}
}
