/*
A palindromic number reads the same both ways. The largest palindrome made from
the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import "fmt"
import "github.com/jiqiang/project-euler/p4/lib"

func main() {

  lowBoundary := 100 * 100
  highBoundary := 999 * 999

  palindrome, d1, d2, ok := -1, -1, -1, false

  for n := highBoundary; n >= lowBoundary; n-- {

    if !lib.IsPalindrome(n) {
      continue
    }

    d1, d2, ok = lib.ProductOfTwoThreeDigits(n)

    if ok == true {
      palindrome = n
      break
    }

  }

  fmt.Printf("%d x %d = %d", d1, d2, palindrome)
  fmt.Println()

}
