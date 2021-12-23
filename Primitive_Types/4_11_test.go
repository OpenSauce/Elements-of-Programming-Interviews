package chapter_4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangleIntersection(t *testing.T) {
	testCases := map[string]struct {
		input1, input2, expected Rect
	}{
		"Intersecting": {
			input1:   Rect{1, 1, 4, 4},
			input2:   Rect{2, 2, 4, 4},
			expected: Rect{2, 2, 3, 3},
		},
		"Intersecting #2": {
			input1:   Rect{3, 2, 4, 4},
			input2:   Rect{1, 1, 4, 4},
			expected: Rect{3, 2, 3, 2},
		},
		"Not intersecting": {
			input1:   Rect{5, 6, 2, 2},
			input2:   Rect{1, 1, 2, 2},
			expected: Rect{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := RectangleIntersection(tc.input1, tc.input2)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
