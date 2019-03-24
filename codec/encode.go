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
		return utils.AddBytePrefix(b, len(b)), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return utils.IntToByte(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return utils.UintToByte(val.Uint()), nil
	case reflect.Bool:
		return utils.BoolToByte(val.Bool()), nil
	case reflect.Slice:
		return decodeSlice(val)
	case reflect.Struct:
		return decodeStruct(val)
	}
	return nil, errors.New("unknown encode type " + t.Name())
}

func decodeSlice(val reflect.Value) ([]byte, error) {
	var b []byte
	for i := 0; i < val.Len(); i++ {
		v := val.Index(i).Interface()
		d, err := Encode(v)
		if err != nil {
			return nil, err
		}
		b = append(b, d...)
	}
	return utils.AddBytePrefix(b, val.Len()), nil
}

func decodeStruct(val reflect.Value) ([]byte, error) {
	t := val.Type()
	name := t.Name()
	if encoder, ok := encodeMap[name]; ok {
		return encoder(val.Interface()), nil
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
	return utils.AddBytePrefix(b, len(b)), nil
}
