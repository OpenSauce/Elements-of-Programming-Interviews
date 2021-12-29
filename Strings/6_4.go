package chapter_6

func ReplaceAndRemove(size int, input []rune) []rune {
	writeIndex := 0
	count := 0
	for i := 0; i < size; i++ {
		if input[i] != 'b' {
			input[writeIndex] = input[i]
			writeIndex++
		}
		if input[i] == 'a' {
			count++
		}
	}

	currentIndex := writeIndex - 1
	writeIndex = writeIndex + count - 1
	res := make([]rune, writeIndex+1)
	copy(res, input)

	for currentIndex >= 0 {
		if input[currentIndex] == 'a' {
			res[writeIndex] = 'd'
			writeIndex--
			res[writeIndex] = 'd'
			writeIndex--
		} else {
			res[writeIndex] = input[currentIndex]
			writeIndex--
		}
		currentIndex--
	}
	return res
}
