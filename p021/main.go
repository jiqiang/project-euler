// Amicable numbers
//
// Let d(n) be defined as the sum of proper divisors of n (numbers less than n
// which divide evenly into n). If d(a) = b and d(b) = a, where a ≠ b, then a
// and b are an amicable pair and each of a and b are called amicable numbers.
//
// For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44,
// 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4,
// 71 and 142; so d(284) = 220.
//
// Evaluate the sum of all the amicable numbers under 10000.
package main

import "fmt"

func sumOfSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func divisors(n int) []int {
	ds := []int{}
	for i := 1; i < n; i++ {
		if contains(ds, i) {
			continue
		}
		m := n % i
		if m == 0 {
			if i == 1 {
				ds = append(ds, i)
			} else if i == n/i {
				ds = append(ds, i)
			} else {
				ds = append(ds, i, n/i)
			}
		}
	}
	return ds
}

func contains(ds []int, n int) bool {
	for _, v := range ds {
		if v == n {
			return true
		}
	}
	return false
}

func sumOfDivisors(n int) int {
	ds := divisors(n)
	sum := sumOfSlice(ds)
	return sum
}

func run(n int) int {
	s := []int{}
	for i := 1; i < n; i++ {
		if contains(s, i) {
			continue
		}
		j := sumOfDivisors(i)
		if sumOfDivisors(j) == i && i != j {
			s = append(s, i, j)
		}
	}
	return sumOfSlice(s)
}

func run2(n int) int {
	c := make(chan int)
	sum := 0
	for i := 1; i < n; i++ {
		go func(cc chan int, ii int) {
			j := sumOfDivisors(ii)
			if sumOfDivisors(j) == ii && ii != j {
				c <- ii
			} else {
				c <- 0
			}
		}(c, i)
	}

	for i := 1; i < n; i++ {
		sum += <-c
	}

	return sum
}

func main() {
	fmt.Println(run2(10000))
}
