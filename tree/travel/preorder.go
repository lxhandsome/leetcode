package main

func recursionPreOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	recursionPreOrder(root.Left, res)
	recursionPreOrder(root.Right, res)
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || root != nil {
		for root != nil {
			// 先打印根节点
			res = append(res, root.Val)
			stack = append(stack, root)
			// 向左走
			root = root.Left
		}
		// 走到最左边了
		// 从栈里pop一个然后向右边走
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 当前tmp的左节点和根节点已经都被访问过了
		root = tmp.Right
	}
	return res
}
