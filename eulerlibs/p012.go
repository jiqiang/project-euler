package eulerlibs

// FindPrimeFactorsCount returns prime factors counts of an integer number.
func FindPrimeFactorsCount(n int) int {

	pfc := make(map[int]int)
	result := 1

	for i := 2; i <= n; i++ {
		// count prime only
		if !IsPrime2(i) {
			continue
		}

		if n%i == 0 {
			n = n / i
			_, ok := pfc[i]
			if ok {
				pfc[i]++
			} else {
				pfc[i] = 1
			}
			i--
		}
	}

	for _, count := range pfc {
		result = result * (count + 1)
	}

	return result
}
