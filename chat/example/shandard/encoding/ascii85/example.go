package main

import (
	"bytes"
	"chat/pkg/log"
	"encoding/ascii85"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

// ascii85包实现了ascii85数据编码（5个ascii字符表示4个字节），该编码用于btoa工具和Adobe的PostScript语言和PDF文档格式
func main()  {
	var src = []byte("hello world")

	num := ascii85.MaxEncodedLen(len(src))

	var dst = make([]byte, num)

	ascii85.Encode(dst, src)

	// 转base64字符串输出
	fmt.Println("Ascii85编码内容: ", base64.StdEncoding.EncodeToString(dst))

	var origin = make([]byte, num)

	_, _, err := ascii85.Decode(origin, dst, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ascii85解码内容: ", string(origin))

	var buf bytes.Buffer

	w := ascii85.NewEncoder(&buf)

	if _, err = w.Write([]byte("hello world")); err != nil {
		log.Fatal(err)
	}

	// 关闭
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	// 转base64字符串输出
	fmt.Println("Ascii85编码内容: ", base64.StdEncoding.EncodeToString(buf.Bytes()))

	r := ascii85.NewDecoder(&buf)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ascii85解码内容: ", string(b))
}
