package main

import "math"

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路：
最大路径和的可能有三种情况
1，左子树的最大路径和
2，右子树的最大路径和
3，左子树提供的最大路径 + 根节点值 + 右子树提供的最大路径和
*/

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	var dfs func(r *TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		left := dfs(r.Left)
		right := dfs(r.Right)
		fullMax := left + right + r.Val
		maxSum = max(maxSum, fullMax)
		return max(max(left, right)+r.Val, 0)
	}
	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
