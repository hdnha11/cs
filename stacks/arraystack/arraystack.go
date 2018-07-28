package arraystack

import (
	"errors"
)

// ArrayStack stack implementation using array
type ArrayStack struct {
	items []interface{}
	size  int
}

// ErrNotEnoughSpace Not enough space error
var ErrNotEnoughSpace = errors.New("Not enough space error")

// New create new ArrayStack
func New(cap int) *ArrayStack {
	return &ArrayStack{make([]interface{}, cap), 0}
}

// Push push an item to top of the stack
func (as *ArrayStack) Push(item interface{}) error {
	if as.size >= cap(as.items) {
		return ErrNotEnoughSpace
	}

	as.items[as.size] = item
	as.size++

	return nil
}

// Pop remove and return an item from top of the stack
func (as *ArrayStack) Pop() interface{} {
	if as.Empty() {
		return nil
	}

	as.size--

	return as.items[as.size]
}

// Empty check if the stack is empty
func (as ArrayStack) Empty() bool {
	return as.size == 0
}
