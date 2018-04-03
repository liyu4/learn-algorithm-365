package search

import (
	"fmt"
)

// In computer science, binary search, also known as half-interval search,
// logarithmic search, or binary chop,
// is a search algorithm that finds the position of a target value within a sorted array.
func binarySearch(in []int, target int) int {
	i, j := 0, len(in)-1
	for i < j {
		mid := (i + j) / 2
		if target > in[mid] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	if in[i] == target {
		return i
	}
	return 0
}

var i int = 1

// 使用递归去实现这个过程
// 确切的说这是一个尾递归（不保存调用者的状态，直接丢弃）
func binarySearchRecursive(start, end int, in []int, target int) int {
	fmt.Println("------------------------------------------------------------------------")
	fmt.Printf("if start %d == end %d\n", start, end)
	fmt.Printf("frame %v: start: %d, end: %d, in: %v, target: %d\n", i, start, end, in, target)
	fmt.Println("call binarySearchRecursive")
	fmt.Println("------------------------------------------------------------------------")
	i++

	if start == end {
		return start
	}
	if start != end {
		mid := (start + end) / 2
		if target > in[mid] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return binarySearchRecursive(start, end, in, target)
}

// 非尾递归的实现，其实是一个函数调用栈
func fact(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * fact(x-1)
	}
}
