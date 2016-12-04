package sort

import ()

func merge(left, right []int) []int {
	arr := make([]int, len(left)+len(right))

	// index p for left, q for right

	p, q := 0, 0

	for i := 0; i < len(arr); i++ {

		// fix for index out of range without using sentinel
		{
			if p >= len(left) {
				arr[i] = right[q]
				q = q + 1
				continue
			} else if q >= len(right) {
				arr[i] = left[p]
				p = p + 1
				continue
			}
		}

		// default loop conditon define in algorithm
		if left[p] >= right[q] {
			arr[i] = right[q]
			q = q + 1
		} else {
			arr[i] = left[p]
			p = p + 1
		}
	}
	return arr
}

func sort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2

	left := array[:mid]
	right := array[mid:]

	left = sort(left)
	right = sort(right)

	return merge(left, right)
}
