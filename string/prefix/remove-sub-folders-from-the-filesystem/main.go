package main

import "fmt"

/*字典树，前缀树*/
type node struct {
	folder string
	sub    map[string]*node
	end    bool
}

func newNode(f string) *node {
	return &node{
		folder: f,
		sub:    map[string]*node{},
	}
}

func (n *node) addFolder(s string) {
	if len(s) == 0 {
		n.end = true
		return
	}
	descriptor := 0
	i := 0
	for ; i < len(s); i++ {
		if s[i] == '/' {
			descriptor++
		}
		if descriptor == 2 {
			break
		}

	}
	f := s[:i]
	if _, ok := n.sub[f]; !ok {
		n.sub[f] = newNode(f)
	}
	n.sub[f].addFolder(s[i:])
}

func removeSubfolders(folder []string) (res []string) {
	root := newNode("")
	for _, f := range folder {
		root.addFolder(f)
	}
	var dfs func(*node, string)
	dfs = func(n *node, current string) {
		current += n.folder
		if n.end {
			res = append(res, current)
			return
		}
		for _, node := range n.sub {
			dfs(node, current)
		}
	}

	dfs(root, "")
	return
}

func main() {
	res := removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"})
	fmt.Println(res)
}
