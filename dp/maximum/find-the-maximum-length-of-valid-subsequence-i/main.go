package main

import "fmt"

// https://leetcode.cn/problems/find-the-maximum-length-of-valid-subsequence-i/description/?envType=daily-question&envId=2025-07-16

/*
1. f[i][j] 从i - j最大有效子序列的长度

2. f[i][j] =
	if (nums[i] + nums[i+1]) %2 == (nums[j] + nums[j-1] ) %2 f[i][j] = f[i+1][j-1]
	else f[i][j] = f[i-1][j+1]


3. init

	f[i-1][i] = 2


*/

func main() {
	fmt.Printf("maximumLength([]int{1, 2, 3, 4}): %v\n", maximumLength([]int{1, 2, 1, 1, 2, 1, 2}))
}

func maximumLength1(nums []int) int {

	n := len(nums)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if i == j || i-j == 1 {
				f[j][i] = 1
				continue
			}
			f[j][i] = max(f[j+1][i], f[j][i-1])
			p2r := (nums[i-1] + nums[i]) % 2
			p2l := (nums[j] + nums[j+1]) % 2
			if p2l == p2r {
				m := f[j+1][i-1] + 2
				if i-j == 3 && (nums[j+1]+nums[i-1])%2 == p2l {
					m += 1
				}
				f[j][i] = max(m, f[j][i])
			}
		}
	}
	return f[0][n-1]
}


/*
根据 模运算的世界：当加减乘除遇上取模，可以移项，得

(a+b−(b+c))modk=0
化简得

(a−c)modk=0
这意味着 a 与 c 关于模 k 同余。即题目式子中的 sub[i] 与 sub[i+2] 关于模 k 同余。换句话说，有效子序列的偶数项 sub[0],sub[2],sub[4],… 都关于模 k 同余，奇数项 sub[1],sub[3],sub[5],… 都关于模 k 同余。

如果把每个 nums[i] 都改成 nums[i]modk，问题等价于：

求最长子序列的长度，该子序列的奇数项都相同，偶数项都相同。

比如 nums=[1,2,1,2,1,2]：

遍历到 1 的时候，在「末尾为 1,2 的子序列」的末尾添加一个 1，得到「末尾为 2,1 的子序列」。
遍历到 2 的时候，在「末尾为 2,1 的子序列」的末尾添加一个 2，得到「末尾为 1,2 的子序列」。



*/

func maximumLength(nums []int) int {
	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, 2)
	}
	res := 0
	for _, x := range nums {
		x %= 2
		for y, fxy := range f[x] {
			f[y][x] = fxy + 1
			res = max(res, f[y][x])
		}
	}
	return res
}
