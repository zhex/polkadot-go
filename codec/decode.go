package codec

import (
	"encoding/binary"
	"errors"
	"reflect"

	"github.com/zhex/polkadot-go/utils"
)

func Decode(s interface{}, data []byte) error {
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
			offset, length = GetBytesInfo(data)
			v := string(data[offset : offset+length])
			fVal.SetString(v)
		case reflect.Uint32:
			offset = 0
			length = 4
			v := binary.LittleEndian.Uint32(data[offset:length])
			fVal.SetUint(uint64(v))
		case reflect.Struct:
			// todo: need to decode with the decoder registered in the type registry
		default:
			return errors.New("unknown type to decode")
		}
		data = data[offset+length:]
	}

	return nil
}

func GetBytesInfo(data []byte) (uint64, uint64) {
	flag := data[0] & 0x03
	if flag == 0x00 {
		return 1, uint64(data[0] >> 2)
	} else if flag == 0x01 {
		d := utils.FixByteWidth(data[:2], 8, true)
		return 2, binary.LittleEndian.Uint64(d) >> 2
	} else if flag == 0x02 {
		d := utils.FixByteWidth(data[:4], 8, true)
		return 4, binary.LittleEndian.Uint64(d) >> 2
	}
	l := uint64(data[0]>>2) + 4
	offset := l + 1
	d := utils.FixByteWidth(data[1:offset], 8, true)
	return offset, binary.LittleEndian.Uint64(d)
}
