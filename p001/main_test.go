package main

import "testing"

func TestIsMultiples(t *testing.T) {

	h := p001Helper{}

	var tests = []struct {
		target   int
		base     int
		expected bool
	}{
		{6, 3, true},
		{3, 3, true},
		{0, 3, false},
		{7, 3, false},
	}

	for _, tt := range tests {
		actual := h.isMultiples(tt.target, tt.base)
		if actual != tt.expected {
			t.Errorf("%d is not multiples of %d", tt.target, tt.base)
		}
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
