package eulerlibs

import "testing"

func TestLeastCommonMultiple(t *testing.T) {
	expected := 60
	actual := LeastCommonMultiple(6)
	if expected != actual {
		t.Errorf("least common multiple of 1 to 6 should be %d", expected)
	}
}

func BenchmarkLeastCommonMultiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeastCommonMultiple(20)
	}
}
