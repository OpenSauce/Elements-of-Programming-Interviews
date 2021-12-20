package chapter_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumerateAllPrimes(t *testing.T) {
	testCases := map[string]struct {
		input1   int
		expected []int
	}{
		"Primes Leading upto 18": {
			input1:   18,
			expected: []int{2, 3, 5, 7, 11, 13, 17},
		},
		"Primes Leading upto 60": {
			input1:   60,
			expected: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := EnumerateAllPrimes(tc.input1)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
