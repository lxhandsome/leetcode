package main

import (
	"container/heap"

)

// https://leetcode.cn/problems/minimum-difference-in-sums-after-removal-of-elements/description/?envType=daily-question&envId=2025-07-18

var _ heap.Interface = (*minHeap)(nil)
var _ heap.Interface = (*maxHeap)(nil)

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

// Swap implements heap.Interface.
func (m *minHeap) Swap(i int, j int) {
	m.array[i], m.array[j] = m.array[j], m.array[i]
}

// Pop implements heap.Interface.
func (m *minHeap) Pop() any {
	l := len(m.array)
	top := m.array[l-1]
	m.array = m.array[:l-1]
	return top
}

// Push implements heap.Interface.
func (m *minHeap) Push(x any) {
	m.array = append(m.array, x.(int))
}

type maxHeap struct {
	array []int
}

// Len implements heap.Interface.
func (m *maxHeap) Len() int {
	return len(m.array)
}

// Less implements heap.Interface.
func (m *maxHeap) Less(i int, j int) bool {
	return m.array[j] < m.array[i]
}

// Pop implements heap.Interface.
func (m *maxHeap) Pop() any {
	l := len(m.array)
	top := m.array[l-1]
	m.array = m.array[:l-1]
	return top
}

// Push implements heap.Interface.
func (m *maxHeap) Push(x any) {
	m.array = append(m.array, x.(int))
}

// Swap implements heap.Interface.
func (m *maxHeap) Swap(i int, j int) {
	m.array[i], m.array[j] = m.array[j], m.array[i]
}

/*
1. 数组从i下标进行分割 [0,i] 中选择n个使得和最小
2. [i+1,3n-1]选择n个使的和最大
3. 最小前缀后 - 最大后缀和 = 最小值
*/
func minimumDifference(nums []int) int64 {
	// 求最大后缀
	m := len(nums)
	n := m / 3
	mh := &minHeap{array: nums[2*n:]}
	sum := 0
	heap.Init(mh)
	for _, num := range mh.array {
		sum += num
	}
	suffixMax := make([]int, n*2+1)
	suffixMax[2*n] = sum
	for i := 2*n - 1; i >= n; i-- {
		if nums[i] > mh.array[0] {
			sum -= heap.Pop(mh).(int)
			heap.Push(mh, nums[i])
			sum += nums[i]
		}
		suffixMax[i] = sum
	}

	maxh := &maxHeap{array: nums[:n]}
	heap.Init(maxh)

	prefixSum := 0

	for _, n := range maxh.array {
		prefixSum += n
	}

	ans := prefixSum - suffixMax[n]

	for i := n; i < 2*n; i++ {
		if nums[i] < maxh.array[0] {
			prefixSum -= heap.Pop(maxh).(int)
			heap.Push(mh, nums[i])
			prefixSum += nums[i]
		}
		ans = min(ans, prefixSum-suffixMax[i+1])
	}

	return int64(ans)

}
