package sort

//快速排序的最简单的思想可能就是选取一个基准数（可以是随机的也可以是数组的第一个数）
//然后将大于这个数的数放在右边，将小于这个数的数放在左边
//最后递归调用这个过程
func quick_sort(arr []int, start, lenght int) []int {
	//显然如果剩下一个元素的时候就不需要排序了。
	if lenght-1 > start {
		i := start
		j := lenght - 1
		x := arr[i]

		//开始左右交换
		//直到i = j 说明左右的步骤执行完毕
		for i < j {
			//从右边开始找到第一个小于x的值，则停止循环
			for i < j && (arr[j] >= x) {
				j--
			}

			//将找到的第一个小于x的值填入第一个基准点
			if i < j {
				arr[i] = arr[j]
				i++
			}

			// 从右边开始找第一个大于x的值，找到则停止循环
			for i < j && (arr[i] < x) {
				i++
			}

			//将找到的第一个比x大的值填入之前空白的地方，也就是比x小的那个值的地方
			if i < j {
				arr[j-1] = arr[i]
			}
		}

		arr[i] = x

		quick_sort(arr, start, i-1)
		quick_sort(arr, i+1, lenght)
	}
	return arr
}
