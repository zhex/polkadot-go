package codec

type Struct struct {
	codecMap map[string]Codec
	Types    []interface{}
}

func (s *Struct) Values() []Codec {
	var arr []Codec
	for _, v := range s.codecMap {
		arr = append(arr, v)
	}
	return arr
}

func (s *Struct) EncodedLength() int {
	l := 0
	for _, v := range s.codecMap {
		l += v.EncodedLength()
	}
	return l
}

func (s *Struct) Empty() bool {
	for _, v := range s.codecMap {
		if !v.Empty() {
			return false
		}
	}
	return true
}

func (s *Struct) U8a() *U8a {
	var u []byte
	for _, v := range s.codecMap {
		u = append(u, v.U8a().Value()...)
	}
	return &U8a{
		value: u,
	}
}

func (s *Struct) Hex() string {
	return s.U8a().Hex()
}
