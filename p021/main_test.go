package main

import "testing"
import "reflect"

func TestSumOfSlice(t *testing.T) {
	var tests = []struct {
		target   []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 3},
	}

	for _, tt := range tests {
		actual := sumOfSlice(tt.target)
		if actual != tt.expected {
			t.Errorf("sum of %v shouldn't be equal to %d", tt.target, actual)
		}
	}
}

func TestDivisors(t *testing.T) {
	var tests = []struct {
		target   int
		expected map[int]int
	}{
		{220, map[int]int{1: 220, 2: 110, 4: 55, 5: 44, 10: 22, 11: 20}},
		{284, map[int]int{1: 284, 2: 142, 4: 71}},
	}

	for _, tt := range tests {
		actual := divisors(tt.target)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("expecting divisors of %d to be %v, actually got %v", tt.target, tt.expected, actual)
		}
	}
}
