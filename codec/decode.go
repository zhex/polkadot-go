package codec

import (
	"errors"
	"reflect"
)

func Decode(data []byte, s interface{}) (*ByteInfo, error) {
	val := reflect.Indirect(reflect.ValueOf(s))
	t := val.Type()
	// if out source is a reflect.Value reference, use it directly
	if t.PkgPath() == "reflect" && t.Name() == "Value" {
		val = *(s.(*reflect.Value))
	}
	switch val.Kind() {
	case reflect.String:
		return decodeString(data, val)
	case reflect.Bool:
		return decodeBool(data, val)
	case reflect.Slice:
		return decodeSlice(data, val)
	case reflect.Struct:
	}
	return nil, errors.New("can not decode with type " + val.Kind().String())
}

func decodeString(b []byte, val reflect.Value) (*ByteInfo, error) {
	info := GetBytesInfo(b)
	v := string(b[info.Offset:info.End()])
	val.SetString(v)
	return info, nil
}

func decodeBool(b []byte, val reflect.Value) (*ByteInfo, error) {
	info := &ByteInfo{
		Offset: 0,
		Len:    1,
	}
	if b[0] == 1 {
		val.SetBool(true)
	} else {
		val.SetBool(false)
	}
	return info, nil
}

func decodeSlice(b []byte, val reflect.Value) (*ByteInfo, error) {
	l := int(b[0])
	b = b[1:]
	valSlice := reflect.MakeSlice(val.Type(), l, l)
	for i := 0; i < l; i++ {
		v := valSlice.Index(i)
		info, err := Decode(b, &v)
		if err != nil {
			return nil, err
		}
		b = b[info.End():]
	}
	val.Set(valSlice)
	return nil, nil
}
