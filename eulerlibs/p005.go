package eulerlibs

// LeastCommonMultiple finds least common multiple for a given number.
func LeastCommonMultiple(max int) int {

	var isLcm bool

	var lcm int

	lcm = max

	for i := 2; ; i++ {

		isLcm = true

		for n := 2; n < max; n++ {

			if lcm%n != 0 {
				isLcm = false
				break
			}
		}

		if isLcm {
			break
		}

		lcm = max * i
	}

	return lcm

}
