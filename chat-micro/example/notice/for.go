package main

import (
	"fmt"
	"time"
)

//for 语句中的迭代变量在每次迭代中都会重用，即 for 中创建的闭包函数接收到的参数始终是同一个变量，
//在 goroutine 开始执行时都会得到同一个迭代值：
func forRange() {
	data := []string{"one", "two", "three"}

	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}

	time.Sleep(3 * time.Second)
	// 输出 three three three
}

func forRange2() {
	data := []string{"one", "two", "three"}

	for _, v := range data {
		vCopy := v
		go func() {
			fmt.Println(vCopy)
		}()
	}

	time.Sleep(3 * time.Second)
	// 输出 one two three
}

func main() {
	data := []string{"one", "two", "three"}

	for _, v := range data {
		go func(in string) {
			fmt.Println(in)
		}(v)
	}

	time.Sleep(3 * time.Second)
	// 输出 one two three
}