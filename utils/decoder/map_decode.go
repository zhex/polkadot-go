package decoder

import (
	"github.com/mitchellh/mapstructure"
)

var hooks = map[string]mapstructure.DecodeHookFunc{}

func RegisterHook(name string, f mapstructure.DecodeHookFunc) {
	hooks[name] = f
}

func MapDecode(input interface{}, out interface{}) error {
	config := mapstructure.DecoderConfig{
		Result: out,
	}
	var fn []mapstructure.DecodeHookFunc
	for _, f := range hooks {
		fn = append(fn, f)
	}
	if len(fn) > 0 {
		config.DecodeHook = mapstructure.ComposeDecodeHookFunc(fn...)
	}
	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return err
	}
	err = decoder.Decode(input)
	return nil
}
