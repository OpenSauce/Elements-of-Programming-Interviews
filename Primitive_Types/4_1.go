package chapter_4

func Parity(x int64) int {
	var result int
	for x > 0 {
		result ^= 1
		x &= (x - 1)
	}
	return result
}
