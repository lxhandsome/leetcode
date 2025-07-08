package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended/solutions/3707151/liang-chong-fang-fa-zui-xiao-dui-bing-ch-ijbf/?envType=daily-question&envId=2025-07-07

/*

以开始时间相同对会议进行分组，组里放置的是会议的结束时间

1，针对第i天，有会议安排[i,j],第i天必须参加这场会议没如果不参加，有可能会错过这场会议e.g [i,i]
2, 以第i天为开始时间的会议会议安排可能有多场次 [i,j] [i,j+1] ... ，应该参加结束时间最早的会议，如果解释参加晚点的会议那么有可能错过更多的会议


*/

var _ heap.Interface = (*minHeap)(nil)

type minHeap struct {
	array []int
}

// Len implements heap.Interface.
func (m *minHeap) Len() int {
	return len(m.array)
}

// Less implements heap.Interface.
func (m *minHeap) Less(i int, j int) bool {
	return m.array[i] < m.array[j]
}

// Pop implements heap.Interface.
func (m *minHeap) Pop() any {
	end := m.array[m.Len()-1]
	m.array = m.array[:m.Len()-1]
	return end
}

func (m *minHeap) Top() int {
	top := m.array[0]
	return top
}

// Push implements heap.Interface.
func (m *minHeap) Push(x any) {
	m.array = append(m.array, x.(int))
}

// Swap implements heap.Interface.
func (m *minHeap) Swap(i int, j int) {
	m.array[i], m.array[j] = m.array[j], m.array[i]
}

func maxEvents(events [][]int) int {

	maxEnd := 0
	// 找到会议的最大结束时间
	for _, event := range events {
		maxEnd = max(maxEnd, event[1])
	}
	group := make([][]int, maxEnd+1)
	// 使用开始时间对会议进行分组
	for _, event := range events {
		group[event[0]] = append(group[event[0]], event[1])
	}

	h := &minHeap{}

	heap.Init(h)

	var res int

	for start, eventEnd := range group {
		// 将堆中小于当前时间的会议删除
		for h.Len() > 0 && h.Top() < start {
			heap.Pop(h)
		}

		for _, end := range eventEnd {
			heap.Push(h, end)
		}
		if h.Len() > 0 {
			res++
			heap.Pop(h)
		}
	}

	return res
}

func main() {
	res := maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}})
	fmt.Println(res)
}
