package main

import (
	"crypto"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

func main()  {
	sha384Demo()
	sha512Demo()
}

func sha384Demo()  {
	h := sha512.New384()
	h.Write([]byte("hello world"))

	m := h.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m))

	hash := crypto.SHA384
	h2 := hash.New()
	h2.Write([]byte("hello world"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))

	m3 := sha512.Sum384([]byte("hello world"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}

func sha512Demo()  {
	h := sha512.New()
	h.Write([]byte("hello world"))
	m := h.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m))

	hash := crypto.SHA512
	h2 := hash.New()
	h2.Write([]byte("hello world"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))

	m3 := sha512.Sum512([]byte("hello world"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}