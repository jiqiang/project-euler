package lib

import "testing"

func TestSumOfMultiples(t *testing.T) {
  expected := 23
  actual := SumOfMultiples(1, 10)
  if actual != expected {
    t.Errorf("(p1) expected:%d - actual:%d", expected, actual)
  }
}
