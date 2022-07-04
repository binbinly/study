package net

import (
	"testing"
)

func TestGetIp(t *testing.T) {
	ip, err := GetIP("http://www.baidu.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(ip)
}
