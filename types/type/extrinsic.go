package _type

import "reflect"

type Extrinsic struct {
}

func EncodeExtrinsic(ex interface{}) ([]byte, error) {
	// todo
	return []byte("mock"), nil
}

func DecodeExtrinsic(b []byte, target reflect.Value) (interface{}, error) {
	// todo
	return nil, nil
}
