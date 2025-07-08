package main

import (
	"slices"
	"sort"
)

func maxValue(events [][]int, k int) int {
	// 特判 k=1 的情况可以更快
	if k == 1 {
		mx := 0
		for _, e := range events {
			mx = max(mx, e[2])
		}
		return mx
	}

	slices.SortFunc(events, func(a, b []int) int { return a[1] - b[1] })
	n := len(events)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}
	for i, e := range events {
		startDay, value := e[0], e[2]
		p := sort.Search(i, func(j int) bool { return events[j][1] >= startDay })
		for j := 1; j <= k; j++ {
			// 为什么是 p 不是 p+1：上面算的是 >= startDay，-1 后得到 < startDay，但由于还要 +1，抵消了
			f[i+1][j] = max(f[i][j], f[p][j-1]+value)
		}
	}
	return f[n][k]
}
