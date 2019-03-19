package codec

import (
	"bytes"
	"encoding/hex"

	"github.com/zhex/polkadot-go/utils"
)

type U8a struct {
	value []byte
}

func (u *U8a) Value() []byte {
	return u.value
}

func (u *U8a) EncodedLength() int {
	return len(u.value)
}

func (u *U8a) Hex() string {
	return u.Hex()
}

func (u *U8a) String() string {
	return hex.EncodeToString(u.value)
}

func (u *U8a) Equal(b interface{}) bool {
	var d *U8a
	switch v := b.(type) {
	case U8a:
		d = &v
	default:
		d = NewU8a(b)
	}
	return bytes.Equal(u.value, d.Value())
}

func (u *U8a) Empty() bool {
	return len(u.value) == 0
}

func NewU8a(input interface{}) *U8a {
	var value []byte
	switch v := input.(type) {
	case []byte:
		value = v
	case string:
		if utils.IsHex(v) {
			value, _ = utils.HexToBytes(v)
		} else {
			value = []byte(v)
		}
	case bytes.Buffer:
		value = v.Bytes()
	}
	return &U8a{
		value: value,
	}
}
