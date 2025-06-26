package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 6, 1, 2, 5}
	k := 2
	fmt.Println(partitionArray(nums, k))
}

func partitionArray(nums []int, k int) int {
	if len(nums) <= 0 {
		return 0
	}
	sort.Ints(nums)
	res := 0
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[start] > k {
			start = i
			res++
		}
	}
	return res + 1

}
