package sort

import ()

/*
   this is an sample example
   if you want to run it
   MergeSort(testArry, 0, 4, 8)
   you need add above function to main function

   it's like

   packgae main

   import(
		"fmt"
   )

   var testArry = []int{2, 4, 5, 7, 1, 2, 3, 6}
   func main() {
	   MergeSort(testArray, 0, 4, 8)
	   fmt.Println(testArray)
   }

*/

func MergeSort(A []int, p, q, r int) {

	n1 := q - p //4
	n2 := r - q //4

	left := make([]int, n1+1)
	right := make([]int, n2+1)

	left[n1] = 100
	right[n2] = 100

	for i := 0; i < n1; i++ {
		left[i] = A[p+i]
	}

	for j := 0; j < n2; j++ {
		right[j] = A[q+j]
	}

	i := 0
	j := 0

	// r-1 step
	for k := p; k < r; k++ {
		if left[i] <= right[j] {
			A[k] = left[i]
			i = i + 1
		} else {
			A[k] = right[j]
			j = j + 1
		}
	}
}
