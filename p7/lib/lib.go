package lib

import "math"

func IsPrime(n int) bool {
  if n == 2 {
    return true
  } else if n % 2 == 0 {
    return false
  }

  r := int(math.Sqrt(float64(n))) + 1
  for i := 3; i < r; i++ {
    if n % i == 0 {
      return false
    }
  }
  return true
}

func P7(n int) int {
  var p, i, count int

  i, count = 3, 1

  for {
    if IsPrime(i) {
      count++
    }

    if count == n {
      p = i
      break
    }

    i += 2
  }

  return p
}
