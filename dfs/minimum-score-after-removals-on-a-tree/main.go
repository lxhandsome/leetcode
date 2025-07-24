package main

import "math"

// https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree/solutions/1625899/dfs-shi-jian-chuo-chu-li-shu-shang-wen-t-x1kk/?envType=daily-question&envId=2025-07-24
/*

1. x[i] 记为以i节点为根子树的异或值

2. 题目要求删除两条边，生成三个通联分量，使得 分数 = 最大联通分量异或值 - 最小连通分量分量的亦或值
	求删除方案最小的分数值

	假设删除的边分别为 [x1,y1] [x2,y2]

3. 可能有下面几种情况

	a. 两条边都在同一棵子树上并且 y1,是x2的父节点
	此时 三个连通分量的异或值为 x[0] yh x[x1]  &&  x[x1] yh x[x2] && x[y2]

	b. 两条边都在同一个子树上并且 y2,是x1的父节点
	此时 三个连通分量的异或值为 x[0] yh x[x2]  &&  x[x2] yh x[y2] && x[x1]


	c. 两条边在两个互不相交的子树上
	此时 三个连通分量的异或值为 x[0] yh x[x1] yh x[x2] && x[x1] && x[x2]

4. 如何判断是两条边式在同个子树还是不同子树上，同一个子树，如何判断谁是父节点属实子节点
	引入时间戳的概念
	对于每个节点有in[x] out[x]代表进入节点的时间以及出节点的时间

	如果一个x1 ，是x2的父节点

	in[x] < in[x2] <= out[x2] <= out[x1]

	如果两个是互不相交的子树中

	in[x1] out[x1] 与 in[x2] out[x2]互不相交



*/

func minimumScore(nums []int, edges [][]int) int {

	// 计算每个节点的异或值

	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		s := e[0]
		e := e[1]
		g[s] = append(g[s], e)
		g[e] = append(g[e], s)
	}
	xor := make([]int, n)
	in := make([]int, n)
	out := make([]int, n)
	clock := 0

	var dfs func(x int, f int)

	dfs = func(x int, f int) {
		clock++
		in[x] = clock
		xor[x] = nums[x]
		for _, y := range g[x] {
			if y != f {
				dfs(y, x)
				xor[x] ^= xor[y]
			}
		}
		out[x] = clock
	}

	dfs(0, -1)

	// 判断 x 是否为 y 的祖先
	isAncestor := func(x, y int) bool {
		return in[x] < in[y] && in[y] <= out[x]
	}

	ans := math.MaxInt
	// 枚举：删除 x 与 x 父节点之间的边，删除 y 与 y 父节点之间的边
	for x := 2; x < n; x++ {
		for y := 1; y < x; y++ {
			var a, b, c int
			if isAncestor(x, y) { // x 是 y 的祖先
				a, b, c = xor[y], xor[x]^xor[y], xor[0]^xor[x]
			} else if isAncestor(y, x) { // y 是 x 的祖先
				a, b, c = xor[x], xor[x]^xor[y], xor[0]^xor[y]
			} else { // x 和 y 分别属于两棵不相交的子树
				a, b, c = xor[x], xor[y], xor[0]^xor[x]^xor[y]
			}
			ans = min(ans, max(a, b, c)-min(a, b, c))
			if ans == 0 { // 不可能变小
				return 0 // 提前返回
			}
		}
	}

	return ans
}
