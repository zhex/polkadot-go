package metadata

import "github.com/zhex/polkadot-go/codec"

type Metadata struct {
	MagicNumber uint32
	Metadata    codec.EnumType
}

type PlainType struct {
	Value string
}

type MapType struct {
	Key      string
	Value    string
	IsLinked bool
}

type DoubleMapType struct {
	Key1      string
	Key2      string
	Value     string
	KeyHasher string
}
