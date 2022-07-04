package sort

import "log"

//冒泡排序
func bubbleSort(nums []int) {
	var flag bool

	for i := 0; i < len(nums) - 1; i++ {
		flag = true
		for j := i+1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				flag = false
			}
		}
		if flag {
			break
		}
		log.Println(nums)
	}
}
