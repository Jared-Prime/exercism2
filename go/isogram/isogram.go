package isogram

import (
	"regexp"
	"strings"
)

var countable = regexp.MustCompile("[a-z]")

// IsIsogram says if the given string is an isogram
func IsIsogram(word string) bool {
	letters := make(map[rune]int)

	word = strings.ToLower(word)

	for _, letter := range word {
		if countable.MatchString(string(letter)) {
			letters[letter]++
		}

		if letters[letter] > 1 {
			return false
		}
	}

	return true
}
