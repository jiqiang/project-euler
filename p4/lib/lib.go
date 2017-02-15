package lib

import "strconv"

func IsPalindrome(n int) bool {
  s1 := strconv.Itoa(n)

  if s1 == Reverse(s1) {
    return true
  }

  return false
}

func ProductOfTwoThreeDigits(n int)(int, int, bool) {
  d1, d2, ok := -1, -1, false
  for i := 999; i >= 100; i-- {
    if n % i == 0 && n / i >= 100 && n / i <= 999  {
      d1, d2, ok = i, n / i, true
      break
    }
  }
  return d1, d2, ok
}

func Reverse(s string) string {
  r := []rune(s)
  for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
    r[i], r[j] = r[j], r[i]
  }
  return string(r)
}
