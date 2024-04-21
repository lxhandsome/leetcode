package main

import "fmt"

// https://leetcode.cn/problems/number-of-excellent-pairs/description/
/*
解题思路，数据范围超过10^5 所以不能使用双重for来遍历计算 and + or ，数据中某个数字可能会出现多次，所以需要去重

 and + or 的和 其实就等于 两个数字中二进制为为1和 = num1 为1 + num2 为1
可以使用bits[i]来记录每二进制为i的数字数量
*/

func redumpicate(nums []int) (res []int) {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}
	for n := range m {
		res = append(res, n)
	}
	return
}

func count1(a int) int {
	cnt := 0
	for a != 0 {
		r := a & 1
		if r == 1 {
			cnt++
		}
		a = a >> 1
	}
	return cnt
}

func countExcellentPairs(nums []int, k int) int64 {
	nums = redumpicate(nums)
	bits := make([]int, 33)
	for _, num := range nums {
		bits[count1(num)]++
	}
	var cnt int64
	for i := 0; i < 33; i++ {
		for j := 0; j < 33; j++ {
			if j < k-i {
				continue
			}
			cnt += int64(bits[i] * bits[j])
		}
	}
	return cnt
}

func testCase1() []int {
	return []int{1, 2, 3, 1}
}

func main() {
	fmt.Printf("countExcellentPairs(testCase1(), 3): %v\n", countExcellentPairs(testCase1(), 3))
}
