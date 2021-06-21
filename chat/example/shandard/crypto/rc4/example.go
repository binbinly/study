package main

import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"log"
)

// 实现了RC4加密算法
func main()  {
	var key = []byte("example key")

	var text = []byte("hello world")

	c, err := rc4.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	c.XORKeyStream(text, text)
	textStr := hex.EncodeToString(text)

	fmt.Println(textStr)
}