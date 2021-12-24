package chapter_6

import "fmt"

func StringToInteger(num string) int {
	var res int
	for i := len(num) - 1; i > 0; i-- {
		res += int('0' - num[i])
		res *= 10
	}
	return res
}

func IntegerToString(num int) string {
	if num == 0 {
		return "0"
	}

	var res string
	for num != 0 {
		res = fmt.Sprint(rune(num%10)) + res
		num /= 10
	}
	return res
}
