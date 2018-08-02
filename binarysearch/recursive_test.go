package binarysearch_test

import (
	"testing"

	"github.com/hdnha11/cs/binarysearch"
)

func TestRecursiveSearchFound(t *testing.T) {
	for i, el := range array {
		found := binarysearch.Recursive(array, el)
		if found != i {
			t.Error("Expected", found, "equals", i)
		}
	}
}

func TestRecursiveSearchNotFound(t *testing.T) {
	found := binarysearch.Recursive(array, 100)

	if found != -1 {
		t.Error("Expected not found")
	}

	found = binarysearch.Recursive(array, -1)

	if found != -1 {
		t.Error("Expected not found")
	}
}

func TestRecursiveSearchTwoItemsArray(t *testing.T) {
	tia := []int{2, 4}

	found := binarysearch.Recursive(tia, tia[0])

	if found != 0 {
		t.Error("Expected found 0 but got", found)
	}

	found = binarysearch.Recursive(tia, tia[1])

	if found != 1 {
		t.Error("Expected found 1 but got", found)
	}

	found = binarysearch.Recursive(tia, 1)

	if found != -1 {
		t.Error("Expected not found")
	}

	found = binarysearch.Recursive(tia, 5)

	if found != -1 {
		t.Error("Expected not found")
	}
}
