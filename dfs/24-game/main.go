package main

import (
	"fmt"
	"math"

)

func main() {
	fmt.Printf("judgePoint24([]int{4, 1, 8, 7}): %v\n", judgePoint24([]int{3, 3, 8, 8}))
}

func judgePoint24(cards []int) bool {

	fcards := make([]float64, 4)
	for i, x := range cards {
		fcards[i] = float64(x)
	}
	return fjudge(fcards)
}

func fjudge(fcards []float64) bool {
	e := 1e-9
	if len(fcards) == 1 {
		return math.Abs(fcards[0]-24.0) < e
	}
	for i, x := range fcards {
		for j := i + 1; j < len(fcards); j++ {
			y := fcards[j]
			cadidates := []float64{x - y, y - x, x + y, x * y}
			if math.Abs(y) > e {
				cadidates = append(cadidates, x/y)
			}
			if math.Abs(x) > e {
				cadidates = append(cadidates, y/x)
			}
			for _, c := range cadidates {
				newCards := make([]float64, 0)
				newCards = append(newCards, fcards[0:j]...)
				newCards = append(newCards, fcards[j+1:]...)
				newCards[i] = c
				if fjudge(newCards) {
					return true
				}
			}
		}
	}
	return false
}
