package main

import "fmt"

//https://leetcode.cn/problems/longest-increasing-subsequence/description/

/*
解题思路: 当钱的最长严格递增子序列依赖前一个数字或者前面的最长严格递增子序列
最值类型的可能是动态规划
定义状态 dp[i] 为以nums[i]结尾的最长严格递增子序列的长度
状态转移 :
    当nums[i] > nums[i-1]的时候 dp[i] = dp[i-1] + 1
    当nums[i] <= nums[i-1]的时候 dp[i] = max(dp[j])+1 j属于[0,i-1] && nums[i] > nums[j] ,意味着i可以包含在i之前的任意最长一个递增子序列中，在这么多子序列中选择最长的那个

上面的逻辑存在问题: 当nums[i] > nums[i-1] 时 dp[i-1] + 1不一定是当前的最大值，也需要遍历左节点查找最大值

状态转移方程：
    当：nums[i] > nums[j] 时 dp[i] = max(dp[j]+1) j属于[0,i-1]





初始值：
    dp[0]=1

*/

func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	dp := make([]int, l)
	dp[0] = 1
	var res int
	for i := 1; i < l; i++ {
		dp[i] = 1
		// 当nums[i] > nums[j],此时dp[i]也不一定是最大的，所以需要遍历所有的左节点，判断当前i的最大严格递增子序列
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func tastCase() []int {
	return []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
}

func main() {
	fmt.Println(lengthOfLIS(tastCase()))
}
