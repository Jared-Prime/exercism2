package raindrops

import "fmt"

const (
	plong = "Plong"
	plang = "Plang"
	pling = "Pling"
)

// Convert turns an integer into a fun string
func Convert(input int) string {
	var output string

	if input%3 == 0 {
		output = output + pling
	}

	if input%5 == 0 {
		output = output + plang
	}

	if input%7 == 0 {
		output = output + plong
	}

	if output == "" {
		output = fmt.Sprintf("%v", input)
	}

	return output
}

func isPling(n int) bool {
	return n%3 == 0
}

func isPlong(n int) bool {
	return n%7 == 0
}

func isPlang(n int) bool {
	return n%5 == 0
}
