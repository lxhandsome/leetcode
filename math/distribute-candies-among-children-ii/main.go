package main

// https://leetcode.cn/problems/distribute-candies-among-children-ii/description/?envType=daily-question&envId=2025-06-01

/*
解题思路：

1, 题目要求： n个一模一样的糖果分给3个不同的孩子，每个孩子不超过limit 的分法有多少种

2, n个一模一样的糖果分给3个不同的孩子，有多少种g不同的n分法

	转化为n个位置切两刀，有多少种不同的切法 -》 n + 2个位置选择两个位置用来切割，有多少种不同的方法 -》 c(n+2,2)

3, 答案  n个n一模一样的糖果分给3个孩子所有的分法 - 1个孩子分到的糖果大于limmit（a1） - 2个孩子分到的糖果大于limit（a2） - 3个孩子分到的糖果大于limit（a3）

	c 三个孩子分n个糖果 c(n+2，2）
	c1至少有一个孩子分到的糖果超过limit  3* c(n+2 - (limit +1),2)
	c2至少有两个孩子分到的糖果超过limit  3 * c(n+2 - 2*(limit +1),2)
	c3所有孩子分到的糖果超过limit       c(c +2 -3*(limit +1),2)

	res =  c - (c1  - c2  + c3)
*/
func distributeCandies(n int, limit int) int64 {

	all := cn2(n + 2)
	atLeastOne := cn2(n-(limit+1)+2) * 3
	atLeastTwo := cn2(n-2*(limit+1)+2) * 3
	atLeastThree := cn2(n - 3*(limit+1) + 2)

	return all - atLeastOne + atLeastTwo - atLeastThree
}

func cn2(n int) int64 {
	if n < 0 {
		return 0
	}

	return int64(n) * int64(n-1) / 2
}
