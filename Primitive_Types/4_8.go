package chapter_4

func ReverseDigits(x int) int {
	result := 0
	for x != 0 { // Whilst x is not 0
		result = result*10 + x%10 // Shift result up by a factor of 10, then add the least significant digit from x
		x /= 10                   // Shift x down by 10
	}
	return result
}
