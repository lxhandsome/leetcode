package main

import (
	"fmt"
)

func productQueries(n int, queries [][]int) []int {
	const mod = 1_000_000_007
	powers := []int{}
	for n > 0 {
		lowbit := n & -n
		powers = append(powers, lowbit)
		n ^= lowbit
	}

	m := len(powers)
	res := make([][]int, m)
	for i, x := range powers {
		res[i] = make([]int, m)
		res[i][i] = x
		for j := i + 1; j < m; j++ {
			res[i][j] = res[i][j-1] * powers[j] % mod
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = res[q[0]][q[1]]
	}
	return ans
}
func main() {
	queries := [][]int{{0, 1}, {2, 2}, {0, 3}}

	n := 15

	fmt.Println(productQueries(n, queries))
}
