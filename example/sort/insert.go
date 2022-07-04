package sort

import "log"

//插入排序
func insertSort(nums []int)  {
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > temp{
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = temp
		log.Println(nums)
	}
}