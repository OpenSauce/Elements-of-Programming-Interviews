package chapter_4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	testCases := map[string]struct {
		input1   float64
		input2   int64
		expected float64
	}{
		"3^3=18": {
			input1:   3.0,
			input2:   3,
			expected: 27.0,
		},
		"12^2=144": {
			input1:   12,
			input2:   2,
			expected: 144.0,
		},
		"5^2 =25": {
			input1:   5,
			input2:   2,
			expected: 25.0,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := Compute(tc.input1, int(tc.input2))
			assert.Equal(t, tc.expected, actual)
		})
	}
}
