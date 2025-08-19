package main

import "fmt"

func main() {
	fmt.Printf("zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}): %v\n", zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))
}

func zeroFilledSubarray(nums []int) int64 {
	// return windows(nums)
	return zeroSequence(nums)
}

/*
1. 使用滑动窗口，记录窗口的左端点last为非0的元素下标，右端点为0元素下标

右端点为i时，有 last+1,last+2 ... i 子数组，总共数量是 i -last个

*/
func windows(nums []int) (res int64) {
	last := -1
	for i, x := range nums {
		if x != 0 {
			last = i
			continue
		}
		res += int64(i - last)
	}
	return
}

/*
2. 使用计数法
计算连续0的序列长度 n0，计算长度为n0的序列的连续子序列的数量
*/

func zeroSequence(nums []int) int64 {
	res := int64(0)
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] != 0 {
			i++
		}
		j := i
		for j < n && nums[j] == 0 {
			j++
		}
		res += int64(calculate(j - i))
		i = j
	}
	return res
}

func calculate(n int) int {
	if n%2 == 0 {
		return n / 2 * (n + 1)
	}
	return (n + 1) / 2 * n
}
