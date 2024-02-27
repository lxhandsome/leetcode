package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCases = [][]int{
		// {2, 1, 5, 6, 2, 3},
		// {2, 4},
		// {0},
		{2, 1, 2},
	}

	// testRes = []int{10, 4, 0, 3}
	testRes = []int{3}
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestLargestRectangleArea(t *testing.T) {
	for i, c := range testCases {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			res := largestRectangleArea(c)
			assert.Equal(t, testRes[i], res)
		})
	}
}
