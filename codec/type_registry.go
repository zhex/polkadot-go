package codec

type TypeEncoder = func(interface{}) []byte
type TypeDecoder = func([]byte) interface{}

var encodeMap = map[string]TypeEncoder{}
var decodeMap = map[string]TypeDecoder{}

func RegisterType(name string, encoder TypeEncoder, decoder TypeDecoder) {
	encodeMap[name] = encoder
	decodeMap[name] = decoder
}
