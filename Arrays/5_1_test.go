package chapter_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDutchNationalFlag(t *testing.T) {
	testCases := map[string]struct {
		input1   []int
		input2   int
		expected []int
	}{
		"Index 3 - Pivot of 0": {
			input1:   []int{0, 1, 2, 0, 2, 1, 1},
			input2:   3,
			expected: []int{1, 1, 2, 2, 2, 2, 2},
		},
		"Index 2 - Pivot of 2": {
			input1:   []int{0, 1, 2, 0, 2, 1, 1},
			input2:   2,
			expected: []int{0, 0, 0, 0, 0, 1, 1},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			pivot := tc.input1[tc.input2]
			sortedArray := DutchNationalFlag(tc.input1, tc.input2)
			for i, value := range tc.expected {
				switch value {
				case 0:
					assert.Less(t, sortedArray[i], pivot)
				case 1:
					assert.Equal(t, sortedArray[i], pivot)
				case 2:
					assert.Greater(t, sortedArray[i], pivot)
				}

			}
		})
	}
}
