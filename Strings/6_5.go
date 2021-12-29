package chapter_6

import "unicode"

func Palindromicity(word string) bool {
	firstIndex := 0
	lastIndex := len(word) - 1
	for firstIndex < lastIndex {
		if !unicode.IsLetter(rune(word[firstIndex])) {
			firstIndex++
			continue
		}
		if !unicode.IsLetter(rune(word[lastIndex])) {
			lastIndex--
			continue
		}
		if unicode.ToLower(rune(word[firstIndex])) != unicode.ToLower(rune(word[lastIndex])) {
			return false
		}
		firstIndex++
		lastIndex--
	}
	return true
}
