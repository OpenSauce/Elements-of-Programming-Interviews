package chapter_5

func DutchNationalFlag(array []int, i int) []int {
	pivot := array[i]
	smaller := 0
	equal := 0
	larger := len(array)

	for equal < larger {
		if array[equal] < pivot {
			array[smaller], array[equal] = array[equal], array[smaller]
			smaller++
			equal++
		} else if array[equal] == pivot {
			equal++
		} else {
			larger--
			array[equal], array[larger] = array[larger], array[equal]
		}
	}
	return array
}
