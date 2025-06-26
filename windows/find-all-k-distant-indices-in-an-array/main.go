package main

import "fmt"

func main() {
	nums := []int{3, 4, 9, 1, 3, 9, 5}
	key := 9
	k := 1
	fmt.Println(findKDistantIndices(nums, key, k))
}

func findKDistantIndices(nums []int, key, k int) (ans []int) {
	last := -k - 1 // 保证 key 不存在时 last < i-k
	for i := k - 1; i >= 0; i-- {
		if nums[i] == key {
			last = i
			break
		}
	}

	for i := range nums {
		if i+k < len(nums) && nums[i+k] == key {
			last = i + k
		}
		if last >= i-k { // key 在窗口中
			ans = append(ans, i)
		}
	}
	return
}
