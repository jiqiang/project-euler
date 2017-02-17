package eulerlibs

// SumOfSquares calculates sum of square of a series of numbers.
func SumOfSquares(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

// SquareOfSum calculates square of sum of a series of numbers.
func SquareOfSum(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result * result
}

// RunP6 runs problem 6 solution.
func RunP6(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
