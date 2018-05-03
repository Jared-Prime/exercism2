// Package acronym creates an acronym from a given string
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns a jargony looking acronym for a given string
func Abbreviate(s string) string {
	var acronym []byte
	for _, word := range words(s) {
		acronym = append(acronym, firstLetter(word))
	}

	return strings.ToUpper(string(acronym))
}

func words(longName string) []string {
	return strings.FieldsFunc(longName, func(c rune) bool {
		return unicode.IsSpace(c) || c == '-'
	})
}

func firstLetter(word string) byte {
	return word[0]
}
