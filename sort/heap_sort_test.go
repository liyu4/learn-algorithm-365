package sort

import (
	"testing"
)

func TestHeapsort(t *testing.T) {
	arr := []int{100, 2, 4, 67, 2222, 84, 23, 155, 6, 7, 34, 565, 6}
	heap_sort(arr)

	t.Log(arr)

	arr2 := []int{2, 3, 2, 5, 333, 5}

	build_max_heap(arr2)

	// newarr := max_heap_insert(arr2, 100)
	// t.Log(newarr)

	max, result, err := heap_extract_max(arr2)

	if err != nil {
		t.Log(err)
	}

	t.Log(max)
	t.Log(result)
}
