package main

import "math"

// https://leetcode.cn/problems/largest-rectangle-in-histogram/description/
/*
思路:
求以每个矩形为高度的可围成的最大矩形的面积

暴力：遍历每个矩形高举，分别向右向左遍历找到最边上大于当前高度的矩形左右下标，计算出宽度，然后乘以当前高度就是当前高度的最大面积

单调栈：

遍历过程中可以看到当，矩形出现下降后，之前所有大于等于当前矩形的围成的最大面积到可以知道

由此可以使用单调栈，保持栈单调非递减，栈中保存当前元素的下表
如果扫描到的元素小于栈的最大值，那么栈顶这个元素所围成的最大矩形面积可以确定（当前元素的下表 - 栈顶元素 +1 ）* 栈顶元素所记录的高度
依次次出栈，每次都计算某个高度的围成最大矩形，知道所有元素到访问完毕
此时栈中还可能有元素，依次计算剩余元素围成的最大值,右边可以添加一个哨兵，高度为0，用户计算其他未计算最高度所围成的最大面积
*/

func largestRectangleArea(heights []int) int {
	// 用来记录每个高度的下标
	if len(heights) == 1 {
		return heights[0]
	}
	stack := []int{}
	max := math.MinInt
	for i := 0; i <= len(heights); i++ {
		current := -1
		if i != len(heights) {
			current = heights[i]
		}
		// 计算当前最高高度可能围成的最大面积
		for len(stack) > 0 && heights[stack[len(stack)-1]] > current {
			// currentMaxIndex := stack[len(stack)-1]
			// w := i - currentMaxIndex
			// h := heights[currentMaxIndex]
			// max = Max(w*h, max)
			// stack = stack[0 : len(stack)-1]
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 对于h来说当前i就是第一个小于h的右边下标
			h := heights[pop]
			// 计算宽度
			w := i
			if len(stack) != 0 {
				// peek就是左边第一个小于等于h的下标
				peek := stack[len(stack)-1]
				w = i - peek - 1
			}
			max = Max(max, h*w)
		}
		stack = append(stack, i)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
