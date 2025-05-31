package main

import "fmt"

// https://leetcode.cn/problems/snakes-and-ladders/description/?envType=daily-question&envId=2025-05-31

func main() {
	res := snakesAndLadders([][]int{
		{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1},
	})
	fmt.Println(res)
}

func snakesAndLadders(board [][]int) int {
	queue := []int{1}
	n := len(board)
	visit := make([]bool, n*n+1)
	visit[1] = true
	step := 0
	for len(queue) != 0 {
		tmp := queue
		queue = make([]int, 0)
		for _, l := range tmp {
			if l == n*n {
				return step
			}
			start := l + 1
			end := min(l+6, n*n)
			for i := start; i <= end; i++ {
				x := (i - 1) / n
				y := (i - 1) % n
				if x%2 > 0 {
					y = n - y - 1
				}
				next := board[n-1-x][y]
				if next == -1 {
					next = i
				}
				if !visit[next] {
					visit[next] = true
					queue = append(queue, next)
				}
			}
		}
		step++
	}

	return -1

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
