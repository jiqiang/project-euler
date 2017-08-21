package main

import "testing"

func TestSumOfEvenFibs(t *testing.T) {
	var tests = []struct {
		max      int
		expected int
	}{
		{1, 0},
		{4, 2},
		{5, 2},
		{10, 10},
	}

	for _, tt := range tests {
		var actual = sumOfEvenFibs(tt.max)
		if actual != tt.expected {
			t.Errorf("sum of even fibonacci number less than %d should be %d, actually got %d", tt.max, tt.expected, actual)
		}
	}
}
