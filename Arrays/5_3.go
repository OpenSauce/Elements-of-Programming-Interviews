package chapter_5

import (
	"math"
)

func MultiplyTwoIntegers(n1, n2 []int) []int {
	sign := 1
	if n1[0] < 0 && n2[0] >= 0 || n2[0] < 0 && n1[0] >= 0 {
		sign = -1
	}
	n1[0] = int(math.Abs(float64(n1[0])))
	n2[0] = int(math.Abs(float64(n2[0])))

	var nums = make([]int, len(n1)+len(n2))
	for i := len(n1) - 1; i >= 0; i-- {
		for j := len(n2) - 1; j >= 0; j-- {
			nums[i+j+1] += n1[i] * n2[j]
			nums[i+j] += nums[i+j+1] / 10
			nums[i+j+1] %= 10
		}
	}

	index := 0
	for i, j := range nums {
		if j != 0 {
			index = i
			break
		}
	}

	nums = nums[index:]
	nums[0] *= sign
	return nums
}
