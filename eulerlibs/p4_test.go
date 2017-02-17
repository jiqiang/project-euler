package eulerlibs

import "testing"

func TestIsPalindrome(t *testing.T) {
	t1 := IsPalindrome(10001)
	if t1 == false {
		t.Errorf("(p4) 10001 should be a palindrome")
	}

	t2 := IsPalindrome(1000)
	if t2 == true {
		t.Errorf("(p4) 1000 shouldn't be a palindrome")
	}
}

func TestProductOfTwoThreeDigits(t *testing.T) {
	var d1, d2 int
	var ok bool

	d1, d2, ok = ProductOfTwoThreeDigits(10001)
	if d1 != -1 && d2 != -1 && ok != false {
		t.Errorf("(p4) ProductOfTwoThreeDigits(10001) should return -1, -1 and false")
	}

	d1, d2, ok = ProductOfTwoThreeDigits(906609)
	if d1 != 993 && d2 != 913 && ok != true {
		t.Errorf("(p4) ProductOfTwoThreeDigits(906609) should return 993, 913 and true")
	}
}

func TestReverse(t *testing.T) {
	expected := "54321"
	actual := Reverse("12345")
	if expected != actual {
		t.Errorf("the reverse string of '12345' should be '54321'")
	}
}

func BenchmarkP4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunP4()
	}
}
