package codec

type Codec interface {
	EncodedLength() int
	Hex() string
	String() string
	Equal(b Codec) bool
	Empty() bool
}
