package utils

func ReverseBytes(data []byte) []byte {
	var r []byte
	for _, d := range data {
		r = append([]byte{d}, r...)
	}
	return r
}
