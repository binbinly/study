package main

import "math"

// math包提供了基本的数学常数和数学函数
func main()  {

	// 返回一个NaN
	math.NaN()

	math.IsNaN(12.34)

	f := math.Inf(0)

	math.IsInf(12.34, 1)
	math.IsInf(f, 0)


}
