package main

import "fmt"

// https://leetcode.cn/problems/word-break/description/

/*
解题思路：
如果字符串 j - i组成的字符串在单词表中存在 && 0-j也是有单词表中的单词组成的 那么 0 - i 也是可以有单词表中的单词组成的
由上分析可知当前的状态依赖前面的状态，综上考虑使用动态规划

定义状态：dp[i] 以s[i]结尾的字符串能不能由单词表中的单词组成
状态转移方程  dp[i] = s[j+1:i]在单词表中 && dp[j] 为 true
初始状态： dp[0] = s[0] in 单词表
*/

func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	wordM := make(map[string]struct{})
	for _, s := range wordDict {
		wordM[s] = struct{}{}
	}
	if _, ok := wordM[s]; ok {
		return true
	}
	dp := make([]bool, l)
	for i := 0; i < l; i++ {
		// 0 - i字符串的长度 取的切片范围为s[:i+1]
		w := string(s[:i+1])
		if _, ok := wordM[w]; ok {
			dp[i] = true
			continue
		}
		for j := i - 1; j >= 0; j-- {
			w := string(s[j+1 : i+1])
			if _, ok := wordM[w]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[l-1]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
