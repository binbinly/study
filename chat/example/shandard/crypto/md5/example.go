package main

import (
	"crypto"
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

// 实现了MD5哈希算法
func main()  {
	// 返回一个新的使用MD5校验的hash.HasH
	h := md5.New()
	// 写入
	h.Write([]byte("hello world"))

	m := h.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m))

	hash := crypto.MD5
	h2 := hash.New()
	h2.Write([]byte("hello world"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))

	m3 := md5.Sum([]byte("hello world"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}