package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testCase() ([][]int, [][]int) {
	res := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 2, 1},
	}
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	return mat, res
}

func TestUpdateMatrix(t *testing.T) {
	mat, expect := testCase()
	actual := updateMatrix(mat)
	for i := 0; i < len(expect); i++ {
		for j := 0; j < len(expect[i]); j++ {
			assert.Equal(t, expect[i][j], actual[i][j], fmt.Sprintf("第 %d 层 第 %d 个元素不同", i+1, j+1))
		}
	}

}
