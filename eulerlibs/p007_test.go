package eulerlibs

import "testing"

func TestIsPrime(t *testing.T) {
	target := 7
	expected := true
	actual := IsPrime(target)
	if expected != actual {
		t.Errorf("%d is a prime", target)
	}

	target = 8
	expected = false
	actual = IsPrime(target)
	if expected != actual {
		t.Errorf("%d is not a prime", target)
	}
}

func BenchmarkP7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunP7(10001)
	}
}
