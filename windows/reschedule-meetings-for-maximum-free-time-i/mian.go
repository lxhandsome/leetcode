package main

// https://leetcode.cn/problems/reschedule-meetings-for-maximum-free-time-i/description/?envType=daily-question&envId=2025-07-09

/*
1. 当k等于1的时候相当于合并两段空闲时间，结果就是在所有的合并空闲时间中秋最大值
2. 当k等于2的时候相当于合并三段空闲时间，结果就是在所有合并的空闲时间中求最大值
3. 题目就可以转换为在连续n + 1段空闲时间(0,startTime[0],startTime[1]-endTime[0],startTime[2]-endTime[1],...,endTime -endTime[n-1])中合并k+1段空闲时间，求最大值
4. 可以使用滑动窗口维护k+1段空闲时间的和，求最大值

*/

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	free := make([]int, 0, n+1)
	// 求出所有的空闲时间放到free数组中
	free = append(free, startTime[0])
	for i := 0; i < n-1; i++ {
		free = append(free, startTime[i+1]-endTime[i])
	}
	free = append(free, eventTime-endTime[n-1])
	m := 0
	l := 0
	sum := 0
	for r := range free {
		if r-l > k {
			sum -= free[l]
			l++
		}
		sum += free[r]
		m = max(sum, m)
	}
	return m

}

// 空间优化不使用ree数组
func maxFreeTime2(eventTime, k int, startTime, endTime []int) (ans int) {
	n := len(startTime)
	get := func(i int) int {
		if i == 0 {
			return startTime[0] // 最左边的空余时间段
		}
		if i == n {
			return eventTime - endTime[n-1] // 最右边的空余时间段
		}
		return startTime[i] - endTime[i-1] // 中间的空余时间段
	}

	s := 0
	for i := range n + 1 {
		s += get(i)
		if i < k {
			continue
		}
		ans = max(ans, s)
		s -= get(i - k)
	}
	return
}
