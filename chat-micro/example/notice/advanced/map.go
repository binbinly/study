package main

import "fmt"

type data struct {
	name string
}

func main() {
	m := map[string]data{
		"x": {"Tom"},
	}
	_ = m
	// 无法直接更新 struct 的字段值
	//m["x"].name = "Jerry"

	//因为 map 中的元素是不可寻址的。需区分开的是，slice 的元素可寻址
	s := []data{{"Tom"}}
	s[0].name = "Jerry"
	fmt.Println(s)    // [{Jerry}]

	//更新 map 中 struct 元素的字段值，有 2 个方法：
	// 提取整个 struct 到局部变量中，修改字段值后再整个赋值
	m = map[string]data{
		"x": {"Tom"},
	}
	r := m["x"]
	r.name = "Jerry"
	m["x"] = r
	fmt.Println(m)    // map[x:{Jerry}]

	//使用指向元素的 map 指针
	mm := map[string]*data{
		"x": {"Tom"},
	}

	mm["x"].name = "Jerry"    // 直接修改 m["x"] 中的字段
	fmt.Println(m["x"])    // &{Jerry}
}
