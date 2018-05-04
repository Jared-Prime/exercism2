package diffsquares

// SquareOfSums returns square of sums
func SquareOfSums(n int) int {
	var result int

	for i := 1; i <= n; i++ {
		result += i
	}

	return result * result
}

// SumOfSquares returns sum of squares
func SumOfSquares(n int) int {
	var result int

	for i := 1; i <= n; i++ {
		result += (i * i)
	}

	return result
}

// Difference returns difference of squares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
