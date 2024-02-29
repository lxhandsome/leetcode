package main

// https://leetcode.cn/problems/01-matrix/description/
/*
思路:
计算每个格子到对应理她最近的0的距离
可转化为从0开始bfs,计算到达非0位置所需要的距离,并且标记当前位置已经标记   当前位置到离他最近的0的距离等于 前一个节点到0的距离 + 1
*/
func updateMatrix(mat [][]int) [][]int {
	queue := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			// 从0出发，所以需要将所有的0入队
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				// 标记当前坐标没有被访问过
				mat[i][j] = -1
			}
		}
	}
	// 需要向当前点的四个方向出发，坐标还没有被访问过，计算当前点到最近的0的距离
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]
		for _, d := range directions {
			x := front[0] + d[0]
			y := front[1] + d[1]
			// 当前节点没有被访问过
			if x >= 0 && x < len(mat) && y >= 0 && y < len(mat[x]) && mat[x][y] == -1 {
				// 计算当前节点离最近0的距离
				mat[x][y] = mat[front[0]][front[1]] + 1
				// 从当前位置出发可能还有节点没有被访问，所以需要将当前节点入队
				queue = append(queue, []int{x, y})
			}
		}
	}
	return mat
}
