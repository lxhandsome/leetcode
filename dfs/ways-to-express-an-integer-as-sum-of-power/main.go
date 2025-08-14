package main

import "fmt"

func main() {
	fmt.Printf("numberOfWays(10, 2): %v\n", numberOfWays(4, 1))
}

func fastPower(a, b int) int {
	result := 1
	for b > 0 {
		if b%2 == 1 {
			result *= a
		}
		a *= a
		b /= 2
	}
	return result
}

func numberOfWays(n int, x int) (res int) {

	var pows [301]int
	const mod = 1e9 + 7

	var dfs func(num int, sum int)

	dfs = func(num int, sum int) {
		if num > n {
			return
		}
		if sum > n {
			return
		}
		if sum == n {
			res = (res % mod) + 1
			return
		}
		if pows[num] == 0 {
			pows[num] = fastPower(num, x)
		}
		dfs(num+1, sum+pows[num])
		dfs(num+1, sum)
	}

	dfs(1, 0)

	return res

}
