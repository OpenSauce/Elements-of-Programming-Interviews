package chapter_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrdering(t *testing.T) {
	testCases := map[string]struct {
		input1   [][]int
		expected []int
	}{
		"2x2 Matrix": {
			input1: [][]int{
				{1, 2},
				{4, 3},
			},
			expected: []int{1, 2, 3, 4},
		},
		"3x3 Matrix": {
			input1: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		"4x4 Matrix": {
			input1: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			expected: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := SpiralOrdering(tc.input1)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
