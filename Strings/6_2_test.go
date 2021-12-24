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
		"String Input 30": {
			input1:   "30",
			input2:   10,
			input3:   2,
			expected: "00001",
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
