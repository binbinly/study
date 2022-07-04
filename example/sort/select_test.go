package sort

import "testing"

func Test_selectSort(t *testing.T) {
	nums := []int{3, 1, 8, 5, 7, 2, 4, 9, 6}
	t.Log("beg", nums)
	selectSort(nums)
	t.Log("end", nums)
}
