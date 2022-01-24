package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

type data1 struct {
	num     int
	fp      float32
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string
	handler interface{}
	ref     *byte
	raw     [10]byte
	//checks [10]func() bool        // 无法比较
	//doIt   func() bool        // 无法比较
	//m      map[string]string    // 无法比较
	//bytes  []byte            // 无法比较
}

//可以使用相等运算符 == 来比较结构体变量，前提是两个结构体的成员都是可比较的类型：
func structCompare() {
	v1 := data1{}
	v2 := data1{}
	fmt.Println("v1 == v2: ", v1 == v2)    // true
}

// 比较相等运算符无法比较的元素
func deepEqual() {
	v1 := data1{}
	v2 := data1{}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(v1, v2))    // true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a1"}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(m1, m2))    // true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	// 注意两个 slice 相等，值和顺序必须一致
	fmt.Println("v1 == v2: ", reflect.DeepEqual(s1, s2))    // true
}

func deepEqual2() {
	var str = "one"
	var in interface{} = "one"
	fmt.Println("str == in: ", reflect.DeepEqual(str, in))    // true

	v1 := []string{"one", "two"}
	v2 := []string{"two", "one"}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(v1, v2))    // false

	data := map[string]interface{}{
		"code":  200,
		"value": []string{"one", "two"},
	}
	encoded, _ := json.Marshal(data)
	var decoded map[string]interface{}
	json.Unmarshal(encoded, &decoded)
	fmt.Println("data == decoded: ", reflect.DeepEqual(data, decoded))    // false
}

//reflect.DeepEqual() 认为空 slice 与 nil slice 并不相等，但注意 byte.Equal() 会认为二者相等：
func main() {
	var b1 []byte = nil
	b2 := []byte{}

	// b1 与 b2 长度相等、有相同的字节序
	// nil 与 slice 在字节上是相同的
	fmt.Println("b1 == b2: ", bytes.Equal(b1, b2))    // true
}
