package main

import "testing"
import "reflect"

func TestExtractNamesFromString(t *testing.T) {
	var tests = []struct {
		target   string
		expected []string
	}{
		{``, []string{}},
		{` `, []string{}},
		{`"GLENN"`, []string{"GLENN"}},
		{`"GLENN", "AUSTYN"`, []string{"AUSTYN", "GLENN"}},
		{`GLENN JOY AUSTYN`, []string{"LENN JOY AUSTY"}},
	}
	for _, tt := range tests {
		actual := ExtractNamesFromString(tt.target)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("expecting %v, but actually get %v", tt.expected, actual)
		}
	}
}

func TestGetContentStringFromFile(t *testing.T) {
	var tests = []struct {
		filename string
		expected string
	}{
		{"names1_test.txt", `"GLENN", "AUSTYN"`},
	}
	for _, tt := range tests {
		actual, _ := GetContentStringFromFile(tt.filename)
		if actual != tt.expected {
			t.Errorf("expecting %v, but actually get %v", tt.expected, actual)
		}
	}
}

func TestGetLetterPositionSum(t *testing.T) {
	var tests = []struct {
		target   string
		expected int64
	}{
		{"COLIN", 53},
		{"", 0},
	}
	for _, tt := range tests {
		actual := GetLetterPositionSum(tt.target)
		if actual != tt.expected {
			t.Errorf("expecting %v, but actually get %v", tt.expected, actual)
		}
	}
}
