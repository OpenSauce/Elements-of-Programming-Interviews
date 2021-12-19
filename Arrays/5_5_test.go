package chapter_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	testCases := map[string]struct {
		input1   []int
		expected []int
	}{
		"Duplicate Example": {
			input1:   []int{2, 3, 5, 5, 7, 11, 11, 13},
			expected: []int{2, 3, 5, 7, 11, 13},
		},
		"All Duplicates": {
			input1:   []int{1, 1, 1, 1, 1},
			expected: []int{1},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := DeleteDuplicates(tc.input1)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
