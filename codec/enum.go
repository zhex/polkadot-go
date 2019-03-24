package codec

import "reflect"

type Enum uint8

type EnumType struct {
	Enum
	Def []interface{}
}

func (e *EnumType) Type() interface{} {
	return e.Def[e.Enum]
}

type EnumTypeDecoded struct {
	Index uint8
	Value interface{}
}

func EncodeEnumType(e interface{}) []byte {
	et := e.(EnumType)
	sub, _ := Encode(et.Type())
	return append([]byte{uint8(et.Enum)}, sub...)
}

func DecodeEnumType(b []byte, target reflect.Value) (interface{}, error) {
	et := target.Interface().(EnumType)
	et.Enum = Enum(b[0])
	t := et.Type()

	err := Decode(t, b[1:])
	if err != nil {
		return nil, err
	}
	return EnumTypeDecoded{
		Index: b[0],
		Value: t,
	}, nil
}
