package chapter_5

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleOfflineData(t *testing.T) {
	testCases := map[string]struct {
		input1 []int
		input2 int
	}{
		"Random 4 Subset": {
			input1: []int{310, 315, 275, 295, 260, 270, 290, 230, 255, 250},
			input2: 4,
		},
		"Random 5 Subset": {
			input1: []int{500, 510, 450, 455, 500, 480, 476},
			input2: 5,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var actual []int
			firstVals := SampleOfflineData(tc.input1, tc.input2)
			copy(actual, firstVals)
			actual2 := SampleOfflineData(tc.input1, tc.input2)
			assert.False(t, reflect.DeepEqual(actual, actual2))
		})
	}
}
