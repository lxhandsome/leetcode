package main

// https://leetcode.cn/problems/sum-of-k-mirror-numbers/description/?envType=daily-question&envId=2025-06-23

/*
1, 生成[2,9]进制以及在10进制下 k镜像数的所有可能
2, 最直接生成累加求结果
3，现生成10进制的回文数
	[base,10*base) 范围内不动最后一位，前面反转，然后拼接生成奇数个的回文数组
	[base,10*base) 将所有数组反装，然后拼接到后面生成偶数个的回文数

4，判断是否为k进制的回文
*/

const maxN = 30

var ans [10][]int

// 力扣 9. 回文数
func isKPalindrome(x, k int) bool {
	if x%k == 0 {
		return false
	}
	rev := 0
	for rev < x/k {
		rev = rev*k + x%k
		x /= k
	}
	return rev == x || rev == x/k
}

func doPalindrome(x int) bool {
	done := true
	for k := 2; k < 10; k++ {
		if len(ans[k]) < maxN && isKPalindrome(x, k) {
			ans[k] = append(ans[k], x)
		}
		if len(ans[k]) < maxN {
			done = false
		}
	}
	if !done {
		return false
	}

	for k := 2; k < 10; k++ {
		// 计算前缀和
		for i := 1; i < maxN; i++ {
			ans[k][i] += ans[k][i-1]
		}
	}
	return true
}

func init() {
	for k := 2; k < 10; k++ {
		ans[k] = make([]int, 0, maxN) // 预分配空间
	}
	for base := 1; ; base *= 10 {
		// 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			if doPalindrome(x) {
				return
			}
		}
		// 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
		for i := base; i < base*10; i++ {
			x := i
			for t := i; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			if doPalindrome(x) {
				return
			}
		}
	}
}

func kMirror(k, n int) int64 {
	return int64(ans[k][n-1])
}
