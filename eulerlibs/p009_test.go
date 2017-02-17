package eulerlibs

import "testing"

func TestPythagoreanTriplet(t *testing.T) {
	expected := PythagoreanTriple{3, 4, 5}
	pts := PythagoreanTriplet(2)
	actual := pts[0]
	if expected.x != actual.x || expected.y != actual.y || expected.z != actual.z {
		t.Errorf("expected:%d,%d,%d - actual:%d,%d,%d", expected.x, expected.y, expected.z, actual.x, actual.y, actual.z)
	}
}

func TestScaleMatch(t *testing.T) {
	expected := false
	pt := PythagoreanTriple{3, 4, 5}
	actual := ScaleMatch(pt, 1000)
	if expected != actual {
		t.Errorf("ScaleMatch 3, 4, 5 should not match 1000")
	}

	expected = true
	pt = PythagoreanTriple{200, 375, 425}
	actual = ScaleMatch(pt, 1000)
	if expected != actual {
		t.Errorf("ScaleMatch 200, 375, 425 should match 1000")
	}
}

func TestP9(t *testing.T) {
	expected := 31875000
	_, _, _, actual := P9()
	if expected != actual {
		t.Errorf("P9() should return %d", expected)
	}
}

func BenchmarkP9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		P9()
	}
}
