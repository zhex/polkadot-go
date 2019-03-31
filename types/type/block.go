package _type

import (
	"github.com/zhex/polkadot-go/types/primitives"
)

type Block struct {
	Header     *BlockHeader
	Extrinsics []interface{}
}

type BlockHeader struct {
	ParentHash     primitives.Hash256
	Number         uint64
	StateRoot      primitives.Hash256
	ExtrinsicsRoot primitives.Hash256
	Digest         Digest
}

type SignedBlock struct {
	Block         *Block
	Justification interface{}
}
