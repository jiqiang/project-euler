package eulerlibs

import "testing"

func TestIsPrime2(t *testing.T) {
	expected := true
	actual := IsPrime2(5)
	if expected != actual {
		t.Errorf("expected:%v - actual:%v", expected, actual)
	}

	expected = false
	actual = IsPrime2(4)
	if expected != actual {
		t.Errorf("expected:%v - actual:%v", expected, actual)
	}
}

func BenchmarkRunP10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunP10(2000000)
	}
}
