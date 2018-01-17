/*
The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	f1, f2, count := big.NewInt(1), big.NewInt(1), 2
	var limit big.Int
	// the smallest number contains 1000 digit is 10^999
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)
	for f2.Cmp(&limit) < 0 {
		f1, f2 = f2, f1.Add(f1, f2)
		count++
	}
	fmt.Println(count)
}
