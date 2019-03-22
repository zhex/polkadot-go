package hook

import (
	"github.com/zhex/polkadot-go/types/primitives"
	"github.com/zhex/polkadot-go/utils/decoder"
	"reflect"
	"regexp"
)

func init() {
	decoder.RegisterHook("hash", Hash)
}

func Hash(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	re := regexp.MustCompile(`^Hash(\d*)$`)
	matched := re.FindStringSubmatch(t.Name())
	if f.Kind() == reflect.String && t.Kind() == reflect.Struct && len(matched) > 0 {
		str := data.(string)
		var hash interface{}
		switch matched[1] {
		case "160":
			hash = primitives.NewHash160(str)
		case "256":
			hash = primitives.NewHash256(str)
		case "512":
			hash = primitives.NewHash512(str)
		case "":
			hash = primitives.NewHash256(str)
		}
		return hash, nil
	}
	return data, nil
}
