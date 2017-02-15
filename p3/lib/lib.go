package lib

func MaxPrimeFactor(num int) int {
  for i := 2; i < num; i++ {
    if num % i == 0 {
      num = num / i
      i = i - 1
    }
  }
  return num
}
