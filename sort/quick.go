package sort

/*
快速排序思路：

分治
1，在数组中选取一个基准key，找到key在数组中的位置，保证key左边的元素小于等于key，key右边的元素大于等于key

然后再key的左右两边执行同样的操作[0,keyIndex][keyIndex+1:len(arr)]，直到输入的下标相同为止

当结束的时候就能保证数组有序
*/

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	key := arr[start]
	i := start
	j := end
	for i < j {
		// j从走往左走找到第一个小于key的元素
		for i < j && j > 0 && arr[j] > key {
			j--
		}
		// i从左往右找找到第一个大于key的元素
		for i < j && i < len(arr)-1 && arr[i] < key {
			i++
		}
		if i < j {
			swap(arr, i, j)
		}
	}
	arr[i] = key
	quickSort(arr, start, i)
	quickSort(arr, i+1, end)
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
