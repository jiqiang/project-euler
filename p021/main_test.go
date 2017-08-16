package main

import "testing"
import "reflect"
import "sort"

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
		expected []int
	}{
		{220, []int{1, 2, 4, 5, 10, 11, 20, 22, 44, 55, 110}},
		{284, []int{1, 2, 4, 71, 142}},
		{1, []int{}},
		{2, []int{1}},
		{3, []int{1}},
		{4, []int{1, 2}},
		{5, []int{1}},
		{6, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		actual := divisors(tt.target)
		sort.Ints(actual)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("expecting divisors of %d to be %v, actually got %v", tt.target, tt.expected, actual)
		}
	}
}

func TestContains(t *testing.T) {
	var tests = []struct {
		divisors []int
		target   int
		expected bool
	}{
		{[]int{}, 1, false},
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 4, false},
	}

	for _, tt := range tests {
		actual := contains(tt.divisors, tt.target)
		if actual != tt.expected {
			t.Errorf("%v doesn't contain %d", tt.divisors, tt.target)
		}
	}
}

func TestSumOfDivisors(t *testing.T) {
	var tests = []struct {
		target   int
		expected int
	}{
		{220, 284},
		{284, 220},
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 3},
	}

	for _, tt := range tests {
		actual := sumOfDivisors(tt.target)
		if actual != tt.expected {
			t.Errorf("sum of divisor %d should be equal to %d, it is actually equal to %d", tt.target, tt.expected, actual)
		}
	}
}

func TestRun(t *testing.T) {
	var tests = []struct {
		target   int
		expected int
	}{
		{285, 504},
		{1211, 2898},
	}

	for _, tt := range tests {
		actual := run(tt.target)
		if actual != tt.expected {
			t.Errorf("expecting %d, actually got %d", tt.expected, actual)
		}
	}
}

func TestRun2(t *testing.T) {
	var tests = []struct {
		target   int
		expected int
	}{
		{285, 504},
		{1211, 2898},
	}

	for _, tt := range tests {
		actual := run2(tt.target)
		if actual != tt.expected {
			t.Errorf("expecting %d, actually got %d", tt.expected, actual)
		}
	}
}
