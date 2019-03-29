package v3

import (
	"github.com/zhex/polkadot-go/codec"
	"github.com/zhex/polkadot-go/metadata"
)

type Metadata struct {
	Modules []Module
}

type Module struct {
	Name    string
	Prefix  string
	Storage codec.Option
	Calls   codec.Option
	Events  codec.Option
}

type Storage struct {
	Name     string
	Modifier *codec.EnumType
	Type     *codec.EnumType
	Fallback []byte
	Docs     []string
}

func CreateType() *codec.EnumType {
	return &codec.EnumType{
		Def: []interface{}{
			metadata.PlainType{},
			metadata.MapType{},
			metadata.DoubleMapType{},
		},
	}
}

func CreateModifier() *codec.EnumType {
	return &codec.EnumType{
		Def: []interface{}{
			nil,
			nil,
		},
	}
}
