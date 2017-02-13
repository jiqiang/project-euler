package libs

import "testing"

func TestSumOfMultiples(t *testing.T) {
  expected := 23
  actual := SumOfMultiples(10)
  if actual != expected {
    t.Errorf("expected:%d - actual:%d", expected, actual)
  }
}
