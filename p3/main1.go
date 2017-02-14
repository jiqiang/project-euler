/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import "fmt"

func main() {

  start := 600851475143

  for i := 2; i < start; i++ {
    if start % i == 0 {
      start = start / i;
      i = i - 1
    }
  }

  fmt.Println(start)

}
