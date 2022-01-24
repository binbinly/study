package main

import "fmt"

func base() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil)    // <nil> true
	fmt.Println(in, in == nil)    // <nil> true

	in = data
	fmt.Println(in, in == nil)    // <nil> false    // data 值为 nil，但 in 值不为 nil
}

// 正确示例
func main() {
	doIt := func(arg int) interface{} {
		var result *struct{} = nil
		if arg > 0 {
			result = &struct{}{}
		} else {
			return nil    // 明确指明返回 nil
		}
		return result
	}

	if res := doIt(-1); res != nil {
		fmt.Println("Good result: ", res)
	} else {
		fmt.Println("Bad result: ", res)    // Bad result:  <nil>
	}
}