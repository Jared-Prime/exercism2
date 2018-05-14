package binarygap

// LongestZeros finds the longest sequence of zeros in a binary representation of an integer
func LongestZeros(n int) int {
	var initial, maximum int

	_, _, maximum = recur(n, initial, maximum)

	return maximum
}

func recur(n, count, maximum int) (int, int, int) {
	if n < 2 {
		return n, count, maximum
	}

	if n%2 == 0 {
		count++
	} else {
		count = 0
	}

	n = n >> 1

	if maximum <= count {
		maximum = count
	}

	return recur(n, count, maximum)
}
