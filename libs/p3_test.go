package libs

import "testing"

func TestMaxPrimeFactor(t *testing.T) {
  expected := 5
  actual := MaxPrimeFactor(10)
  if expected != actual {
    t.Errorf("(p3) expected:%d - actual:%d", expected, actual)
  }
}
