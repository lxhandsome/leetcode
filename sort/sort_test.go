package sort_test

import (
	"fmt"
	"leetcode/sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testCase() [][]int {
	return [][]int{
		{},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 4, 5, 6, 2, 3},
	}
}

func TestQuickSort(t *testing.T) {
	cases := testCase()
	for i, c := range cases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			sort.QuickSort(c)
			for i, j := 0, len(c)-1; i < j; {
				assert.Equal(t, true, c[i] <= c[j], fmt.Sprintf("下标 %d 值 %d 大于 下标 %d值%d", i, j, c[i], c[j]))
				i++
				j--
			}
		})
	}
}
