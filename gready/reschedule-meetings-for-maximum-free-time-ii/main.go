package main

// https://leetcode.cn/problems/reschedule-meetings-for-maximum-free-time-ii/?envType=daily-question&envId=2025-07-10

/*
1. 将会议移动到可以移动的最大空闲位置上去，当前的最大空闲长度 = max(当前最大空闲长度，会议leftFree + 会议长度 + 会议rightFree)
2. 如果不能移动到最最大空闲位置上去，当前会议智能左右移动    最大空闲长度 = (当前最大空闲长度，会议leftFree + 会议长度 + 会议rightFree)
3. 会议移动过程中，可以会议动到与当前会议享有的左右空闲位置，相当于没移动，第一种情况我们需要记录前三大的空闲位置
4. 所以要记录前三个空闲位置（a ，b ，c），分类讨论
	当桌子能够一移动到三空闲位置
	if 当前会议左右两边有 a 尝试b
	if 当前会议左右两边有 b 尝试c
	在移动的过程中更新最大的空闲长度
	当不能移动到前线空闲位置，那么该会议就只能左右移动


*/

func maxFreeTime(eventTime int, startTime []int, endTime []int) int {

	n := len(startTime)

	// 获取第i个空闲长度
	get := func(i int) int {
		if i == 0 {
			return startTime[0]
		}
		if i == n {
			return eventTime - endTime[n-1]
		}
		return startTime[i] - endTime[i-1]
	}

	// 获取前三大的空闲位置
	a, b, c := 0, -1, -1
	for i := 1; i <= n; i++ {
		free := get(i)
		if free > get(a) {
			a, b, c = i, a, b
		} else if b < 0 || free > get(b) {
			b, c = i, b
		} else if c < 0 || free > get(c) {
			c = i
		}
	}

	res := 0
	for i := range n {
		hy := endTime[i] - startTime[i]
		if (i != a && i+1 != a && get(a) >= hy) || (i != b && i+1 != b && get(b) >= hy) || get(c) >= hy {
			res = max(res, get(i)+hy+get(i+1))
		} else {
			res = max(res, get(i)+get(i+1))
		}
	}

	return res
}