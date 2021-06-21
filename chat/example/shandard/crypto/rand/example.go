package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

// 实现了用于加解密的更安全的随机数生成器
func main()  {
	b := make([]byte, 10)

	if _, err := rand.Read(b); err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)

	i, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i.Bytes())

	p, err := rand.Prime(rand.Reader, 11)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.Bytes())
}