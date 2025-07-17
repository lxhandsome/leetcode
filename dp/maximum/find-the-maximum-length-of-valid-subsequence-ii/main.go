package main

// https://leetcode.cn/problems/find-the-maximum-length-of-valid-subsequence-ii/?envType=daily-question&envId=2025-07-17
/*
1. *(sub[0] + sub[1]) mod k == (sub[1] + sub[2]) mod k  == > (sub[2] - sub[0]) mod k == 0

子序列同相同奇偶下表的元素mod k应该相同

考虑子序列任意相邻的两个元素 a，b mod k == > a1,b1 子序列下一个元素肯定为 a1
记f[x][y] 为 mod 结果为以x开头，以y结尾对的子序列的长度
便利nums ==> x


*/
func maximumLength(nums []int, k int) (res int) {
	f := make([][]int, k)
	for i := range f {
		f[i] = make([]int, k)
	}
	for _, x := range nums {
		x %= k
		for y, fxy := range f[x] {
			f[y][x] = fxy + 1
			res = max(res, f[y][x])
		}
	}
	return res
}
