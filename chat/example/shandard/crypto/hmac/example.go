package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// 实现了U.S. Federal Information Processing Standards Publication 198规定的HMAC（加密哈希信息认证码）
func main()  {
	// 声明秘钥
	var key = []byte("test hmac key")

	// 声明内容
	var content = []byte("hello world")

	h := hmac.New(sha256.New, key)

	h.Write(content)

	expectedMAC := h.Sum(nil)

	b := hmac.Equal(expectedMAC, expectedMAC)
	fmt.Println("b", b)
}