package codec

import (
	"errors"
	"github.com/zhex/polkadot-go/utils"
	"reflect"
)

func Encode(data interface{}) ([]byte, error) {
	val := reflect.Indirect(reflect.ValueOf(data))
	t := val.Type()

	switch t.Kind() {
	case reflect.String:
		b := []byte(val.String())
		return AddLengthPrefix(b, len(b)), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return utils.IntToByte(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return utils.UintToByte(val.Uint()), nil
	case reflect.Bool:
		return encodeBool(val)
	case reflect.Slice:
		return encodeSlice(val)
	case reflect.Struct:
		return encodeStruct(val)
	}
	return nil, errors.New("unknown encode type " + t.Name())
}

func encodeSlice(val reflect.Value) ([]byte, error) {
	var b []byte
	for i := 0; i < val.Len(); i++ {
		v := val.Index(i).Interface()
		d, err := Encode(v)
		if err != nil {
			return nil, err
		}
		b = append(b, d...)
	}
	return AddLengthPrefix(b, val.Len()), nil
}

func encodeStruct(val reflect.Value) ([]byte, error) {
	t := val.Type()
	name := t.Name()
	if encoder, ok := encodeMap[name]; ok {
		return encoder(val.Interface())
	}
	var b []byte
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		v := val.FieldByIndex(f.Index).Interface()
		d, err := Encode(v)
		if err != nil {
			return nil, err
		}
		b = append(b, d...)
	}
	return AddLengthPrefix(b, len(b)), nil
}

func encodeBool(val reflect.Value) ([]byte, error) {
	if val.Bool() {
		return []byte{1}, nil
	}
	return []byte{0}, nil
}