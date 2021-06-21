package main

import (
	"crypto/subtle"
	"fmt"
)

// 实现了在加密代码中常用的功能，但需要仔细考虑才能正确使用
// 比如[]byte中含有验证用户身份的数据（密文哈希、token等）的时候使用
func main()  {

	var x, y int32 = 6, 8
	var a, b byte = 1, 2
	var c, d = 1, 2
	var e, f = []byte("22"), []byte("33")

	fmt.Println(subtle.ConstantTimeByteEq(a, b))

	fmt.Println(subtle.ConstantTimeEq(x, y))

	fmt.Println(subtle.ConstantTimeLessOrEq(c, d))

	fmt.Println(subtle.ConstantTimeCompare(e, f))

	subtle.ConstantTimeCopy(1, e, f)

	subtle.ConstantTimeSelect(0, c, d)
}