package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	var expected, actual *big.Int
	var target uint64

	// test factorial of 0
	target = uint64(0)
	expected = big.NewInt(1)
	actual = Factorial(target)
	if actual.Cmp(expected) != 0 {
		t.Errorf("factorial of %d should be %d, actually get %d", target, expected, actual)
	}

	// test factorial of 1
	target = uint64(1)
	expected = big.NewInt(1)
	actual = Factorial(target)
	if actual.Cmp(expected) != 0 {
		t.Errorf("factorial of %d should be %d, actually get %d", target, expected, actual)
	}

	// test factorial of 10
	target = uint64(10)
	expected = big.NewInt(3628800)
	actual = Factorial(target)
	if actual.Cmp(expected) != 0 {
		t.Errorf("factorial of %d should be %d, actually get %d", target, expected, actual)
	}
}

func TestDigitArray(t *testing.T) {
	var actual, expected []uint
	var target *big.Int

	// test 0
	target = big.NewInt(0)
	expected = []uint{0}
	actual = DigitArray(target)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("DigitArray takes %d and should return %v, actually get %v", target, expected, actual)
	}

	// test 987654321
	target = big.NewInt(9876543210)
	expected = []uint{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	actual = DigitArray(target)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("DigitArray takes %d and should return %v, actually get %v", target, expected, actual)
	}
}

func TestSumOfDigits(t *testing.T) {
	var actual, expected uint64
	var target []uint

	// test []
	target = []uint{}
	expected = uint64(0)
	actual = SumOfDigits(target)
	if actual != expected {
		t.Errorf("SumOfDigits takes %v and should return %d, actually get %d", target, expected, actual)
	}

	// test [2]
	target = []uint{2}
	expected = uint64(2)
	actual = SumOfDigits(target)
	if actual != expected {
		t.Errorf("SumOfDigits takes %v and should return %d, actually get %d", target, expected, actual)
	}

	// test [1 2 3]
	target = []uint{1, 2, 3}
	expected = uint64(6)
	actual = SumOfDigits(target)
	if actual != expected {
		t.Errorf("SumOfDigits takes %v and should return %d, actually get %d", target, expected, actual)
	}
}
