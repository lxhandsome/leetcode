package main

// https://leetcode.cn/problems/climbing-stairs/description/

/*
思路：
dp动态规划

定义状态 f[i] 有多少种方法到达i层
状态转移 f[i] = f[i-1] + f[i-1]
初始状态:

	f[0] =1
	f[1] =1

优化：根据状态转移方程可以看到f[i]只与前两次的和关，可以使用两个长度的数组来保存状态

状态转移方程：f[i%2] = f[(i-1)%2] + f[(i-2)%2]
*/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, 2)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i%2] = dp[(i-1)%2] + dp[(i-2)%2]
	}
	return dp[(n-1)%2]
}
