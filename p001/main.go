// Multiples of 3 and 5
//
// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of all the
// multiples of 3 or 5 below 1000.

package main

import (
	"fmt"
	"math"
)

type helper interface {
	isMultiples(target int, base int) bool
}

type p001Helper struct{}

// isMultiples check if given target is multiples of given base.
func (s p001Helper) isMultiples(target int, base int) bool {
	if target < base {
		return false
	}
	return math.Mod(float64(target), float64(base)) == float64(0)
}

// run runs p001 solution
func run(target int, b1 int, b2 int) int {

	h := p001Helper{}

	i := b1
	if b1 > b2 {
		i = b2
	}

	sum := 0
	for {
		if i >= target {
			break
		}

		if h.isMultiples(i, b1) || h.isMultiples(i, b2) {
			sum += i
		}
		i++
	}

	return sum
}

func main() {
	fmt.Println(run(1000, 3, 5))
}
