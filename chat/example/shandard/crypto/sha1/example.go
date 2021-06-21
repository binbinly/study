package main

import (
	"crypto"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main()  {

	h := sha1.New()
	h.Write([]byte("hello world"))
	m := h.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m))

	hash := crypto.SHA1
	h2 := hash.New()
	h2.Write([]byte("hello world"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))

	m3 := sha1.Sum([]byte("hello world"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}