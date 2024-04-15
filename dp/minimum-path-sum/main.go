package main

// https://leetcode.cn/problems/minimum-path-sum/description/

/*
思路：动态规划
定义状态： f[i][j] 从00出发到达ij所走的最短路径
初始条件 f[0][0] = a[0][0]
循环不变式：f[i][j] = min(f[i-1][j],f[i][j-1]) + a[i][j]
结果 f[n-1][m-1]
*/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}
	n := len(grid)
	m := len(grid[0])
	res := make([][]int, 2)
	// 初始化
	res[0] = make([]int, m)
	res[1] = make([]int, m)
	res[0][0] = grid[0][0]
	for j := 1; j < m; j++ {
		res[0][j] = res[0][j-1] + grid[0][j]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 0 {
				res[i%2][j] = res[(i-1)%2][j] + grid[i][j]
				continue
			}
			res[i%2][j] = min(res[(i-1)%2][j], res[i%2][j-1]) + grid[i][j]
		}
	}
	return res[(n-1)%2][m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
