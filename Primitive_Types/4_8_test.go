package chapter_4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseDigits(t *testing.T) {
	testCases := map[string]struct {
		input    int
		expected int
	}{
		"30 becomes 03": {
			input:    30,
			expected: 03,
		},
		"155 becomes 551": {
			input:    155,
			expected: 551,
		},
		"-24 becomes -42": {
			input:    -24,
			expected: -42,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := ReverseDigits(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
