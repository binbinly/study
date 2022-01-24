package main

import "fmt"

//在 range 迭代中，得到的值其实是元素的一份值拷贝，更新拷贝并不会更改原来的元素，即是拷贝的地址并不是原有元素的地址：
func main() {
	data := []int{1, 2, 3}
	for _, v := range data {
		v *= 10        // data 中原有元素是不会被修改的
	}
	fmt.Println("data: ", data)    // data:  [1 2 3]

	//如果要修改原有元素的值，应该使用索引直接访问：
	for i, v := range data {
		data[i] = v * 10
	}
	fmt.Println("data: ", data)    // data:  [10 20 30]

	//如果你的集合保存的是指向值的指针，需稍作修改。依旧需要使用索引访问元素，不过可以使用 range 出来的元素直接更新原有值：
	data1 := []*struct{ num int }{{1}, {2}, {3}}
	for _, v := range data1 {
		v.num *= 10    // 直接使用指针更新
	}
	fmt.Println(data1[0], data1[1], data1[2])    // &{10} &{20} &{30}
}