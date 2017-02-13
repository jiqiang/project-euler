package libs

import "testing"

func TestEvenFibonacci(t *testing.T) {
  expected := []int{2, 8}
  actual := EvenFibonacci(10)

  if len(expected) != len(actual) {
    t.Errorf("length of expected:%d - length of actual:%d", expected, actual)
  }

  for i := 0; i < len(expected); i++ {
    if expected[i] != actual[i] {
      t.Errorf("expected[%d]=%d - actual[%d]=%d", i, expected[i], i, actual[i])
    }
  }
}
