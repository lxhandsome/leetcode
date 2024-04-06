package main

import "fmt"

// https://leetcode.cn/problems/search-a-2d-matrix/

/*
思路：
每一行按照非严格递增排列 a[i] <= a[i+1]
每行的第一个数大于等于前一行的最后一个数
所以整个二维 数组可以看做一个一位数组，对一位数组进行二分查找即可
所以需要根据mid的值找到元素在二维数组中的实际位置
0 <= mid <= m*n -1
r := mid / n
c := mid % n
*/
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	start := 0
	end := m*n - 1
	for start <= end {
		mid := (start + end) / 2
		i := mid / n
		j := mid % n
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1}, {3}}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
}
