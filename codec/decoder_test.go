package codec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeStruct(t *testing.T) {
	s := struct {
		Foo  string
		Bar  uint32
		Book *struct {
			Name string
		}
	}{}

	DecodeStruct(&s, []byte{28, 98, 97, 122, 122, 105, 110, 103, 69, 0, 0, 0})
	assert.Equal(t, "bazzing", s.Foo)
	assert.Equal(t, 69, int(s.Bar))
}
