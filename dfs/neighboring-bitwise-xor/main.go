package main

import "fmt"

func main() {
	fmt.Printf("doesValidArrayExist([]int{1, 1, 0}): %v\n", doesValidArrayExist([]int{1, 1, 0}))
}

func doesValidArrayExist(derived []int) bool {
	n := len(derived)
	original := make([]int, n)

	var dfs func(index int) bool

	dfs = func(index int) bool {
		if index == n {
			return derived[index-1] == original[index-1]^original[0]
		}
		if index == 0 {
			for i := 0; i <= 1; i++ {
				original[index] = i
				if dfs(index + 1) {
					return true
				}
			}
			return false
		}
		for i := 0; i <= 1; i++ {
			if derived[index-1] == original[index-1]^i {
				original[index] = i
				if dfs(index + 1) {
					return true
				}
			}
		}
		return false
	}

	return dfs(0)

}
