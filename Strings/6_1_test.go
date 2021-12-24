package chapter_6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToInteger(t *testing.T) {
	testCases := map[string]struct {
		input1   string
		expected int
	}{
		"String Input 30": {
			input1:   "30",
			expected: 30,
		},
		"Minus Input": {
			input1:   "-30",
			expected: -30,
		},
		"Multiple Digits": {
			input1:   "485",
			expected: 485,
		},
		"Zero Digit": {
			input1:   "0",
			expected: 0,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := StringToInteger(tc.input1)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestIntegerToString(t *testing.T) {
	testCases := map[string]struct {
		input1   int
		expected string
	}{
		"Integer Input 30": {
			input1:   30,
			expected: "30",
		},
		"Minus Input": {
			input1:   -30,
			expected: "-30",
		},
		"Multiple Characters": {
			input1:   485,
			expected: "485",
		},
		"Zero Character": {
			input1:   0,
			expected: "0",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := IntegerToString(tc.input1)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
