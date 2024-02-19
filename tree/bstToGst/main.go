package main

// 1038. 从二叉搜索树到更大和树
// https://leetcode.cn/problems/binary-search-tree-to-greater-sum-tree/description/

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	total := 0
	var dfs func(root *TreeNode)

	// 先遍历二叉搜索树，获得该节点对应的右子树的和
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		total += root.Val
		root.Val = total
		dfs(root.Left)
	}
	dfs(root)
	return root
}

func newNode(left, right *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	node.Left = left
	node.Right = right
	return node
}

func makeTestCase() *TreeNode {
	node8 := newNode(nil, nil, 8)
	node7 := newNode(nil, node8, 7)
	node5 := newNode(nil, nil, 5)
	node6 := newNode(node5, node7, 6)
	node3 := newNode(nil, nil, 3)
	node2 := newNode(nil, node3, 2)
	node0 := newNode(nil, nil, 0)
	node1 := newNode(node0, node2, 1)
	node4 := newNode(node1, node6, 4)
	return node4
}

func main() {
	root := makeTestCase()
	res := bstToGst(root)
	fmt.Println(res)
}
