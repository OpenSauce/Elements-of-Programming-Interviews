package chapter_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplyTwoIntegers(t *testing.T) {
	testCases := map[string]struct {
		input1   []int
		input2   []int
		expected []int
	}{
		"Big Negative": {
			input1:   []int{1, 9, 3, 7, 0, 7, 7, 2, 1},
			input2:   []int{-7, 6, 1, 8, 3, 8, 2, 5, 7, 2, 8, 7},
			expected: []int{-1, 4, 7, 5, 7, 3, 9, 5, 2, 5, 8, 9, 6, 7, 6, 4, 1, 2, 9, 2, 7},
		},
		"Simple Multiplication": {
			input1:   []int{1, 5},
			input2:   []int{3},
			expected: []int{4, 5},
		},
		"Another Simple Multiplication": {
			input1:   []int{3, 0},
			input2:   []int{4, 2, 5},
			expected: []int{1, 2, 7, 5, 0},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := MultiplyTwoIntegers(tc.input1, tc.input2)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
