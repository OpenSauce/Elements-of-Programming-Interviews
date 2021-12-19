package chapter_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrementAnInteger(t *testing.T) {
	testCases := map[string]struct {
		input1   []int
		expected []int
	}{
		"Base Case": {
			input1:   []int{1, 3, 2, 1},
			expected: []int{1, 3, 2, 2},
		},
		"Single Carry Over": {
			input1:   []int{1, 2, 9},
			expected: []int{1, 3, 0},
		},
		"Multiple Carry Over": {
			input1:   []int{1, 9, 9, 9},
			expected: []int{2, 0, 0, 0},
		},
		"Multiple Carry Overflowing": {
			input1:   []int{9, 9, 9, 9},
			expected: []int{1, 0, 0, 0, 0},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := IncrementAnInteger(tc.input1)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
