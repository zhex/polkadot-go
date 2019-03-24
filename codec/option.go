package codec

import "reflect"

type Option struct {
	Value interface{}
}

func (o *Option) Empty() bool {
	return o.Value == nil
}

func EncodeOption(data interface{}) ([]byte, error) {
	d := data.(Option)
	if d.Empty() {
		return []byte{0}, nil
	}
	sub, err := Encode(d.Value)
	if err != nil {
		return nil, err
	}
	return append([]byte{1}, sub...), nil
}

func DecodeOption(b []byte, target reflect.Value) (interface{}, error) {
	if b[0] == 0 {
		return Option{Value: nil}, nil
	}
	v := target.Field(0).Interface()
	if _, err := Decode(b[1:], v); err != nil {
		return nil, err
	}
	return Option{Value: v}, nil
}
