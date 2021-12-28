package chapter_6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseConversion(t *testing.T) {
	testCases := map[string]struct {
		input1   string
		input2   int
		input3   int
		expected string
	}{
		"615 Base 7 to Base 13": {
			input1:   "615",
			input2:   7,
			input3:   13,
			expected: "1A7",
		},
		"1A7 Base 13 to Base 7": {
			input1:   "1A7",
			input2:   7,
			input3:   13,
			expected: "615",
		},
		"BBBA Base 16 to Base 4": {
			input1:   "BBBA",
			input2:   16,
			input3:   4,
			expected: "23232322",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := BaseConversion(tc.input1, tc.input2, tc.input3)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
