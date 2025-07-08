package main

import (
	"sort"

)

// https://leetcode.cn/problems/maximum-profit-in-job-scheduling/solutions/1913089/dong-tai-gui-hua-er-fen-cha-zhao-you-hua-zkcg/

type job struct {
	start  int
	end    int
	profit int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([]job, n)
	for i := range startTime {
		jobs[i] = job{
			start:  startTime[i],
			end:    endTime[i],
			profit: profit[i],
		}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})
	//f[i] 为前i个工作最多能获得的利润
	//f[0] =0
	//f[i] = max(f[i-1], f[j]+profit[i]),j 为第一个小于等于start[i]的工作的下标

	f := make([]int, n+1)
	for i, job := range jobs {
		// 二分查找找到第一个结束时间大于当前开始时间的下表
		// j - 1就是结束时间 <= 当前工作的开始时间的小标
		j := sort.Search(i, func(j int) bool { return jobs[j].end > job.start })
		// 状态转移中，为什么是 j 不是 j+1：上面算的是 > start，-1 后得到 <= start，但由于还要 +1，抵消了
		// 前i个工作的做大获利 = max(不做第i个工作的获利，做第i个工作的获利 + 做前j个工作的获利)
		f[i+1] = max(f[i], f[j]+job.profit)
	}
	return f[len(jobs)]

}
