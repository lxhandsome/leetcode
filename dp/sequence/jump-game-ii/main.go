package main

import "math"

// https://leetcode.cn/problems/jump-game-ii/description/

/*
思路：动态规划dp
定义状态： dp[i]为到达i出需要的最少跳跃次数
状态转移方程：dp[i] =  if (num[j] + j >= i) 能到达i位置 min(dp[j] + 1,m) j属于[0,i-1)
初始值： dp[0] = 0
*/

func jump1(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		m := math.MaxInt
		for j := i - 1; j >= 0; j-- {
			if nums[j]+j >= i {
				// 从j位置一步能够跳道i位置
				m = min(m, dp[j]+1)
			}
		}
		dp[i] = m
	}
	return dp[n-1]
}

/*
解题思路：贪心算法
从最后一个位置出发找到最远的能够跳到当前节点的位置pos记录下来，下次从pos接着向前走，知道0
*/
func jump(nums []int) int {
	n := len(nums)
	pos := n - 1
	step := 0
	for pos > 0 {
		for i := 0; i < n; i++ {
			if nums[i]+i >= pos {
				pos = i
				step++
				break
			}
		}
	}
	return step

}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func testCase() []int {
	return []int{2, 3, 1, 1, 4}
}

func main() {
	print(jump(testCase()))
}
