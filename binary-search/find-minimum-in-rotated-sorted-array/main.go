package main

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/
/*
解题思路：
对数据进行旋转后数组分为有序的两部分
a[0 ~ i -1] 递增的
a[i ~ n-1] 递增的
a[i-1] > a[i]
此时可以使用二分排序来寻找第一个无序的元素
*/
func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			if mid == 0 || mid-1 >= 0 && nums[mid-1] > nums[end] {
				return nums[mid]
			}
			end = mid - 1
		}
	}
	return -1
}
