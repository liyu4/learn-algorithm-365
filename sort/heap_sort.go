package sort

import (
	"errors"
)

// 参考算法导论的实现
// 第一步将数组中的元素堆化
func max_heapify(arr []int, position int) {
	heapsize := len(arr) - 1

	largest := position
	//假设现在position为0，也就是说从数组的第一个元素开始，我们开始堆化的过程，就是把一个数据转换成堆（近似于完全二叉树）
	//根据以下的代码我们其实很容易的知道，堆化的结果是一个最大堆
	leftchild := 2*position + 1
	rightchild := 2*position + 2

	if leftchild <= heapsize && arr[leftchild] > arr[position] {
		largest = leftchild
	}

	if rightchild <= heapsize && arr[rightchild] > arr[largest] {
		largest = rightchild
	}

	if position != largest {
		//交换子节点数据和父节点数据
		//这里给一个简单的例子[3,4,16] 堆化的结果是[16,4,3],各位同学请想象一下这棵树，是不是很简单和有趣呢！
		arr[position], arr[largest] = arr[largest], arr[position]
		// 父节点改变，堆化递归
		max_heapify(arr, largest) //recursive
	}
}

//自底向上建立堆化
func build_max_heap(arr []int) {
	var i int
	for i = (len(arr) - 1) / 2; i >= 0; i-- {
		max_heapify(arr, i)
	}
}

func heap_sort(arr []int) {
	build_max_heap(arr)

	var i int
	// i 等于1说明是到了最后一个元素，可以看成是一个玩好的堆，已经不需要维护，直接交换。
	for i = len(arr) - 1; i >= 1; i-- {
		//最大元素总是在arr[0]，移除最大节点数据n，然后重新维护这个最大堆（这个堆已经改变）
		arr[i], arr[0] = arr[0], arr[i]
		// 说明一下：最大的元素和最后的元素互换，最大堆的性质极有可能已经被破坏，所以要重新维护
		max_heapify(arr[:i-1], 0)
	}
}

/*
  这部分实现一个最大优先队列
*/

func heap_maximum(arr []int) int {
	return arr[0]
}

func heap_extract_max(arr []int) (int, []int, error) {
	heapsize := len(arr) - 1

	if heapsize < 0 {
		//下溢
		return -1, nil, errors.New("heap underflow")
	}

	max := arr[0]
	arr[0] = arr[heapsize]

	if heapsize == 0 {
		arr = arr[:0:0]
		return max, nil, nil
	}

	arr = arr[:heapsize-1]
	max_heapify(arr, 0)
	return max, arr, nil
}

func heap_increase_key(arr []int, i, key int) error {
	if key < arr[i] {
		return errors.New("new key is small than current key")
	}

	arr[i] = key

	for i > 0 && arr[i/2] < arr[i] {
		arr[i], arr[i/2] = arr[i/2], arr[i]
		i = i / 2
	}
	return nil
}

func max_heap_insert(arr []int, key int) []int {
	newarr := make([]int, len(arr)+1)
	copy(newarr, arr)

	newarr[len(arr)-1] = -10000000000
	heap_increase_key(newarr, len(newarr)-1, key)
	return newarr
}
