package grains

import "errors"

// Square says how many grains are on each square
func Square(n int) (uint64, error) {
	var count uint64

	if n <= 0 || n > 64 {
		return count, errors.New("Invalid square")
	}
	count = 1 << uint(n-1)

	return count, nil
}

// Total gives total number of grains
func Total() uint64 {
	var total uint64

	for i := 1; i <= 64; i++ {
		count, _ := Square(i)

		total += count
	}

	return total
}
