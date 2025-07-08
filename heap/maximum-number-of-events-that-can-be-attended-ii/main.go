package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended-ii/description/?envType=daily-question&envId=2025-07-08

var _ heap.Interface = (*minHeap)(nil)

type dayValue struct {
	day   int
	value int
}

type minHeap struct {
	array []dayValue
}

// Len implements heap.Interface.
func (m *minHeap) Len() int {
	return len(m.array)
}

// Less implements heap.Interface.
func (m *minHeap) Less(i int, j int) bool {
	if m.array[i].value != m.array[i].value {
		return m.array[i].value > m.array[i].value
	}
	return m.array[i].day < m.array[i].day
}

// Pop implements heap.Interface.
func (m *minHeap) Pop() any {
	end := m.array[m.Len()-1]
	m.array = m.array[:m.Len()-1]
	return end
}

func (m *minHeap) Peek() dayValue {
	top := m.array[0]
	return top
}

// Push implements heap.Interface.
func (m *minHeap) Push(x any) {
	m.array = append(m.array, x.(dayValue))
}

// Swap implements heap.Interface.
func (m *minHeap) Swap(i int, j int) {
	m.array[i], m.array[j] = m.array[j], m.array[i]
}

func maxValue(events [][]int, k int) int {
	// 找到会议最大结束时间
	maxEnd := 0
	for _, event := range events {
		maxEnd = max(maxEnd, event[1])
	}
	// 使用开始时间进行分组
	group := make([][]dayValue, maxEnd+1)
	for _, event := range events {
		group[event[0]] = append(group[event[0]], dayValue{day: event[1], value: event[2]})
	}
	visited := make([]bool, len(group))
	h := &minHeap{}
	res := 0
	for day, eventEnd := range group {
		if k <= 0 {
			break
		}
		for h.Len() > 0 && h.Peek().day < day {
			heap.Pop(h)
		}
		for _, event := range eventEnd {
			heap.Push(h, event)
		}

		if h.Len() > 0 {
			top := h.Peek()
			if !visited[top.day] && k > 0 {
				heap.Pop(h)
				visited[top.day] = true
				res += top.value
				k--
			}
		}

	}

	return res
}

func main() {
	res := maxValue([][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}, 2)
	fmt.Println(res)
}
