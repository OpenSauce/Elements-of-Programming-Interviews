package chapter_4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimalPalindrome(t *testing.T) {
	testCases := map[string]struct {
		input    int
		expected bool
	}{
		"7": {
			input:    7,
			expected: true,
		},
		"121": {
			input:    121,
			expected: true,
		},
		"2147447412": {
			input:    2147447412,
			expected: true,
		},
		"-1": {
			input:    -1,
			expected: false,
		},
		"12": {
			input:    12,
			expected: false,
		},
		"2147483647": {
			input:    2147483647,
			expected: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := DecimalPalindrome(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
