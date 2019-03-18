package types

type Codec interface {
	EncodedLength() int
	IsEmpty() bool
	Hex() string
	Sting() string
	Bytes(isBare bool) []byte
}
