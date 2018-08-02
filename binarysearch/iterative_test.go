package binarysearch_test

import (
	"testing"

	"github.com/hdnha11/cs/binarysearch"
)

var array = []int{0, 5, 13, 19, 22, 41, 55, 68, 72, 81, 98}

func TestIterativeSearchFound(t *testing.T) {
	for i, el := range array {
		found := binarysearch.Iterative(array, el)
		if found != i {
			t.Error("Expected", found, "equals", i)
		}
	}
}

func TestIterativeSearchNotFound(t *testing.T) {
	found := binarysearch.Iterative(array, 100)

	if found != -1 {
		t.Error("Expected not found")
	}

	found = binarysearch.Iterative(array, -1)

	if found != -1 {
		t.Error("Expected not found")
	}
}

func TestIterativeSearchTwoItemsArray(t *testing.T) {
	tia := []int{2, 4}

	found := binarysearch.Iterative(tia, tia[0])

	if found != 0 {
		t.Error("Expected found 0 but got", found)
	}

	found = binarysearch.Iterative(tia, tia[1])

	if found != 1 {
		t.Error("Expected found 1 but got", found)
	}

	found = binarysearch.Iterative(tia, 1)

	if found != -1 {
		t.Error("Expected not found")
	}

	found = binarysearch.Iterative(tia, 5)

	if found != -1 {
		t.Error("Expected not found")
	}
}
