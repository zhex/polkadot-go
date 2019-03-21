package codec

import (
	"encoding/binary"
	"errors"
	"reflect"

	"github.com/zhex/polkadot-go/utils"
)

func DecodeStruct(s interface{}, data []byte) error {
	val := reflect.Indirect(reflect.ValueOf(s))
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		var offset uint64
		var length uint64

		f := t.Field(i)
		fVal := val.FieldByIndex(f.Index)

		if !fVal.CanSet() {
			return errors.New("can not set field")
		}

		switch f.Type.Kind() {
		case reflect.String:
			offset, length = utils.DecodeBytes(data)
			v := string(data[offset : offset+length])
			fVal.SetString(v)
		case reflect.Uint32:
			offset = 0
			length = 4
			v := binary.LittleEndian.Uint32(data[offset:length])
			fVal.SetUint(uint64(v))
		default:
			return errors.New("unknow type to decode")
		}
		data = data[offset+length:]
	}

	return nil
}
