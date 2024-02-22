package main

// https://leetcode.cn/problems/number-of-islands/description/
// 思路: 使用深度优先搜索，将1的上下左右的1都遍历完才算找到一块陆地
// 一个岛屿：当且仅当它的上下左右都是0时才是岛屿，当踩到岛屿的时候，将当前踩到的陆地置为0继续遍历，当1的前后左右都是0是岛屿搜寻结束，返回1(代表一块陆地)
// 当前graph[i][j] == '1'  graph[i-1][j] graph[i+1][j] graph[i][j-1] graph[i][j+1]
func numIslands(grid [][]byte) (count int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && dfs(grid, i, j) > 0 {
				count++
			}
		}
	}
	return
}

func dfs(grid [][]byte, i, j int) int {
	// 超出边界 陆地数量为0
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		return dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j-1) + dfs(grid, i, j+1) + 1
	}
	// 当前为水陆地数量为0
	return 0
}
