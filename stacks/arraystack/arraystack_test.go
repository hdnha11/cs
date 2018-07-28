package arraystack_test

import (
	"testing"

	"github.com/hdnha11/cs/stacks/arraystack"
)

func TestNew(t *testing.T) {
	as := arraystack.New(2)

	if as == nil {
		t.Error("Should return an not nil stack")
	}

	if !as.Empty() {
		t.Error("Expected empty stack")
	}
}

func TestPushPop(t *testing.T) {
	as := arraystack.New(0)

	if item := as.Pop(); item != nil {
		t.Error("Expected nil value")
	}

	err := as.Push(1)

	if err == nil {
		t.Error("Expected ErrNotEnoughSpace")
	}

	as = arraystack.New(4)

	as.Push(1)

	if as.Empty() {
		t.Error("Expected the stack is not empty")
	}

	if item := as.Pop(); item != 1 {
		t.Error("Expected item value = 1 but got", item)
	}

	if !as.Empty() {
		t.Error("Expected an empty stack")
	}

	as.Push(2)
	as.Push(3)
	as.Push(4)

	if item := as.Pop(); item != 4 {
		t.Error("Expected item value = 4 but got", item)
	}

	if as.Empty() {
		t.Error("Expected the stack is not empty")
	}

	if item := as.Pop(); item != 3 {
		t.Error("Expected item value = 3 but got", item)
	}

	if as.Empty() {
		t.Error("Expected the stack is not empty")
	}

	if item := as.Pop(); item != 2 {
		t.Error("Expected item value = 2 but got", item)
	}

	if !as.Empty() {
		t.Error("Expected an empty stack")
	}
}
