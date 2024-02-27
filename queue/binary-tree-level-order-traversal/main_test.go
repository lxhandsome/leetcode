package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newNode(left, right *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	node.Left = left
	node.Right = right
	return node
}

func makeTestCase() (*TreeNode, [][]int) {
	res := [][]int{
		{3},
		{9, 20},
		{15, 7},
	}
	node15 := newNode(nil, nil, 15)
	node7 := newNode(nil, nil, 7)
	node20 := newNode(node15, node7, 20)
	node9 := newNode(nil, nil, 9)
	node3 := newNode(node9, node20, 3)
	return node3, res
}

func TestLevelOrder(t *testing.T) {
	node, expect := makeTestCase()
	actual := levelOrder(node)
	assert.Equal(t, len(expect), len(actual), "层数不相同")
	for i := 0; i < len(expect); i++ {
		assert.Equal(t, len(expect[i]), len(actual[i]), fmt.Sprintf("第 %d 层长度不同", i+1))
		for j := 0; j < len(expect[i]); j++ {
			assert.Equal(t, expect[i][j], actual[i][j], fmt.Sprintf("第 %d 层 第 %d 个元素不同", i+1, j+1))
		}
	}
}
