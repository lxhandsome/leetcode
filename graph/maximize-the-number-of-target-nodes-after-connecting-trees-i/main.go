package main

import "fmt"

// https://leetcode.cn/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/?envType=daily-question&envId=2025-05-28

func main() {

	edges1 := [][]int{
		{0, 1},
		{0, 2},
		{2, 3},
		{2, 4},
	}
	edges2 := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{2, 7},
		{1, 4},
		{4, 5},
		{4, 6},
	}

	res := maxTargetNodes(edges1, edges2, 2)

	fmt.Println(res)
}

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {

	max2 := 0
	if k > 0 {
		dfs := buildTree(edges2, k-1)
		for i := 0; i <= len(edges2); i++ {
			max2 = max(max2, dfs(i, -1, 0))
		}
	}

	dfs := buildTree(edges1, k)
	ans := make([]int, len(edges1)+1)
	for i := range ans {
		ans[i] = dfs(i, -1, 0) + max2
	}
	return ans
}

func buildTree(edges [][]int, k int) func(int, int, int) int {

	g := make([][]int, len(edges)+1)

	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int, int) int

	dfs = func(x, from, d int) int {
		if d > k {
			return 0
		}
		cnt := 1
		for _, y := range g[x] {
			if y != from {
				cnt += dfs(y, x, d+1)
			}
		}
		return cnt
	}

	return dfs

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
