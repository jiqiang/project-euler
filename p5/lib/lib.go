package lib

import "fmt"

func EvenlyDivisibleNumber(from int, to int) int {

  value := 1
  for j := 2; j <= 20; j++ {
    if IsPrimeNumber(j) {
      value = value * j
    }
  }

  fmt.Println(value)

  return 1
}

func IsPrimeNumber(n int) bool {
  for i := 2; i < n; i++ {
    if n % i == 0 {
      return false
    }
  }
  return true
}
