package libs

func SumOfMultiples(count int) int {
  sum := 0;

  for n := 1; n < count; n++ {
    if n % 3 == 0 || n % 5 == 0 {
      sum += n
    }
  }

  return sum
}
