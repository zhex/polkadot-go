package _type

type Digest struct {
	Logs []DigestItem
}

type DigestItem struct {
}

type Signature []byte
type Seal map[uint64]Signature
