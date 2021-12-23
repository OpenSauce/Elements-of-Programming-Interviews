package chapter_5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuySellStockOnce(t *testing.T) {
	testCases := map[string]struct {
		input1   []int
		expected int
	}{
		"Buy at 260 sell at 290": {
			input1:   []int{310, 315, 275, 295, 260, 270, 290, 230, 255, 250},
			expected: 30,
		},
		"Buy at 450 sell at 500": {
			input1:   []int{500, 510, 450, 455, 500, 480, 476},
			expected: 50,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := BuySellStockOnce(tc.input1)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
