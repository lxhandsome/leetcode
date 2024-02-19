package main

import "fmt"

func makeSmallestPalindrome(s string) string {
	sByte := []byte(s)
	for i, j := 0, len(sByte)-1; i < j; {
		if sByte[i] < sByte[j] {
			sByte[j] = sByte[i]
		} else if sByte[i] > sByte[j] {
			sByte[i] = sByte[j]
		}
		i++
		j--
	}
	return string(sByte)
}

func testCase() []string {
	return []string{
		"egcfe",
		"abcd",
		"seven",
	}
}

func main() {
	for _, test := range testCase() {
		fmt.Println(makeSmallestPalindrome(test))
	}
}
