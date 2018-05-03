// Package bob allows you to say Hey to Bob, who does not care
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

var (
	uppercase = regexp.MustCompile("[A-Z]")
	lowercase = regexp.MustCompile("[a-z]")
)

// Hey returns a very limited response from Bob.
func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")

	switch {
	case isNothing(remark):
		return "Fine. Be that way!"
	case isShoutQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case isShout(remark):
		return "Whoa, chill out!"
	case isQuestion(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isShout(remark string) bool {
	return !lowercase.MatchString(remark) && uppercase.MatchString(remark)
}

func isShoutQuestion(remark string) bool {
	return isQuestion(remark) && isShout(remark)
}

func isNothing(remark string) bool {
	return remark == ""
}
