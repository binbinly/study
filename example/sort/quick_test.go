package sort

import "testing"

func Test_quickSort(t *testing.T) {
	nums := []int{3, 1, 8, 5, 7, 2, 4, 9, 6}
	t.Log("beg", nums)
	quickSort(nums)
	t.Log("end", nums)
}