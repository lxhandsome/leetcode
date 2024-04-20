package main

// https://leetcode.cn/problems/palindrome-partitioning-ii/description/
/*
思路： 求最小值，可能为动态固话类型
定义状态：f[i]为字符串[0,i]切割成回文字符串需要切割的最少次数
状态转移：
1. 当[0,i]为回文字符串的话需要切割的次数为0
2. 当[0,i]不是回文字符串的话,此时需要枚举左节点j
	当[j,i]为回文字符串的时候，f[i] = min(f[j-1]+1) j 属于 [0,i)

初始化值 f[0] = 0


*/
func minCut(s string) int {
	l := len(s)
	if l <= 1 {
		return 0
	}
	dp := make([]int, l)
	dp[0] = 0
	is := getIsPalindromeGrid(s)

	for i := 1; i < l; i++ {
		// 0, i 是回文串的情况
		if is[0][i] {
			dp[i] = 0
			continue
		}
		dp[i] = i
		for j := 1; j <= i; j++ {
			if is[j][i] {
				dp[i] = min(dp[j-1]+1, dp[i])
			}
		}
	}

	return dp[l-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 字符串i，j为回文字符串，需要满足
// grid[i-1][j+1]为回文字符串 && s[i] == s[j]
// i 状态依赖于i-1状态, j状态依赖于j+1

/*
使用dp来保存i-j是否为回文字符串
对于一个字符串 有左右两个指针 l，r如果 l+1，r-1是回文字符串的话
if s[l] == s[j] l, r也是回文字符串
根据上式，我们需要知道l+1的状态 和r-1的状态
r 应该从左往右遍历
l 应该从右往左遍历
*/

func getIsPalindromeGrid(s string) [][]bool {
	n := len(s)
	grid := make([][]bool, n)
	for i := range grid {
		grid[i] = make([]bool, n)
	}
	for r := 0; r < n; r++ {
		for l := r; l >= 0; l-- {
			if l == r {
				grid[l][r] = true
				continue
			}
			if s[l] == s[r] {
				if r-l == 1 || grid[l+1][r-1] {
					grid[l][r] = true
				}
			}
		}
	}
	return grid
}

func main() {
	min := minCut("abbab")
	print(min)
}
