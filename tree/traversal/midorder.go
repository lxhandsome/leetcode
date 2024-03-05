package main

func recursionMidOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	recursionMidOrder(root.Left, res)
	*res = append(*res, root.Val)
	recursionMidOrder(root.Right, res)
}

func midOrder(root *TreeNode, res *[]int) {
	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		*res = append(*res, tmp.Val)
		root = tmp.Right
	}
}
