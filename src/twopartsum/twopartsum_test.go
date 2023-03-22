package main

import (
	"testing"
)

// test that ConcurrentSum sums an even-length array correctly
func TestSumConcurrentCorrectlySumsEvenArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 55

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

// test that ConcurrentSum sums an odd-lenght array correctly

func TestSumConcurrentCorrectlySumsOddArray(t *testing.T) {
	arr := []int{3, 6, 1, 9, 8, 50, 87}
	expected := 164

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

func TestSumConcurrentCorrectlySumsNegativeIntegers(t *testing.T) {
	arr := []int{-7, -10, -15, -11, - 3}
	expected := -46

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}