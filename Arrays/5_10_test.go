package chapter_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermuteArrayElements(t *testing.T) {
	testCases := map[string]struct {
		input1   []int
		input2   []int
		expected []int
	}{
		"Perm Example #1": {
			input1:   []int{1, 2, 3, 4},
			input2:   []int{2, 0, 1, 3},
			expected: []int{2, 3, 1, 4},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := PermuteArrayElements(tc.input1, tc.input2)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
