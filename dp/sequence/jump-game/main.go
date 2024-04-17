package main

// https://leetcode.cn/problems/jump-game/description/

/*
解题思路：能够到达的问题，使用动态规划
状态：f[i] 能否到达i位置
状态转移方程： f[i] = f[i-j] &&  nums[i-j] + i - j > i  j 属于(0, i]
初始值：f[0] =true

*/

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = true
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[n-1]
}
