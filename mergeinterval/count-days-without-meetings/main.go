package main

import (
	"fmt"
	"slices"

)

// https://leetcode.cn/problems/count-days-without-meetings/?envType=daily-question&envId=2025-07-11

/*
1. 合并区间，使用总时间减去每段区间的长度

*/

func main() {
	res := countDays(10, [][]int{{5, 7}, {1, 3}, {9, 10}})
	fmt.Println(res)
}

func countDays(days int, meetings [][]int) int {
	slices.SortFunc(meetings, func(a, b []int) int {
		return a[0] - b[0]
	})

	// 用来上一段区间的开始和结束时间
	start, end := 1, 0
	for _, m := range meetings {
		s := m[0]
		e := m[1]
		if s > end {
			// 上一段时间已经结束，减去会议时间
			days -= end - start + 1
			start = s
		}
		end = max(end, e)
	}
	days -= end - start + 1
	return days
}
