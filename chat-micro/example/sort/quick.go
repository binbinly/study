package main

import "fmt"

//see https://www.topgoer.cn/docs/gosuanfa/gosuanfa-1c906k4cpjfnp
//算法描述：是对插入算法的一种优化，利用对问题的二分化，实现递归完成快速排序
//在所有算法中二分化是最常用的方式，将问题尽量的分成两种情况加以分析，
//最终以形成类似树的方式加以利用，因为在比较模型中的算法中，最快的排序时间
//负载度为 O(nlgn).
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitdata := arr[0]          //第一个数据
	low := make([]int, 0, 0)     //比我小的数据
	hight := make([]int, 0, 0)   //比我大的数据
	mid := make([]int, 0, 0)     //与我一样大的数据
	mid = append(mid, splitdata) //加入一个
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitdata {
			low = append(low, arr[i])
		} else if arr[i] > splitdata {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight)
	myarr := append(append(low, mid...), hight...)
	return myarr
}

//快读排序算法
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr))
}
