package main

import "fmt"

// https://leetcode.cn/problems/divisible-and-non-divisible-sums-difference/?envType=daily-question&envId=2025-05-27
func differenceOfSums(n int, m int) int {
	var num1, num2 int
	for i := 1; i <= n; i++ {
		if i%m == 0 {
			num2 += i
		} else {
			num1 += i
		}
	}
	return num1 - num2
}

func main() {
	res := numOfUnplacedFruits([]int{4, 2, 5}, []int{3, 5, 4})
	fmt.Println(res)
}

func numOfUnplacedFruits(fruits []int, baskets []int) (res int) {
	n := len(baskets)
	flag := make([]bool, n)
	for i := 0; i < n; i++ {
		f := false
		for j := 0; j < n; j++ {
			if baskets[j] >= fruits[i] && !flag[j] {
				flag[j] = true
				f = true
				break
			}
		}
		if !f {
			res++
		}
	}
	return res
}
