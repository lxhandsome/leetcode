package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
// 有两种方式实现
// 第一种使用递归,来实现中序遍历
func inorderTraversal1(root *TreeNode) []int {
	res := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

// 第二种使用stack来实现中序遍历
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		root = top.Right
	}
	return res
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
	inorderTraversal(makeTestCase())
}
