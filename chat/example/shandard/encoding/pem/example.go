package main

import (
	"bytes"
	"encoding/pem"
	"fmt"
	"log"
)

func main()  {
	var buf bytes.Buffer

	block := &pem.Block{
		Type: "PUBLIC KEY",
		Headers: map[string]string{
			"author":"zc",
			"name":"gopher",
		},
		Bytes: []byte("test"),
	}

	if err := pem.Encode(&buf, block); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))

	dk, _ := pem.Decode(buf.Bytes())
	fmt.Println(dk)
}