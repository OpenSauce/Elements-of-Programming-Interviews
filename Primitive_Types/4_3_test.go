package chapter_4

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBits(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"1000111 becomes 1110001": {
			input:    "1000111",
			expected: "1110001",
		},
		"110 becomes 011": {
			input:    "110",
			expected: "011",
		},
		"10101011 becomes 11010101": {
			input:    "10101011",
			expected: "11010101",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			value, _ := strconv.ParseUint(tc.input, 2, 64)
			actual := ReverseBits(value)
			assert.Equal(t, tc.expected, fmt.Sprintf("%b", actual))
		})
	}
}
