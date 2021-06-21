package main

import (
	"encoding/asn1"
	"fmt"
	"log"
)

type Hello struct {
	Num int
	Str string
}

// asn1包实现了DER编码的ASN.1数据结构的解析
func main() {
	hello := Hello{
		Num: 12,
		Str: "Go",
	}

	b, err := asn1.Marshal(hello)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b)

	var h Hello

	if _, err = asn1.Unmarshal(b, &h); err != nil {
		log.Fatal(err)
	}
	fmt.Println(h)
}
