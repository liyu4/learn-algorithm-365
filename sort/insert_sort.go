package sort

func insertSort(arr []int) []int {
	var (
		i   int
		j   int
		key int
	)

	for i = 1; i < len(arr); i++ {
		// 这一步相当于从底牌中抽出一张
		key = arr[i]

		// 这一步的意义是将key插入排序好的[0:i]因为是左开右闭，所以其实是0~j是排序好的
		j = i - 1

		for j >= 0 && arr[j] > key {
			// 记录比后面元素大的值
			temp := arr[j]
			// 小的元素往前移动一步
			arr[j] = arr[i]
			// 前面大的元素往后移动一步
			arr[j+1] = temp

			// 条件往前移动，已经找的小的元素跟排序的好的所有元素比较，直到找到自己合适的位置
			j--
			arr[j+1] = key
		}
	}
	return arr
}
