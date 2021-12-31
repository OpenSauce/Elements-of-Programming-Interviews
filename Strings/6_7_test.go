package chapter_6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookAndSay(t *testing.T) {
	testCases := map[string]struct {
		input    int
		expected string
	}{
		"First Eight": {
			input:    8,
			expected: "1113213211",
		},
		"First Two": {
			input:    3,
			expected: "21",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := LookAndSay(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
