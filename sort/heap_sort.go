package sort

import (
	"fmt"
)

// 参考算法导论的实现
// 第一步将数组中的元素堆化
func max_heapify(arr []int, position int) {
	heapsize := len(arr)
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
