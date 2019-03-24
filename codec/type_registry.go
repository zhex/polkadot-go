package codec

import "reflect"

type TypeEncoder = func(interface{}) []byte
type TypeDecoder = func([]byte, reflect.Value) (interface{}, error)

var encodeMap = map[string]TypeEncoder{}
var decodeMap = map[string]TypeDecoder{}

func RegisterType(name string, encoder TypeEncoder, decoder TypeDecoder) {
	encodeMap[name] = encoder
	decodeMap[name] = decoder
}

func init() {
	RegisterType("EnumType", EncodeEnumType, DecodeEnumType)
}
