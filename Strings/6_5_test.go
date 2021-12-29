package chapter_6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromicity(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected bool
	}{
		"A Palindromic String": {
			input:    "ray a yar",
			expected: true,
		},
		"A Palindromic, Capitalized String": {
			input:    "Ray a yar",
			expected: true,
		},
		"A Non-Palindromic String": {
			input:    "Barbara strisand",
			expected: false,
		},
		"A Palindromic String With Punctuation": {
			input:    "ray, a yar!",
			expected: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := Palindromicity(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
