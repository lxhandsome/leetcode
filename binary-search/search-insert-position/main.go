package main

import "fmt"

/// https://leetcode.cn/problems/search-insert-position/description/

/*
思路：
使用二分查找查找第一个等于target的数字，如果找到则直接返回下标
如果找不到，结束时候start == end，应该插入的下标 = start + 1
*/
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			if mid == 0 || mid-1 >= 0 && nums[mid-1] != target {
				return mid
			}
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return start
}

func main() {
	nums := []int{1, 3, 5, 7}
	res := searchInsert(nums, 6)
	fmt.Println(res)
}
