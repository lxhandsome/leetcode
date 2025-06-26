package main

// https://leetcode.cn/problems/count-the-number-of-arrays-with-k-matching-adjacent-elements/?envType=daily-question&envId=2025-06-17

/*
有n-1对相邻元素相同 ==》 有n-1-k对相邻元素不同

1，先在n-1中选择n -1- k对
2，n -1- k对元素各不相同
3，n -1- k对元素的值在[1,m]之间
4，由于没对相邻元素对元素不相同， c(n-1-k,n-1) * m * (m-1)^(n-1-k)

 c(n-1-k,n-1) = c(k,n-1) =  (n-1)! / k! * (n-1-k)!

*/

import "fmt"

func main() {
	res := countGoodArrays(3, 2, 1)
	fmt.Println(res)
}

const (
	maxn = 1_00000
	mod  = 1_00000_0007
)

// 阶乘
var fact, invF [maxn]int

func init() {
	fact[0] = 1
	for i := 1; i < maxn; i++ {
		fact[i] = fact[i-1] * i % mod
	}
	invF[maxn-1] = pow(fact[maxn-1], mod-2)
	for i := maxn - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func c(n, k int) int {
	return fact[n] * invF[k] % mod * invF[n-k] % mod
}

func countGoodArrays(n int, m int, k int) int {
	return c(n-1, k) * m % mod * pow(m-1, n-1-k) % mod
}
