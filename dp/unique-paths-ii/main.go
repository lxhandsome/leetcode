package main

// https://leetcode.cn/problems/unique-paths-ii/description/

/*
思路：
dp

定义状态： f[i][j] 为到达ij不同路径总和
状态转移方程:

	当前是障碍物的话 obstacleGrid[i][j] == 1 f[i][j] = 0
	当前不是障碍物的话 obstacleGrid[i][j] == 0  f[i][j] =  f[i-1][j] + f[i][j-1]

初始状态：f[0][0] = 0
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				if obstacleGrid[i][j] == 1 || dp[i][j-1] == 0 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i][j-1]
				}
				continue
			}
			if j == 0 {
				if obstacleGrid[i][j] == 1 || dp[i-1][j] == 0 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i-1][j]
				}
				continue
			}
			if obstacleGrid[i][j] == 1 || (dp[i-1][j] == 0 && dp[i][j-1] == 0) {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func testCase() [][]int {
	return [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
}

func main() {

}
