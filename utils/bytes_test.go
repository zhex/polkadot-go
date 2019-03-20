package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeBytes(t *testing.T) {
	offset, l := DecodeBytes([]byte{252})
	assert.Equal(t, uint64(1), offset)
	assert.Equal(t, uint64(63), l)

	offset, l = DecodeBytes([]byte{253, 7})
	assert.Equal(t, uint64(2), offset)
	assert.Equal(t, uint64(511), l)

	offset, l = DecodeBytes([]byte{254, 255, 3, 0})
	assert.Equal(t, uint64(4), offset)
	assert.Equal(t, uint64(0xffff), l)

	offset, l = DecodeBytes([]byte{3, 249, 255, 255, 255})
	assert.Equal(t, uint64(5), offset)
	assert.Equal(t, uint64(0xfffffff9), l)
}
