package util

import "testing"

func TestRandom(t *testing.T) {
	str := Random(32)
	t.Logf("str: %v", str)
}
