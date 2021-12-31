package chapter_6

import "strconv"

func LookAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		res = NextNumber(res)
	}
	return res
}

func NextNumber(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		count := 1
		for i+1 < len(s) && s[i] == s[i+1] {
			i++
			count++
		}
		res += strconv.Itoa(count) + string(s[i])
	}
	return res
}
