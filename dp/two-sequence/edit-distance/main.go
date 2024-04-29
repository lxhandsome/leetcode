package main

// https://leetcode.cn/problems/edit-distance/description/

/*
解题思路：动态规划
定义状态：dp[i][j] 为 word1前i个字符转换为word2前j个字符所操作的最少次数
转移方程：
    当 word1[i] == word2[j] 时候：dp[i][j] = dp[i-1][j-1]
    当不相同的时候 dp[i][j] = max()
*/

func minDistance(word1 string, word2 string) int {
	return 0
}
