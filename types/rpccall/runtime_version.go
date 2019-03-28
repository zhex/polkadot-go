package rpccall

type RuntimeVersionApiValue [2]interface{}

func (r RuntimeVersionApiValue) ID() string {
	return r[0].(string)
}

func (r RuntimeVersionApiValue) Version() uint {
	return r[1].(uint)
}

type RuntimeVersion struct {
	SpecName         string
	ImplName         string
	AuthoringVersion uint32
	SpecVersion      uint32
	ImplVersion      uint32
	APIs             []RuntimeVersionApiValue
}
