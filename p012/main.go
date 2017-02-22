package main

import "fmt"
import "github.com/jiqiang/project-euler/eulerlibs"

func main() {
	count, sum := 0, 0

	for i := 1; ; i++ {
		sum = sum + i
		if eulerlibs.IsPrime2(sum) {
			continue
		}

		thisCount := eulerlibs.FindPrimeFactorsCount(sum)

		if count < thisCount {
			count = thisCount
		}

		if count > 500 {
			break
		}
	}

	fmt.Printf("%d is the first triangle number to have over five hundred divisors(%d)\n", sum, count)
}
