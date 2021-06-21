package main

import (
	"bytes"
	"chat/pkg/log"
	"encoding/gob"
	"fmt"
)

func main()  {
	// 数据编码解码
	example()

	// 对interface编码解码
	exampleInterface()
}

func example()  {

	var data = "hello world"

	var data2 = []string{"1", "2", "hello go"}

	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(data); err != nil {
		log.Fatal(err)
	}

	if err := enc.Encode(data2); err != nil {
		log.Fatal(err)
	}

	var origin string
	var origin2 []string

	dec := gob.NewDecoder(&buf)

	if err := dec.Decode(&origin); err != nil {
		log.Fatal(err)
	}
	if err := dec.Decode(&origin2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(origin, origin2)
}

func exampleInterface()  {
	var buf bytes.Buffer

	gob.Register(Hello{})

	enc := gob.NewEncoder(&buf)

	interfaceEncode(enc, Hello{"world"})

	dec := gob.NewDecoder(&buf)

	interfaceDecode(dec).Say()
}

type Hello struct {
	Name string
}

func (h Hello) Say()  {
	fmt.Println("hello ", h.Name)
}

type Pythagoras interface {
	Say()
}

func interfaceEncode(enc *gob.Encoder, p Pythagoras)  {
	err := enc.Encode(&p)
	if err != nil {
		log.Fatal(err)
	}
}

func interfaceDecode(dec *gob.Decoder) (p Pythagoras) {
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
	return
}
