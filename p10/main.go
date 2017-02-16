package main

import "fmt"

func IsPrime2(n int) bool {
  if n <= 1 {
    return false
  } else if n <= 3 {
    return true
  } else if n % 2 == 0 || n % 3 == 0 {
    return false
  }

  for i := 5; i * i <= n; i += 6 {
    if n % i == 0 || n % (i + 2) == 0 {
      return false
    }
  }
  return true
}

func main() {

  sum := 2
  for i := 3; i <= 2000000; i += 2 {
    if IsPrime2(i) {
      sum += i
    }
  }

  fmt.Println(sum)

}
