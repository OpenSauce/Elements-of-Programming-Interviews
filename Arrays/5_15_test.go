package chapter_5

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeRandomSubset(t *testing.T) {
	testCases := map[string]struct {
		input1 int
		input2 int
	}{
		"Random Subset #1": {
			input1: 4,
			input2: 2,
		},
		"Random Subset #2": {
			input1: 8,
			input2: 8,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var actual []int
			firstVals := ComputeRandomSubset(tc.input1, tc.input2)
			copy(actual, firstVals)
			actual2 := ComputeRandomSubset(tc.input1, tc.input2)
			assert.False(t, reflect.DeepEqual(actual, actual2))
		})
	}
}
