package main

import (
	"bytes"
	"chat/pkg/log"
	"encoding/json"
	"fmt"
)

func main()  {
	// 声明内容
	var data = []byte(`{"message": "Hello Gopher!"}`)
	// 验证内容是否符合json格式
	json.Valid(data)

	example()
	example2()
}

type Hello struct {
	Name string
	Sex int
}

func example()  {
	hello := Hello{"gopher", 1}

	b, err := json.Marshal(hello)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	originHello := Hello{}
	if err = json.Unmarshal(b, &originHello); err != nil {
		log.Fatal(err)
	}
	fmt.Println(originHello)
}

func example2()  {

	var buf bytes.Buffer

	hello := Hello{"<div>world</div>", 2}

	e := json.NewEncoder(&buf)

	e.SetEscapeHTML(false)

	e.SetIndent("", "")

	if err := e.Encode(hello); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))

	originHello := Hello{}

	d := json.NewDecoder(&buf)

	d.Buffered()

	d.DisallowUnknownFields()

	d.UseNumber()

	d.More()

	if err := d.Decode(&originHello); err != nil {
		log.Fatal(err)
	}
	fmt.Println(originHello)
}