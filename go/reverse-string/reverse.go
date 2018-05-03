package reverse

// String reverses the input string
func String(input string) string {
	inputRunes := []rune(input)
	inputLength := len(inputRunes)
	outputRunes := make([]rune, inputLength)

	for i, inputRune := range inputRunes {
		outputRunes[inputLength-i-1] = inputRune
	}

	return string(outputRunes)
}
