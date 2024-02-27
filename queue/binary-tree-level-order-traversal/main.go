package main

// https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路: 使用队列来对树进行层次遍历
使用一个指针保存下一层的最后一个，元素遍历到下一层的最后一个时，更新下一层的最后一个为队列末尾元素
*/
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	nextLevelEnd := root
	currentLevel := 0
	queue = append(queue, root)
	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]
		if len(res) < currentLevel+1 {
			res = append(res, make([]int, 0))
		}
		res[currentLevel] = append(res[currentLevel], front.Val)
		if front.Left != nil {
			queue = append(queue, front.Left)
		}
		if front.Right != nil {
			queue = append(queue, front.Right)
		}
		// 当前层的最后一个
		if front == nextLevelEnd {
			if len(queue) != 0 {
				nextLevelEnd = queue[len(queue)-1]
				currentLevel++
			}
		}
	}
	return res
}
