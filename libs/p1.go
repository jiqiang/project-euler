package libs

func SumOfMultiples(start int, end int) int {
  sum := 0;

  for n := start; n < end; n++ {
    if n % 3 == 0 || n % 5 == 0 {
      sum += n
    }
  }

  return sum
}
