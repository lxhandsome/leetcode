package main

type segTree struct {
	max []int
}

func newSegTree(a []int) *segTree {
	n := len(a)
	tree := &segTree{
		max: make([]int, n*4),
	}
	tree.build(a, 1, 0, n-1)
	return tree
}

func (t *segTree) build(baskets []int, o, l, r int) {
	if l == r {
		t.max[o] = baskets[l]
		return
	}
	m := (l + r) >> 1
	t.build(baskets, 2*o, l, m)
	t.build(baskets, 2*o+1, m+1, r)
	t.maintain(o)
}

func (t *segTree) maintain(o int) {
	t.max[o] = max(t.max[2*o], t.max[2*o+1])
}

func (t *segTree) findFristAndUpdate(o int, x int, l, r int) int {
	if t.max[o] < x {
		return -1
	}
	if l == r {
		t.max[o] = -1
		return l
	}
	m := (l + r) >> 1
	i := t.findFristAndUpdate(2*o, x, l, m)
	if i < 0 {
		i = t.findFristAndUpdate(2*o+1, x, m+1, r)
	}
	t.maintain(o)
	return i

}

func numOfUnplacedFruits(fruits []int, baskets []int) (res int) {
	segTree := newSegTree(baskets)
	n := len(fruits)
	for _, f := range fruits {
		if segTree.findFristAndUpdate(1, f, 0, n-1) < 0 {
			res++
		}
	}
	return
}
