package main

import "sort"

type DSU struct {
	parent []int
	rank   []int
}

// 初始化并查集
func NewDSU(size int) *DSU {
	dsu := &DSU{
		parent: make([]int, size),
		rank:   make([]int, size),
	}
	for i := range dsu.parent {
		dsu.parent[i] = i
		dsu.rank[i] = 1
	}
	return dsu
}

// 查找根节点，带路径压缩
func (d *DSU) Find(u int) int {
	if d.parent[u] != u {
		d.parent[u] = d.Find(d.parent[u]) // 路径压缩
	}
	return d.parent[u]
}

// 合并两个集合，按秩合并
func (d *DSU) Union(u, v int) {
	rootU := d.Find(u)
	rootV := d.Find(v)
	if rootU == rootV {
		return // 已经在同一集合
	}

	// 按秩合并
	if d.rank[rootU] > d.rank[rootV] {
		d.parent[rootV] = rootU
	} else if d.rank[rootU] < d.rank[rootV] {
		d.parent[rootU] = rootV
	} else {
		d.parent[rootV] = rootU
		d.rank[rootU]++
	}
}

func mergeArrays(arrays [][]int) [][]int {
	n := len(arrays)
	dsu := NewDSU(n)

	// 预处理：将每个数组转为集合，方便计算交集
	sets := make([]map[int]bool, n)
	for i, arr := range arrays {
		sets[i] = make(map[int]bool)
		for _, num := range arr {
			sets[i][num] = true
		}
	}

	// 检查所有数组对是否可以合并
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 计算交集大小
			common := 0
			for num := range sets[i] {
				if sets[j][num] {
					common++
					if common >= 2 {
						break
					}
				}
			}
			if common >= 2 {
				dsu.Union(i, j)
			}
		}
	}

	// 收集每个连通分量的数组索引
	components := make(map[int][]int)
	for i := 0; i < n; i++ {
		root := dsu.Find(i)
		components[root] = append(components[root], i)
	}

	// 合并每个分量的数组
	result := make([][]int, 0, len(components))
	for _, group := range components {
		merged := make(map[int]bool)
		for _, idx := range group {
			for _, num := range arrays[idx] {
				merged[num] = true
			}
		}
		// 转为有序切片
		sorted := make([]int, 0, len(merged))
		for num := range merged {
			sorted = append(sorted, num)
		}
		sort.Ints(sorted)
		result = append(result, sorted)
	}

	return result
}
