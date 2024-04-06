package main

import "fmt"

// https://www.lintcode.com/problem/search-for-a-range/description

/*
思路：
两次二分搜索，第一次搜索target开始的索引，第二次搜索target结束的索引
*/
func SearchRange(a []int, target int) []int {
	left := searchLeft(a, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := searchRight(a, target)
	return []int{left, right}
}

// 搜索target第一次出现的地方，如果a[mid] == tartget
// 则此时应该继续向左搜索
// 		如果a[mid] == target 且已经走到最左边，直接返回mid
// 		如果a[mid] == target 且还没有走到最左边，但是mid的前一个已经不等于target，也是直接返回
// 如果a[mid] > target end = mid - 1
// 如果a[mid] < target start = mid + 1

func searchLeft(a []int, target int) int {
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == target {
			if mid == 0 || mid-1 >= 0 && a[mid-1] != target {
				return mid
			}
			end = mid - 1
		} else if a[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func searchRight(a []int, target int) int {
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == target {
			if mid == len(a)-1 || mid+1 < len(a) && a[mid+1] != target {
				return mid
			}
			start = mid + 1
		} else if a[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func main() {
	a := []int{1}
	res := SearchRange(a, 1)
	fmt.Printf("%v", res)
}
