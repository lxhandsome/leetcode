package main


// https://leetcode.cn/problems/is-subsequence/
func isSubsequence(s, t string) bool {
	n := len(t)
	nxt := make([][26]int, n+1)
	for j := range nxt[n] {
		nxt[n][j] = n
	}
	for i := n - 1; i >= 0; i-- {
		nxt[i] = nxt[i+1]
		nxt[i][t[i]-'a'] = i
	}

	// 这个写法无论 s 为空还是 t 为空，都能算出正确答案
	i := -1
	for _, b := range s {
		i = nxt[i+1][b-'a']
		if i == n { // b 不在 t 中，说明 s 不是 t 的子序列
			return false
		}
	}
	return true // s 是 t 的子序列
}
