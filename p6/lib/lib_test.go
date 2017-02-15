package lib

import "testing"

func TestSumOfSquares(t *testing.T) {
  expected := 385
  actual := SumOfSquares(10)
  if expected != actual {
    t.Errorf("expected:%d - actual:%d", expected, actual)
  }
}

func TestSquareOfSum(t *testing.T) {
  expected := 3025
  actual := SquareOfSum(10)
  if expected != actual {
    t.Errorf("expected:%d - actual:%d", expected, actual)
  }
}
