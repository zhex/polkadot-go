package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeBytes(t *testing.T) {
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
		offset, l := DecodeBytes(d.input)
		assert.Equal(t, d.expect, []uint64{offset, l})
	}
}
