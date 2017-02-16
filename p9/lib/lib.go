package lib

// PythagoreanTriple is a function.
type PythagoreanTriple struct {
	x int
	y int
	z int
}

// P9 is a function
func P9() (int, int, int, int) {

	var theOne PythagoreanTriple
	var find bool

	for n := 2; ; n += 2 {
		pts := PythagoreanTriplet(n)
		find = false
		for _, pt := range pts {
			if 1000%(pt.x+pt.y+pt.z) != 0 {
				continue
			}

			if ScaleMatch(pt, 1000) {
				theOne = pt
				find = true
				break
			}
		}
		if find {
			break
		}
	}

	return theOne.x, theOne.y, theOne.z, theOne.x * theOne.y * theOne.z
}

// PythagoreanTriplet uses Dikson's method
// https://en.wikipedia.org/wiki/Formulas_for_generating_Pythagorean_triples
func PythagoreanTriplet(r int) []PythagoreanTriple {
	var pts []PythagoreanTriple
	var pt PythagoreanTriple
	var n, s, t, x, y, z int

	n = r * r / 2
	for s = 1; s < n; s++ {
		if n%s == 0 {
			t = n / s
			x = r + s
			y = r + t
			z = r + s + t
			pt = PythagoreanTriple{x, y, z}
			pts = append(pts, pt)
		}
	}

	return pts
}

// ScaleMatch is a function
func ScaleMatch(pt PythagoreanTriple, sum int) bool {
	for i := 2; ; i++ {

		if pt.x+pt.y+pt.z > sum {
			break
		}

		if pt.x+pt.y+pt.z == sum {
			return true
		}

		pt.x *= i
		pt.y *= i
		pt.z *= i

	}

	return false
}
