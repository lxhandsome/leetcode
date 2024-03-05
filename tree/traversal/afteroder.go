package main

func recursionAfterOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	recursionAfterOrder(root.Left, res)
	recursionAfterOrder(root.Right, res)
	*res = append(*res, root.Val)
}

// 后序遍历非递归的关键是根节点要在右节点弹出之后再弹出，所以需要记录右节点
func afterOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	var last *TreeNode
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 因为此时node的右边节点还没弹出，所以不能弹出root，继续遍历
		node := stack[len(stack)-1]
		if node.Right == last || node.Right == nil {
			stack = stack[:len(stack)-1]
			*res = append(*res, node.Val)
			last = node
		} else {
			root = node.Right
		}

	}
}
