package main

import "fmt"

func main() {
	fmt.Printf("totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}): %v\n", totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}

func totalFruit(fruits []int) (ans int) {
	cnt := map[int]int{}
	left := 0
	for right, in := range fruits {
		cnt[in]++          // fruits[right] 进入窗口
		for len(cnt) > 2 { // 不满足要求
			out := fruits[left]
			cnt[out]-- // fruits[left] 离开窗口
			if cnt[out] == 0 {
				delete(cnt, out)
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
