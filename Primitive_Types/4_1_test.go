package chapter_4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParity(t *testing.T) {
	testCases := map[string]struct {
		input    int64
		expected int
	}{
		"Parity of 0": {
			input:    3,
			expected: 0,
		},
		"Parity of 1": {
			input:    1,
			expected: 1,
		},
		"Big.Int": {
			input:    546546754745674,
			expected: 1,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := Parity(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
