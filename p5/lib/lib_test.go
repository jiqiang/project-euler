package lib

import "testing"

func TestLeastCommonMultiple(t *testing.T) {
  expected := 60
  actual := LeastCommonMultiple(6)
  if expected != actual {
    t.Errorf("least common multiple of 1 to 6 should be %d", expected)
  }
}
