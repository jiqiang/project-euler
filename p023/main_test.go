package main

import "testing"
import "reflect"

func TestGetProperDivisors(t *testing.T) {
	var tests = []struct {
		target   int
		expected []int
	}{
		{1, []int{1}},
		{2, []int{1}},
		{3, []int{1}},
		{4, []int{1, 2}},
		{5, []int{1}},
		{6, []int{1, 2, 3}},
		{10, []int{1, 2, 5}},
		{12, []int{1, 2, 6, 3, 4}},
		{28, []int{1, 2, 14, 4, 7}},
		{16, []int{1, 2, 8, 4}},
	}
	for _, tt := range tests {
		actual := GetProperDivisors(tt.target)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("expecting %v, but actually get %v", tt.expected, actual)
		}
	}
}

func TestGetSumOfIntSlice(t *testing.T) {
	var tests = []struct {
		target   []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{2}, 2},
		{[]int{2, 3, 4}, 9},
	}
	for _, tt := range tests {
		actual := GetSumOfIntSlice(tt.target)
		if actual != tt.expected {
			t.Errorf("expecting %v, but actually get %v", tt.expected, actual)
		}
	}
}

func TestGetSumOfDivisors(t *testing.T) {
	var tests = []struct {
		target   int
		expected int
	}{
		{4, 3},
	}
	for _, tt := range tests {
		actual := GetSumOfDivisors(tt.target)
		if actual != tt.expected {
			t.Errorf("expecting %v, but actually get %v", tt.expected, actual)
		}
	}
}

func TestIsPerfectNumber(t *testing.T) {
	var tests = []struct {
		target   int
		expected bool
	}{
		{26, false},
		{28, true},
	}
	for _, tt := range tests {
		actual := IsPerfectNumber(tt.target)
		if actual != tt.expected {
			t.Errorf("expecting %v, but actually get %v", tt.expected, actual)
		}
	}
}

func TestIsDeficientNumber(t *testing.T) {
	var tests = []struct {
		target   int
		expected bool
	}{
		{26, true},
		{28, false},
		{12, false},
	}
	for _, tt := range tests {
		actual := IsDeficientNumber(tt.target)
		if actual != tt.expected {
			t.Errorf("expecting %v, but actually get %v", tt.expected, actual)
		}
	}
}

func TestIsAbundantNumber(t *testing.T) {
	var tests = []struct {
		target   int
		expected bool
	}{
		{26, false},
		{28, false},
		{12, true},
	}
	for _, tt := range tests {
		actual := IsAbundantNumber(tt.target)
		if actual != tt.expected {
			t.Errorf("expecting %v, but actually get %v", tt.expected, actual)
		}
	}
}

func TestIsSumOfTwoAbundantNumbers(t *testing.T) {
	var tests = []struct {
		target          int
		abundantNumbers map[int]int
		expected        bool
	}{
		{23, map[int]int{12: 12, 18: 18, 20: 20}, false},
		{24, map[int]int{12: 12, 18: 18, 20: 20}, true},
		{30, map[int]int{12: 12, 18: 18, 20: 20}, true},
	}
	for _, tt := range tests {
		actual := IsSumOfTwoAbundantNumbers(tt.target, tt.abundantNumbers)
		if actual != tt.expected {
			t.Errorf("expecting %v, but actually get %v", tt.expected, actual)
		}
	}
}
