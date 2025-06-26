package main

func QuickSort(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) []int {
	if l >= r {
		return arr
	}
	pivot := arr[l]
	i, j := l, r
	for i < j {
		for i < j && arr[j] > pivot {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] < pivot {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = pivot
	quickSort(arr, l, i-1)
	quickSort(arr, i+1, r)
	return arr
}
