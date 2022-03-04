package main

import (
	"testing"
)

// Tests have to start with prefix Test to get recognized by
// the go test tool. We also need a parameter of type *testing.T
// which offers testing functions. The typical order of instructions
// is test setup (data preparation), followed by a function call for
// the functionality to test, followed by a verification step.
func TestProduct(t *testing.T) {
	// Check very basic product.
	result := product(1, 2, 3)
	checkEqual(t, result, 6)

	// Check no factor corner case.
	result = product()
	checkEqual(t, result, 0)

	// Check with negative numbers.
	result = product(-7, 3, 10, 2)
	checkEqual(t, result, -420)

	// Check with one zero element.
	result = product(-3, 5, 20, 0)
	checkEqual(t, result, 0)
}

// Error methods are helpful for failing a test but continuing the test
// function.
func checkEqual(t *testing.T, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestRemove(t *testing.T) {
	slice := []int{4, 2, -3, 8, 0}

	// Remove last element.
	slice = remove(slice, 0)
	expected := []int{4, 2, -3, 8}
	checkEqualSlices(t, slice, expected)

	// Remove first element.
	slice = remove(slice, 4)
	expected = []int{2, -3, 8}
	checkEqualSlices(t, slice, expected)

	// Remove middle element.
	slice = remove(slice, -3)
	expected = []int{2, 8}
	checkEqualSlices(t, slice, expected)
}

// Fatal methods are helpful for failing a test when it makes no sense
// to continue the test function due to results that maybe built on each
// other. For simplicity we can also use reflect.DeepEqual to compare
// slices. But there is no builtin equality check for slices in golang.
func checkEqualSlices(t *testing.T, actual, expected []int) {
	t.Helper()

	if len(actual) != len(expected) {
		t.Fatalf("Expected slice %v but got %v", expected, actual)
	}

	for index := 0; index < len(actual); index++ {
		if actual[index] != expected[index] {
			t.Fatalf("Expected slice %v but got %v", expected, actual)
		}
	}
}
