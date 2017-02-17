package eulerlibs

import "testing"

func TestMaxPrimeFactor(t *testing.T) {
	expected := 5
	actual := MaxPrimeFactor(10)
	if expected != actual {
		t.Errorf("(p3) expected:%d - actual:%d", expected, actual)
	}
}

func BenchmarkPrimeFactor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxPrimeFactor(600851475143)
	}
}
