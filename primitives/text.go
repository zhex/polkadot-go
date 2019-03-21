package primitives

import (
	"encoding/hex"

	"github.com/zhex/polkadot-go/utils"

	"github.com/zhex/polkadot-go/codec"
)

type Text string

func (t Text) EncodedLength() int {
	return len(t.Bytes())
}

func (t Text) Bytes() []byte {
	return []byte(t)
}

func (t Text) Hex() string {
	h := hex.EncodeToString(t.Bytes())
	return utils.HexAddPrefix(h)
}

func (t Text) String() string {
	return string(t)
}

func (t Text) Empty() bool {
	return len(t) == 0
}

func (t Text) Equal(c codec.Codec) bool {
	return t.String() == c.String()
}
