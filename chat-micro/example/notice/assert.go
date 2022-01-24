package main

import "fmt"

// 在类型断言语句中，断言失败则会返回目标类型的“零值”，断言变量与原来变量混用可能出现异常情况：
func main() {
	var data interface{} = "great"

	if res, ok := data.(int); ok {
		fmt.Println("[is an int], data: ", res)
	} else {
		fmt.Println("[not an int], data: ", data)    // [not an int], data:  great
	}
}