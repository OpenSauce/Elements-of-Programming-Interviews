package chapter_4

func Compute(x float64, y int) float64 {
	result := 1.0
	power := y
	if y < 0 {
		power = -power
		x = 1.0 / x
	}
	for power > 0 {
		if power&1 == 1 {
			result *= x
		}
		x *= x
		power >>= 1
	}
	return result
}
