package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const FilePath = "example/testdata/bufio"

func main()  {
	// 缓冲区写入
	buf := write()
	fmt.Printf("%s\n", buf.Bytes())

	//写入文件读取测试使用
	if err := ioutil.WriteFile(FilePath, buf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	//缓冲区读取
	read(FilePath)

	//扫描器
	scannerDemo()
	//扫描器自定义
	scannerCustomDemo()
	//扫描器切割最后为空是避免报错
	scannerSplitWithCommaDemo()
}

func write() bytes.Buffer {
	// 声明 buffer
	var buf bytes.Buffer

	// 初始化 writer
	var w = bufio.NewWriter(&buf)

	w.WriteString("hello,")
	w.WriteRune('中')
	w.WriteByte('o')
	w.Write([]byte("world"))
	// 重置当前缓冲区
	w.Flush()

	return buf
}

func read(path string) {

	var buf bytes.Buffer

	// 读取文件内容
	bf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	//初始化reader
	var r = bufio.NewReader(bytes.NewReader(bf))

	// 读取内容查到除夕拿 o 停止并输出 其余内容扔在缓冲区
	if s, err := r.ReadString('o'); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(s)
	}

	// 返回可以从当前缓冲区读取的字节数
	fmt.Println(r.Buffered())

	// 将缓冲区内容写入buffer
	r.WriteTo(&buf)
	fmt.Printf("%s\n", buf.Bytes())
}

func scannerDemo()  {

	//声明字符串
	input := "foo bar bar"
	// 初始化扫描器
	scanner := bufio.NewScanner(strings.NewReader(input))
	//split调用一个split函数，函数内容可以根据格式自己定义
	// scanwords 是一个自带split函数，用于返回以空格分隔的字符串，删除了周围的空格。它绝不返回空字符串。
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func scannerCustomDemo()  {
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))

	// 自定义split函数
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}

	scanner.Split(split)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Invalid input:", err)
	}
}

func scannerSplitWithCommaDemo()  {

	const input = "1,2,3,4"
	scanner := bufio.NewScanner(strings.NewReader(input))
	onComma := func(data []byte, atEOR bool) (advance int, token []byte, err error)  {
		for i := 0; i<len(data);i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Invalid input: ", err)
	}
}