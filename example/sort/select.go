package sort

import "log"

//选择排序
func selectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if min != i {
			nums[min], nums[i] = nums[i], nums[min]
		}
		log.Println(nums)
	}
}
