package main

import (
	"fmt"
	"unsafe"
)

func main()  {

	var hello = Hello{}

	s := unsafe.Sizeof(hello)
	fmt.Println(s)

	f := unsafe.Offsetof(hello.b)
	fmt.Println(f)

	a := unsafe.Alignof(hello)
	fmt.Println(a)
}

type Hello struct {
	a bool
	b string
	c int
	d []float64
}