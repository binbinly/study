package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

// hex包实现了16进制字符表示的编解码
func main()  {
	// 编码/解码
	example()
	example2()

	// 编码/解码(string)
	exampleString()

	// hex dump编码
	exampleDump()
}

func example()  {
	var src = []byte("hello gopher")

	dst := make([]byte, hex.EncodedLen(len(src)))

	hex.Encode(dst, src)
	fmt.Printf("%s\n", dst)

	origin := make([]byte, hex.DecodedLen(len(dst)))
	_, err := hex.Decode(origin, dst)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", origin)
}

func example2()  {
	var src = []byte("hello gopher")

	var buf bytes.Buffer

	w := hex.NewEncoder(&buf)

	if _, err := w.Write(src); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))

	r := hex.NewDecoder(&buf)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func exampleString()  {
	var src = []byte("hello gopher")

	str := hex.EncodeToString(src)
	fmt.Println(str)

	origin, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(origin))
}

func exampleDump()  {
	content := []byte("go is an open source programming language")

	str := hex.Dump(content)
	fmt.Printf("%s\n", str)

	var buf bytes.Buffer

	dumper := hex.Dumper(&buf)
	defer dumper.Close()

	if _, err := dumper.Write(content); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf.Bytes())
}
