package main

import "fmt"

// 在一个 defer 延迟执行的函数中调用 recover() ，它便能捕捉 / 中断 panic
func main() {
	defer func() {
		fmt.Println("recovered: ", recover())
	}()
	panic("not good")
}