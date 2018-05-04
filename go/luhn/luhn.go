package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid performs a Luhn checksum on the input number string
func Valid(input string) bool {
	var sum int
	var double bool

	input = clean(input)

	if input == "0" {
		return false
	}

	length := len(input)

	for i := length - 1; i > -1; i-- {
		digit, err := digit(input, i)
		if err != nil {
			return false
		}

		if double {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}
		}

		double = !double

		sum += digit
	}

	return sum%10 == 0
}

func digit(input string, i int) (int, error) {
	return strconv.Atoi(string(input[i]))
}

func clean(input string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, input)
}
