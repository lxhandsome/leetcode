package main

// https://leetcode.cn/problems/clone-graph/description/
type Node struct {
	Val       int
	Neighbors []*Node
}

// 思路
// 遍历图,对于遍历过的节点返回新建好的节点,对于没有遍历过的节点新建节点,并copy他的邻居节点
func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if _, ok := visited[node]; ok {
		return visited[node]
	}
	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return newNode
}
