package main

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ld := maxDepth(root.Left) + 1
	rd := maxDepth(root.Right) + 1
	return max(ld, rd)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
