package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{2, 6, 1, 1, 4, 3, 76, 3, 567, 32}
	ret := quick_sort(arr, 0, len(arr))
	t.Log(ret)
}
