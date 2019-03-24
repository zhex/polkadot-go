package codec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	//s := struct {
	//	Foo  string
	//	Bar  uint32
	//}{}
	//
	//_ := Decode(&s, []byte{28, 98, 97, 122, 122, 105, 110, 103, 69, 0, 0, 0})
	//assert.Equal(t, "bazzing", s.Foo)
	//assert.Equal(t, 69, int(s.Bar))
}

func TestGetBytesInfo(t *testing.T) {
	data := []struct {
		input  []byte
		expect []uint64
	}{
		{[]byte{252}, []uint64{1, 63}},
		{[]byte{253, 7}, []uint64{2, 511}},
		{[]byte{254, 255, 3, 0}, []uint64{4, 0xffff}},
		{[]byte{3, 249, 255, 255, 255}, []uint64{5, 0xfffffff9}},
	}

	for _, d := range data {
		offset, l := GetBytesInfo(d.input)
		assert.Equal(t, d.expect, []uint64{offset, l})
	}
}
