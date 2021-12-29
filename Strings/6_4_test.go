package chapter_6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceAndRemove(t *testing.T) {
	testCases := map[string]struct {
		input1   int
		input2   []rune
		expected []rune
	}{
		"Size 4 Example": {
			input1:   4,
			input2:   []rune{'a', 'b', 'a', 'c'},
			expected: []rune{'d', 'd', 'd', 'd', 'c'},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := ReplaceAndRemove(tc.input1, tc.input2)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
