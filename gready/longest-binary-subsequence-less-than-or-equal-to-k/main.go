package main

import (
	"fmt"
	"strconv"

)
// https://leetcode.cn/problems/longest-binary-subsequence-less-than-or-equal-to-k/description/?envType=daily-question&envId=2025-06-26
func main() {
	fmt.Println(longestSubsequence("1001010", 5))
}

func longestSubsequence(s string, k int) int {

	ks := strconv.FormatInt(int64(k), 2)
	if len(ks) > len(s) {
		return len(s)
	}
	sub := s[len(s)-len(ks):]
	pz := len(sub)
	prefix := s[:len(s)-len(ks)]
	for _, c := range prefix {
		if c == '0' {
			pz++
		}
	}
	subInt, _ := strconv.ParseInt(sub, 2, 64)
	if int(subInt) > k {
		pz -= 1
	}
	return pz
}
