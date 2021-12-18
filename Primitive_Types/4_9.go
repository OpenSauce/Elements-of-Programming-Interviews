package chapter_4

import "math"

func DecimalPalindrome(x int) bool {
	if x <= 0 {
		return x == 0
	}

	for x != 0 {
		n := math.Floor(math.Log10(float64(x))) + 1
		LSD := x % 10
		MSD := x / (int(math.Pow(float64(10), float64(n-1))))

		if LSD != MSD {
			return false
		}

		x %= int(math.Pow(float64(10), float64(n-1)))
		x /= 10
	}
	return true
}
