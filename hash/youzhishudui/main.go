package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://www.lanqiao.cn/problems/3877/learning/?page=1&first_category_id=1&problem_id=3877

// 1<=i < j  < N
// a[i] = b[j]
// a[j] = b[i]
/*
使用双层for肯定超时，想到在遍历j的时候能够获取到i的结果
a[i] = b[j]
b[i] = a[j]

由上可以得到以下面等式
a[i]*base + b[i] == b[j] *base + a[j]

我们可以使用hash表保存a[j]*base + b[j]
遍历j的时候其实i的值已经在hash表中存在了
我们要获取i的值可以通过 b[j]*base + a[j]来获取，如果有则count++

*/

var base = 100000001

func main() {
	// lanqiao()

}

func lanqiao() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := convert(scanner.Text())
	scanner.Scan()
	a := convertNumArr(scanner.Text())
	scanner.Scan()
	b := convertNumArr(scanner.Text())
	m := make(map[int]int)
	// n, a, b := testCase()
	cnt := 0
	for j := 0; j < n; j++ {
		// 查询当前下表有多少优质数对
		x := b[j]*base + a[j]
		if _, ok := m[x]; ok {
			cnt += m[x] + 1
		}
		// 当前下表作为历史保存到hash表中,供下一个下标进行查询
		y := a[j]*base + b[j]
		if _, ok := m[y]; ok {
			m[y] += 1
		} else {
			m[y] = 0
		}
	}
	return cnt
}

func convertNumArr(s string) (res []int) {
	ss := strings.Split(s, " ")
	for _, s := range ss {
		res = append(res, convert(s))
	}
	return
}

func convert(s string) int {
	n, _ := strconv.ParseInt(s, 10, 64)
	return int(n)
}

func testCase() (int, []int, []int) {
	return 3, []int{1, 2, 3}, []int{3, 2, 1}
}
