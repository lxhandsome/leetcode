package main

/*
思路：
贪心算法： 每次都选择能够跳到更远距离的位置跳上去

当前在i位置，可以跳到最大的距离是nums[i] 依次判断[nums[i+1] - nums[i+nums[i]]]中能够调大的最大距离，选择能够跳的最远的位置，作为下次落脚点

*/

func jump(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	f[0] = 0
	for i := 1; i < n; i++ {
		// 取第一个能跳到当前位置的点即可
		// 因为跳跃次数的结果集是单调递增的，所以贪心思路是正确的
		idx := 0
		for idx < n && idx+nums[idx] < i {
			idx++
		}
		f[i] = f[idx] + 1
	}
	return f[n-1]
}
