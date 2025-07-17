package main

import "slices"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	slices.SortFunc(players, func(a, b int) int { return a - b })
	slices.SortFunc(trainers, func(a, b int) int { return a - b })
	var p, t int
	res := 0
	for p < len(players) && t < len(trainers) {
		if players[p] <= trainers[t] {
			res++
			p++
			t++
		} else {
			t++
		}
	}
	return res
}
