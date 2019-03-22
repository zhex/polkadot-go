package hook

import (
	"github.com/zhex/polkadot-go/utils"
	"github.com/zhex/polkadot-go/utils/decoder"
	"reflect"
	"strconv"
)

func init() {
	decoder.RegisterHook("uint", Uint)
}

func Uint(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() == reflect.String && t.Kind() == reflect.Uint64 {
		str := data.(string)
		if utils.IsHex(str) {
			d := utils.HexStripPrefix(str)
			return strconv.ParseUint(d, 16, 64)
		} else {
			return strconv.Atoi(str)
		}
	}
	return data, nil
}
