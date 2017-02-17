package eulerlibs

import "testing"

func TestEvenFibonacci(t *testing.T) {
	expected := []int{2, 8}
	actual := EvenFibonacci(10)

	if len(expected) != len(actual) {
		t.Errorf("(p2) length of expected:%d - length of actual:%d", expected, actual)
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("(p2) expected[%d]=%d - actual[%d]=%d", i, expected[i], i, actual[i])
		}
	}
}

func BenchmarkP2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunP2(4000000)
	}
}
