package chapter_6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseAllWords(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"Bob likes Alice": {
			input:    "Bob likes Alice",
			expected: "Alice likes Bob",
		},
		"Another Test Sentence": {
			input:    "Here we go again",
			expected: "again go we Here",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := ReverseAllWords(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
