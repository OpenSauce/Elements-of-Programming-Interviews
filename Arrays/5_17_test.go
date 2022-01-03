package chapter_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSudokuChecker(t *testing.T) {
	testCases := map[string]struct {
		input1   [][]int
		expected bool
	}{
		"Valid Sudoku": {
			input1: [][]int{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: true,
		},
		"Invalid Sudoku - Column": {
			input1: [][]int{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 9},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
		"Invalid Sudoku - Row": {
			input1: [][]int{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 6, 2, 8, 0},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
		"Leetcode TestCase": {
			input1: [][]int{
				{0, 0, 5, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 8, 0, 0, 0, 3, 0},
				{0, 5, 0, 0, 2, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 9},
				{0, 0, 0, 0, 0, 0, 4, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 7},
				{0, 1, 0, 0, 0, 0, 0, 0, 0},
				{2, 4, 0, 0, 0, 0, 9, 0, 0},
			},
			expected: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := SudokuChecker(tc.input1)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
