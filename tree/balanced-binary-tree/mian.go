package main

// https://leetcode.cn/problems/balanced-binary-tree/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路:
高度平衡的二叉树:

当前节点左孩子的高度 - 右孩子的高度 < 1 && 左孩子是高度平衡的二叉树 && 右孩子是高度平衡的二叉树
*/

// 计算树的深度，然后判断平衡
func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftB := isBalanced1(root.Left)
	rightB := isBalanced1(root.Right)
	ld := depth(root.Left)
	rd := depth(root.Right)
	return leftB && rightB && abs(ld-rd) <= 1
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return depth(root.Right) + 1
	}
	if root.Right == nil {
		return depth(root.Left) + 1
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dept(root) != -1
}

func dept(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dept(root.Left)
	// 当左子树不平衡
	if l == -1 {
		return -1
	}
	r := dept(root.Right)
	// 右子树不平衡
	if r == -1 {
		return -1
	}
	// 当前节点不平衡返回-1
	if abs(l-r) > 1 {
		return -1
	}
	// 左右节点都平衡, 当前节点也平衡，返回树的深度
	return max(l, r) + 1
}
