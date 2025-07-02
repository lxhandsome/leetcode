package main

// https://leetcode.cn/problems/find-the-original-typed-string-ii/description/?envType=daily-question&envId=2025-07-02

//题解： https://leetcode.cn/problems/find-the-original-typed-string-ii/solutions/2966856/zheng-nan-ze-fan-qian-zhui-he-you-hua-dp-5mi9/?envType=daily-question&envId=2025-07-02
/*
1，计算不考虑输入长度为k的可能情况 total
2，计算输入长度小于k的可能情况 minK 有m组数字，每组中至少选择一个数字，求在剩余可选数字中选择 k-1-m数字的可能方案 （多重背包问题）
	多重背包问题解题思路：定义 f[i][j] 表示从前 i 种物品中选至多 j 个物品的方案数。

3，使用总的 total - minK 及时结果

*/

func possibleStringCount(word string, k int) int {
	if len(word) < k { // 无法满足要求
		return 0
	}

	const mod = 1_000_000_007
	cnts := []int{}
	ans := 1
	cnt := 0
	for i := range word {
		cnt++
		if i == len(word)-1 || word[i] != word[i+1] {
			// 如果 cnt = 1，这组字符串必选，无需参与计算
			if cnt > 1 {
				if k > 0 {
					cnts = append(cnts, cnt-1)
				}
				ans = ans * cnt % mod
			}
			k-- // 注意这里把 k 减小了
			cnt = 0
		}
	}

	if k <= 0 {
		return ans
	}

	m := len(cnts)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, k)
	}
	for i := range f[0] {
		f[0][i] = 1
	}

	s := make([]int, k+1)
	for i, c := range cnts {
		// 计算 f[i] 的前缀和数组 s
		for j, v := range f[i] {
			s[j+1] = s[j] + v
		}
		// 计算子数组和
		for j := range f[i+1] {
			f[i+1][j] = (s[j+1] - s[max(j-c, 0)]) % mod
		}
	}

	return (ans - f[m][k-1] + mod) % mod // 保证结果非负
}
