package main

import (
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main()  {
	sha224Demo()
	sha256Demo()
}

func sha224Demo()  {
	h := sha256.New224()
	h.Write([]byte("hello world"))

	m := h.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m))

	hash := crypto.SHA224
	h2 := hash.New()
	h2.Write([]byte("hello world"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))

	m3 := sha256.Sum224([]byte("hello world"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}

func sha256Demo()  {
	h := sha256.New()
	h.Write([]byte("hello world"))
	m := h.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m))

	hash := crypto.SHA256
	h2 := hash.New()
	h2.Write([]byte("hello world"))
	m2 := h2.Sum(nil)
	fmt.Println(base64.StdEncoding.EncodeToString(m2))

	m3 := sha256.Sum256([]byte("hello world"))
	fmt.Println(base64.StdEncoding.EncodeToString(m3[:]))
}
