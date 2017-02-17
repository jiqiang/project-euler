package eulerlibs

// EvenFibonacci finds even fibonacci number which is less than given term.
func EvenFibonacci(term int) []int {
	var evenFibs []int

	a, b := 1, 2

	for {
		if a > term {
			break
		}

		if a%2 == 0 {
			evenFibs = append(evenFibs, a)
		}

		a, b = b, a+b
	}

	return evenFibs
}

// RunP2 runs problem 2 solution.
func RunP2(max int) int {
	sum := 0
	for _, n := range EvenFibonacci(max) {
		sum += n
	}
	return sum
}
