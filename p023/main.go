package main

import (
	"fmt"
)

// Limit represents all integers which are greater than this limit can be
// written as the sum of two abundant numbers.
const Limit = 28123

// GetProperDivisors gets all non-negative divisors of n, not inlcluding itself.
func GetProperDivisors(n int) []int {
	divisors := []int{1}
	limit := n
	for i := 2; i < limit; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			limit = n / i
			if limit != i {
				divisors = append(divisors, limit)
			}
		}
	}
	return divisors
}

// GetSumOfIntSlice returns sum of numbers in an interger array.
func GetSumOfIntSlice(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}

// IsPerfectNumber returns true if the sum of all proper divisors of a given
// number is equal to that given number. Otherwise, return false.
func IsPerfectNumber(n int) bool {
	sum := GetSumOfDivisors(n)
	return sum == n
}

// IsDeficientNumber returns true if the sum of all proper divisors of a given
// number is less than that given number. Otherwise, return false.
func IsDeficientNumber(n int) bool {
	sum := GetSumOfDivisors(n)
	return sum < n
}

// IsAbundantNumber returns true if the sum of all proper divisors of a given
// number is greater than that given number. Otherwise, return false.
func IsAbundantNumber(n int) bool {
	sum := GetSumOfDivisors(n)
	return sum > n
}

// GetSumOfDivisors returns sum of proper divisors of a given number.
func GetSumOfDivisors(n int) int {
	divisors := GetProperDivisors(n)
	sum := GetSumOfIntSlice(divisors)
	return sum
}

// IsSumOfTwoAbundantNumbers checks if given number is sum of two abundant numbers.
func IsSumOfTwoAbundantNumbers(n int, abundantNums map[int]int) bool {
	// check if it is sum of two equal abundant numbers.
	if n%2 == 0 {
		_, ok := abundantNums[n/2]
		if ok {
			return true
		}
	}

	// check if it is sum of two not equal abundant numbers.
	for _, m := range abundantNums {
		i := n - m
		_, ok := abundantNums[i]
		if ok {
			return true
		}
	}

	return false
}

func main() {
	abundantNums := map[int]int{}
	sum := 0
	for i := 1; i <= Limit; i++ {
		if !IsSumOfTwoAbundantNumbers(i, abundantNums) {
			sum += i
		}
		if IsAbundantNumber(i) {
			abundantNums[i] = i
		}
	}
	fmt.Println(sum)
}
