package main

import (
	"bytes"
	"chat/pkg/log"
	"encoding/base32"
	"fmt"
	"io/ioutil"
)

func main()  {
	var origin = []byte("hello world")

	var buf bytes.Buffer
	// 自定一个32字节的字符串
	var customEncode = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	e := base32.NewEncoding(customEncode)
	w := base32.NewEncoder(e, &buf)
	if _, err := w.Write(origin); err != nil {
		log.Fatal(err)
	}
	if err := w.Close();err != nil {
		log.Fatal(err)
	}
	fmt.Println("base32编码内容: ", string(buf.Bytes()))

	r :=base32.NewDecoder(base32.StdEncoding, &buf)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("base32解码内容: ", string(b))


	// 使用标准的base32编码字符集编码
	originEncode := base32.StdEncoding.EncodeToString(origin)
	fmt.Println("base32编码内容: ", originEncode)

	// 使用标准的base32编码字符集解码
	originBytes, err := base32.StdEncoding.DecodeString(originEncode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("base32解码内容: ", string(originBytes))


	var ne = base32.StdEncoding.EncodedLen(len(origin))
	var dst = make([]byte, ne)

	base32.StdEncoding.Encode(dst, origin)
	fmt.Println("base32编码内容: ", string(dst))

	var nd = base32.StdEncoding.DecodedLen(len(dst))
	var originText = make([]byte, nd)
	if _, err = base32.StdEncoding.Decode(originText, dst); err != nil {
		log.Fatal(err)
	}
	fmt.Println("base32解码内容: ", string(originText))
}
