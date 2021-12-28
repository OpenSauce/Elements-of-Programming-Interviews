package chapter_6

import (
	"math"
)

func BaseConversion(number string, b1 int, b2 int) string {
	var total int
	for i, r := range number {
		var val int
		if (r - 'A') >= 0 {
			val = int(r - 'A' + 10)
		} else {
			val = int(r - '0')
		}
		currentIteration := len(number) - i - 1
		total += val * int(math.Pow(float64(b1), float64(currentIteration)))
	}
	return constructFromBase(total, b2)
}

func constructFromBase(num int, base int) string {
	if num == 0 {
		return ""
	}

	var numToA rune
	if num%base >= 10 {
		numToA = rune('A' + (num%base - 10))
	} else {
		numToA = rune('0' + num%base)
	}

	return constructFromBase(num/base, base) + string(numToA)
}
