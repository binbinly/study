package sort

import "log"

//归并排序
//申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
//设定两个指针，最初位置分别为两个已经排序序列的起始位置；
//比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
//重复步骤 3 直到某一指针达到序列尾；
//将另一序列剩下的所有元素直接复制到合并序列尾。
func mergeSort(nums []int) []int {
	length := len(nums)
	if length <2 {
		return nums
	}
	mid := length / 2
	return merge(mergeSort(nums[0:mid]), mergeSort(nums[mid:]))
}

func merge(left, right []int) []int {
	log.Println("source", left, right)
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	log.Println(result)
	return result
}