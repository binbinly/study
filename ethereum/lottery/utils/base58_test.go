package utils

import "testing"

func Test_encode(t *testing.T) {
	str := "hello world"
	enc := encode(str)
	t.Logf("encode str:%v", enc)
	dec := decode(enc)
	t.Logf("decode str:%v", dec)
}
