package linkedliststack_test

import (
	"testing"

	"github.com/hdnha11/cs/stacks/linkedliststack"
)

func TestNew(t *testing.T) {
	lls := linkedliststack.New()

	if lls == nil {
		t.Error("Should return an not nil stack")
	}

	if !lls.Empty() {
		t.Error("Expected empty stack")
	}
}

func TestPushPop(t *testing.T) {
	lls := linkedliststack.New()

	if item := lls.Pop(); item != nil {
		t.Error("Expected nil value")
	}

	lls.Push(1)

	if lls.Empty() {
		t.Error("Expected the stack is not empty")
	}

	if item := lls.Pop(); item != 1 {
		t.Error("Expected item value = 1 but got", item)
	}

	if !lls.Empty() {
		t.Error("Expected an empty stack")
	}

	lls.Push(2)
	lls.Push(3)
	lls.Push(4)

	if item := lls.Pop(); item != 4 {
		t.Error("Expected item value = 4 but got", item)
	}

	if lls.Empty() {
		t.Error("Expected the stack is not empty")
	}

	if item := lls.Pop(); item != 3 {
		t.Error("Expected item value = 3 but got", item)
	}

	if lls.Empty() {
		t.Error("Expected the stack is not empty")
	}

	if item := lls.Pop(); item != 2 {
		t.Error("Expected item value = 2 but got", item)
	}

	if !lls.Empty() {
		t.Error("Expected an empty stack")
	}
}
