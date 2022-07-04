package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	nums := []int{3, 1, 8, 5, 7, 2, 4, 9, 6}
	t.Log("beg", nums)
	bubbleSort(nums)
	t.Log("end", nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log("beg", nums)
	bubbleSort(nums)
	t.Log("end", nums)
}
