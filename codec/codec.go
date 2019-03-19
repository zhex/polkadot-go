package codec

type Codec interface {
	EncodedLength() int
	Hex() string
	U8a() *U8a
	String() string
	Equal(b Codec) bool
	Empty() bool
}
