package main

import "fmt"

func main() {
	fmt.Printf("checkPowersOfThree(12): %v\n", checkPowersOfThree(12))
}

func checkPowersOfThree(n int) bool {

	var dfs func(e int, current int) bool

	dfs = func(e int, current int) bool {
		if current < 0 {
			return false
		}
		if current == 0 {
			return true
		}
		e3 := fastPow(3, e)
		if e3 > current {
			return false
		}
		if dfs(e+1, current-e3) {
			return true
		}
		return dfs(e+1, current)
	}

	return dfs(0, n)
}

/*
一个三进制数转换为10进制 (210)3

0*3^0 + 1*3^1 + 2*3^2

可以看出如果n的三进制如果有某一位为2则 那么肯定有相同的两个数相加
*/
func checkPowersOfThree1(n int) bool {
	for ; n > 0; n /= 3 {
		if n%3 == 2 {
			return false
		}
	}
	return true
}

func fastPow(a, b int) int {
	res := 1
	for b > 0 {
		if b%2 == 1 {
			res *= a
		}
		a *= a
		b /= 2
	}
	return res
}
