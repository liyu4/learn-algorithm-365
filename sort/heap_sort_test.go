package sort

import (
	"testing"
)

func TestHeapsort(t *testing.T) {
	arr := []int{100, 2, 4, 67, 2222, 84, 23, 155, 6, 7, 34, 565, 6}
	heap_sort(arr)

	t.Log(arr)
}
