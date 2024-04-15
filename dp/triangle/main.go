package main

// https://leetcode.cn/problems/triangle/

/*
思路：使用动态规划

自底向上
定义状态： f[i][j] 是从[j][j]出发到达地段的最短路径
初始条件: f[n-1][i] = a[i][j]
循环不变式： f[i][j] = min(f[i+1][j],f[i+1][j+1) + a[i][j]
结果：f[0][0]
*/

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	n := len(triangle)
	res := make([][]int, 2)
	res[0] = make([]int, n)
	res[1] = make([]int, n)
	for i, d := range triangle[n-1] {
		res[(n-1)%2][i] = d
	}
	for i := n - 2; i >= 0; i-- {
		// 计算从当前节点出发到达最后一层的最短路径
		for j, d := range triangle[i] {
			res[i%2][j] = min(res[(i+1)%2][j], res[(i+1)%2][j+1]) + d
		}
	}
	return res[0][0]
}

func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	n := len(triangle)
	res := make([][]int, n)
	res[n-1] = make([]int, len(triangle[n-1]))
	for i, d := range triangle[n-1] {
		res[n-1][i] = d
	}
	for i := n - 2; i >= 0; i-- {
		res[i] = make([]int, len(triangle[i]))
		// 计算从当前节点出发到达最后一层的最短路径
		for j, d := range triangle[i] {
			res[i][j] = min(res[i+1][j], res[i+1][j+1]) + d
		}
	}
	return res[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
