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
			input:    "0000000000000000000000000000000000000000000000000000000001000111",
			expected: "1110001000000000000000000000000000000000000000000000000000000000",
		},
		"110 becomes 011": {
			input:    "0000000000000000000000000000000000000000000000000000000000000110",
			expected: "0110000000000000000000000000000000000000000000000000000000000000",
		},
		"1...10101011 becomes 11010101...1": {
			input:    "1000000000000000000000000000000000000000000000000000000010101011",
			expected: "1101010100000000000000000000000000000000000000000000000000000001",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			value, _ := strconv.ParseUint(tc.input, 2, 64)
			expectedBits, _ := strconv.ParseUint(tc.expected, 2, 64)
			actual := ReverseBits(value)
			assert.Equal(t, fmt.Sprintf("%b", uint64(expectedBits)), fmt.Sprintf("%b", actual))
		})
	}
}
