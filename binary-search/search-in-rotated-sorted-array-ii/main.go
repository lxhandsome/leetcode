package main

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[end] < nums[mid] {
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
