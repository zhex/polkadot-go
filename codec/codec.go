package codec

type Codec interface {
	EncodedLength() int
	Hex() string
	Bytes() []byte
	String() string
	Equal(b Codec) bool
	Empty() bool
}

type StructBuilder interface {
	BuildStruct() Codec
	Decode([]byte) Codec
	Encode(Codec) []byte
}
