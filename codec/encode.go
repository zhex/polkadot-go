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
		b := []byte(data.(string))
		return utils.AddBytePrefix(b, len(b)), nil
	case reflect.Int:
		d := data.(int)
		return utils.IntToByte(d), nil
	case reflect.Uint8:
		return utils.Uint8ToByte(data.(uint8)), nil
	case reflect.Uint16:
		return utils.Uint16ToByte(data.(uint16)), nil
	case reflect.Uint32:
		return utils.Uint32ToByte(data.(uint32)), nil
	case reflect.Uint64:
		return utils.Uint64ToByte(data.(uint64)), nil
	case reflect.Bool:
		return utils.BooltoByte(data.(bool)), nil
	case reflect.Slice:
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
	case reflect.Struct:
		name := t.Name()
		if encoder, ok := encodeMap[name]; ok {
			return encoder(data), nil
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
	return nil, errors.New("unknown encode type " + t.Name())
}
