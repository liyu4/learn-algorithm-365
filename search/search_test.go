package search

import (
	"fmt"
	"testing"
)

// binarySearch
func TestBinarySearch(t *testing.T) {
	var in = []int{1, 2, 3, 5, 6, 7, 8, 10, 12, 13, 15, 16, 18, 19, 20, 22}
	pos := binarySearch(in, 19)
	fmt.Println(pos)

	pos = binarySearch(in, 29)
	fmt.Println(pos)
}

func TestBinarySearchRecusive(t *testing.T) {
	var in = []int{1, 2, 3, 4}
	pos := binarySearchRecursive(0, len(in)-1, in, 4)
	if in[pos] == 4 {
		fmt.Println(pos)
	} else {
		fmt.Println("not find")
	}
	fmt.Printf("%d\n", fact(3))
}
