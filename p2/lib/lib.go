package lib

func EvenFibonacci(term int) []int {
  var evenFibs []int

  a, b := 1, 2

  for {
    if a > term {
      break;
    }

    if a % 2 == 0 {
      evenFibs = append(evenFibs, a)
    }

    a, b = b, a + b
  }

  return evenFibs
}
