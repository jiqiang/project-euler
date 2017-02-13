package libs

func EvenFibonacci(term int) <-chan int {
  out := make(chan int)

  a, b := 1, 2
  go func() {
    for {
      if a > term {
        break;
      }

      if a % 2 != 0 {
        out <- a
      }

      a, b = b, a + b
    }
    close(out)
  }()

  return out
}
