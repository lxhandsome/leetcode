package main

import (
	"fmt"
	"slices"

)

func main() {
	res := permuteUnique([]int{1, 1, 2})
	for _, a := range res {
		fmt.Println(a)
	}
}

// https://leetcode.cn/problems/permutations-ii/solutions/3059690/ru-he-qu-zhong-pythonjavacgojsrust-by-en-zlwl/

func permuteUnique(nums []int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	path := make([]int, n)    // 所有排列的长度都是 n
	onPath := make([]bool, n) // onPath[j] 表示 nums[j] 是否已经填入排列
	var dfs func(int)
	dfs = func(i int) { // i 表示当前要填排列的第几个数
		if i == n { // 填完了
			ans = append(ans, slices.Clone(path))
			return
		}
		// 枚举 nums[j] 填入 path[i]
		for j, on := range onPath {
			// 如果 nums[j] 已填入排列，continue
			// 如果 nums[j] 和前一个数 nums[j-1] 相等，且 nums[j-1] 没填入排列，continue
			if on || j > 0 && nums[j] == nums[j-1] && !onPath[j-1] {
				continue
			}
			path[i] = nums[j] // 填入排列
			onPath[j] = true  // nums[j] 已填入排列（注意标记的是下标，不是值）
			dfs(i + 1)        // 填排列的下一个数
			onPath[j] = false // 恢复现场
			// 注意 path 无需恢复现场，因为排列长度固定，直接覆盖就行
		}
	}
	dfs(0)
	return ans
}
