// Factorial digit sum
//
// n! means n x (n - 1) x ... x 3 x 2 x 1
//
// For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
//
// Find the sum of the digits in the number 100!
package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// Factorial takes a non-negative integer n and returns n x (n - 1)
func Factorial(n uint64) *big.Int {
	// 0! = 1
	if n == uint64(0) {
		return big.NewInt(1)
	}
	return big.NewInt(0).Mul(big.NewInt(int64(n)), Factorial(n-1))
}

// DigitArray takes a non-negative integer and returns an array containing all
// digits of that integer
func DigitArray(n *big.Int) []uint {
	da := []uint{}
	str := n.String()
	for _, v := range str {
		d, _ := strconv.Atoi(string(v))
		da = append(da, uint(d))
	}
	return da
}

// SumOfDigits takes an integer array and returns sum of all digits in that array.
func SumOfDigits(digits []uint) uint64 {
	sum := uint64(0)
	for _, d := range digits {
		sum += uint64(d)
	}
	return sum
}

func main() {
	number := uint64(100)
	f := Factorial(number)
	fArray := DigitArray(f)
	fSum := SumOfDigits(fArray)
	fmt.Println(fSum)
}
