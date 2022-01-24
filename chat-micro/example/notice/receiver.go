package main

import "fmt"

type data struct {
	num   int
	key   *string
	items map[string]bool
}

func (this *data) pointerFunc() {
	this.num = 7
}

func (this data) valueFunc() {
	this.num = 8
	*this.key = "valueFunc.key"
	this.items["valueFunc"] = true
}

//35.若函数 receiver 传参是传值方式，则无法修改参数的原有值
//方法 receiver 的参数与一般函数的参数类似：如果声明为值，那方法体得到的是一份参数的值拷贝，此时对参数的任何修改都不会对原有值产生影响。
//除非 receiver 参数是 map 或 slice 类型的变量，并且是以指针方式更新 map 中的字段、slice 中的元素的，才会更新原有值:
func main() {
	key := "key1"

	d := data{1, &key, make(map[string]bool)}
	fmt.Printf("num=%v  key=%v  items=%v\n", d.num, *d.key, d.items)

	d.pointerFunc()    // 修改 num 的值为 7
	fmt.Printf("num=%v  key=%v  items=%v\n", d.num, *d.key, d.items)

	d.valueFunc()    // 修改 key 和 items 的值
	fmt.Printf("num=%v  key=%v  items=%v\n", d.num, *d.key, d.items)
}