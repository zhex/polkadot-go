package codec

import (
	"github.com/zhex/polkadot-go/types/type"
	"reflect"
)

type TypeEncoder = func(interface{}) ([]byte, error)
type TypeDecoder = func([]byte, reflect.Value) (interface{}, error)

var encodeMap = map[string]TypeEncoder{}
var decodeMap = map[string]TypeDecoder{}

func RegisterType(name string, encoder TypeEncoder, decoder TypeDecoder) {
	encodeMap[name] = encoder
	decodeMap[name] = decoder
}

func init() {
	RegisterType("EnumType", EncodeEnumType, DecodeEnumType)
	RegisterType("Option", EncodeOption, DecodeOption)
	RegisterType("Extrinsic", _type.EncodeExtrinsic, _type.DecodeExtrinsic)
}
