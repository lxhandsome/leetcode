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

var (
	pre_res   = []int{4, 1, 0, 2, 3, 6, 5, 7, 8}
	mid_res   = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	after_res = []int{0, 3, 2, 1, 5, 8, 7, 6, 4}
)

func TestRecursionPreOrder(t *testing.T) {
	actual := make([]int, 0)
	root := makeTestCase()
	recursionPreOrder(root, &actual)
	assert.Equal(t, len(pre_res), len(actual), "长度不一致")
	for i := 0; i < len(pre_res); i++ {
		assert.Equal(t, pre_res[i], actual[i], fmt.Sprintf("第 %d 个元素不相同", i+1))
	}
}

func TestPreOrder(t *testing.T) {
	root := makeTestCase()
	actual := preorderTraversal(root)
	assert.Equal(t, len(pre_res), len(actual), "长度不一致")
	for i := 0; i < len(pre_res); i++ {
		assert.Equal(t, pre_res[i], actual[i], fmt.Sprintf("第 %d 个元素不相同", i+1))
	}
}

func TestRecursionMidOrder(t *testing.T) {
	actual := make([]int, 0)
	root := makeTestCase()
	recursionMidOrder(root, &actual)
	assert.Equal(t, len(mid_res), len(actual), "长度不一致")
	for i := 0; i < len(mid_res); i++ {
		assert.Equal(t, mid_res[i], actual[i], fmt.Sprintf("第 %d 个元素不相同", i+1))
	}
}

func TestMidOrder(t *testing.T) {
	actual := make([]int, 0)
	root := makeTestCase()
	midOrder(root, &actual)
	assert.Equal(t, len(mid_res), len(actual), "长度不一致")
	for i := 0; i < len(mid_res); i++ {
		assert.Equal(t, mid_res[i], actual[i], fmt.Sprintf("第 %d 个元素不相同", i+1))
	}
}

func TestAfterOrder(t *testing.T) {
	actual := make([]int, 0)
	root := makeTestCase()
	afterOrder(root, &actual)
	assert.Equal(t, len(after_res), len(actual), "长度不一致")
	for i := 0; i < len(after_res); i++ {
		assert.Equal(t, after_res[i], actual[i], fmt.Sprintf("第 %d 个元素不相同", i+1))
	}
}

func TestRecursionAfterOrder(t *testing.T) {
	actual := make([]int, 0)
	root := makeTestCase()
	recursionAfterOrder(root, &actual)
	assert.Equal(t, len(after_res), len(actual), "长度不一致")
	for i := 0; i < len(after_res); i++ {
		assert.Equal(t, after_res[i], actual[i], fmt.Sprintf("第 %d 个元素不相同", i+1))
	}
}
