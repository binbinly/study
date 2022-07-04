package sort

import "log"

//快速排序
func quickSort(nums []int)  {
	_quickSort(nums, 0, len(nums) - 1)
}

func _quickSort(nums []int, left, right int) {
	if left < right {
		partitionIndex := partition(nums, left, right)
		log.Println(nums, left, partitionIndex, right)
		_quickSort(nums, left, partitionIndex - 1)
		_quickSort(nums, partitionIndex+1, right)
	}
}

//分割
//从数列中挑出一个元素，称为 "基准"（pivot）;
//重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（
//相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。
//这个称为分区（partition）操作；
func partition(nums []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[index] = nums[index], nums[i]
			index += 1
		}
	}
	nums[pivot], nums[index-1] = nums[index-1], nums[pivot]
	return index - 1
}