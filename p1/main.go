/*
If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of all the
multiples of 3 or 5 below 1000.
*/

package main

import "fmt"
import "github.com/jiqiang/project-euler/libs"

func main() {
  results := make(chan int, 2)

  go func() {
    results <- libs.SumOfMultiples(1, 500)
  }()

  go func() {
    results <- libs.SumOfMultiples(500, 1000)
  }()

  fmt.Println(<-results + <-results)

}
