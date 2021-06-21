package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// 读取reader内容向指定大小buffer填充
	exampleReadFull()
	// 向writer中写入字符串
	exampleWriteString()
	// writer拷贝reader内容
	exampleCopy()
	// writer拷贝reader指定字节数的内容
	exampleCopayN()
	// writer拷贝reader内容，提供指定大小的缓冲区
	exampleCopyBuffer()
	// 从reader读取至少n字节的内容填充指定大小的buffer
	exampleReadAtLeast()
	// 从reader中读取n个字节并返回一个新的reader
	exampleLimitReader()
	// 将几个reader内容串联起来并返回一个新的reader
	exampleMultiReader()
	// 返回一个将其从r读取的数据写入w的Reader接口。所有通过该接口对r的读取都会执行对应的对w的写入
	exampleTeeReader()
	// 返回一个从reader第off个字节开始往后n个的内容的新reader
	exampleSectionReader()
	// 创建一个可以同时写入多个writer的writer接口
	exampleMultiWriter()
	//// 管道实例
	examplePipe()
}

func exampleReadFull() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)

	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)
}

func exampleWriteString() {
	var buf bytes.Buffer

	if _, err := io.WriteString(&buf, "Hello World"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func exampleCopy()  {
	var buf bytes.Buffer

	r := strings.NewReader("some io.Reader stream to be read")

	if _, err := io.Copy(&buf, r); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func exampleCopayN()  {

	var buf bytes.Buffer

	r := strings.NewReader("some io.Reader stream to be read")

	if _, err := io.CopyN(&buf, r, 10); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func exampleCopyBuffer()  {
	var buf bytes.Buffer

	r1 := strings.NewReader("copy reader")
	b := make([]byte, 4)

	if _, err := io.CopyBuffer(&buf, r1, b); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func exampleReadAtLeast()  {
	r := strings.NewReader("some io.Reader stream to be read\n")
	buf := make([]byte, 10)

	if _, err := io.ReadAtLeast(r, buf, 6); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)
}

func exampleLimitReader()  {
	r := strings.NewReader("some io.Reader stream to be read\n")

	lr := io.LimitReader(r, 4)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

func exampleMultiReader()  {
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")

	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func exampleTeeReader()  {
	r := strings.NewReader("some io.Reader stream to be read\n")
	var buf bytes.Buffer

	tee := io.TeeReader(r, &buf)

	if _, err := io.Copy(os.Stdout, tee); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

func exampleSectionReader()  {
	r := strings.NewReader("some io.Reader stream to be read\n")

	s := io.NewSectionReader(r, 5, 17)

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	be := make([]byte, 6)
	if _, err := s.ReadAt(be, 10); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(be))

	if _, err := s.Seek(10, io.SeekStart); err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

func exampleMultiWriter()  {
	r := strings.NewReader("some io.Reader stream to be read\n")

	var buf1, buf2 bytes.Buffer

	w := io.MultiWriter(&buf1, &buf2)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf1.String())
	fmt.Println(buf2.String())
}

func examplePipe()  {
	r, w := io.Pipe()

	go func() {
		w.Write([]byte("some text to be read\n"))
		w.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	fmt.Println(buf.String())
}