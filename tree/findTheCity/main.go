package main

func minimumFuelCost(roads [][]int, seats int) int64 {
	n := 100010
	m := 2 * n
	idx := 0

	var ans int64
	header := make([]int, n)
	for i := range header {
		header[i] = -1
	}
	edge := make([]int, m)
	next := make([]int, m)

	add := func(start, end int) {
		edge[idx] = end
		next[idx] = header[start]
		header[start] = idx
		idx++
	}
	for _, road := range roads {
		add(road[0], road[1])
		add(road[1], road[0])
	}
	var dfs func(u, fa, t int) int
	dfs = func(u, fa, t int) int {
		cnt := 1
		for i := header[u]; i != -1; i = next[i] {
			j := edge[i]
			if j == fa {
				continue
			}
			cnt += dfs(j, u, t)
		}
		if u != 0 {
			ans += int64(cnt * 1.0 / t)
		}
		return cnt
	}
	dfs(0, -1, seats)
	return ans
}

func makeTestCase() [][]int {
	return [][]int{
		{3, 1},
		{3, 2},
		{1, 0},
		{0, 4},
		{0, 5},
		{4, 6},
	}
}

func main() {
	minimumFuelCost(makeTestCase(), 2)
}
