package main

import (
	"fmt"
	"slices"

)

// https://leetcode.cn/problems/divide-array-into-arrays-with-max-difference/?envType=daily-question&envId=2025-06-18
func divideArray(nums []int, k int) (ans [][]int) {
	slices.Sort(nums)
	for a := range slices.Chunk(nums, 3) {
		if a[2]-a[0] > k {
			return nil
		}
		ans = append(ans, a)
	}
	return
}

func isOk(current []int, k int) bool {
	for i := 0; i < len(current)-1; i++ {
		for j := i + 1; j < len(current); j++ {
			if (current[i] - current[j]) > k {
				return false
			}
			if (current[j] - current[i]) > k {
				return false
			}
		}
	}
	return true
}

func main() {
	nums := []int{1, 3, 4, 8, 7, 9, 3, 5, 1}
	k := 2
	res := divideArray(nums, k)
	fmt.Println(res)
}
