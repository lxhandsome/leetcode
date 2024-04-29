package main

import "fmt"

// https://leetcode.cn/problems/longest-common-subsequence/description/

/*
解题思路：求最大值可能为动态规划
分析：如果知道text1 0-i , text2 0-j 最长的公共子序列为w ，如果 text[i+1] = text[j+1] 则 text 0-i+1 text 0-j+1 最长公共子序列为 w + 1，所以可得之后的状态是依赖前面的状态的，所以可用动态规划

定义状态：dp[i][j] text1 前i个字符和 text2 前j个字符最长的公共子序列
状态转移方程：
    当 text1[i] = text1[j]时 ， dp[i][j] = dp[i-1][j-1] +1
    当 text1[i] != text1[j], dp[i][j] = max(dp[i-1][j],dp[i][j-1])

*/

func longestCommonSubsequence(text1 string, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)
	if l1 == 0 || l2 == 0 {
		return 0
	}
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
