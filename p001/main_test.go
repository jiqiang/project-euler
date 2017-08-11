package main

import "testing"

func TestIsMultiples(t *testing.T) {
	var actual, expected bool
	var base, target int

	h := p001Helper{}

	target = 6
	base = 3
	expected = true
	actual = h.isMultiples(target, base)
	if actual != expected {
		t.Errorf("%d is not multiples of %d", target, base)
	}

	target = 3
	base = 3
	expected = true
	actual = h.isMultiples(target, base)
	if actual != expected {
		t.Errorf("%d is not multiples of %d", target, base)
	}

	target = 0
	base = 3
	expected = false
	actual = h.isMultiples(target, base)
	if actual != expected {
		t.Errorf("%d is not multiples of %d", target, base)
	}

	target = 7
	base = 3
	expected = false
	actual = h.isMultiples(target, base)
	if actual != expected {
		t.Errorf("%d is not multiples of %d", target, base)
	}
}

func TestRun(t *testing.T) {
	var actual, expected, target int

	expected = 23
	target = 10
	actual = run(target, 3, 5)
	if actual != expected {
		t.Errorf("p001 takes %d as input and should return %d, it actually returns %d", target, expected, actual)
	}
}
