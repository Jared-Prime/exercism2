package binarygap

// LongestZeros finds the longest sequence of zeros in a binary representation of an integer
func LongestZeros(n int) int {
	var count, maximum int

	if isEven(n) {
		n, _ = removeTrailingZeros(n)
	}

	for isPositive(n) {
		n >>= 1

		if isEven(n) {
			n, count = removeTrailingZeros(n)

			if count > maximum {
				maximum = count
			}
		}
	}

	return maximum
}

func removeTrailingZeros(n int) (int, int) {
	var removed int

	for isPositiveAndEven(n) {
		n >>= 1
		removed++
	}

	return n, removed
}

func isPositiveAndEven(n int) bool {
	return isPositive(n) && isEven(n)
}

func isEven(n int) bool {
	return n&1 == 0
}

func isPositive(n int) bool {
	return n > 0
}
