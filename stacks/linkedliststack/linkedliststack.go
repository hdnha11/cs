package linkedliststack

import (
	"github.com/hdnha11/cs/arrays/linkedlist"
)

// LinkedListStack stack implementation using array
type LinkedListStack struct {
	items *linkedlist.LinkedList
}

// New create new LinkedListStack
func New() *LinkedListStack {
	return &LinkedListStack{linkedlist.New()}
}

// Push push an item to top of the stack
func (lls *LinkedListStack) Push(item interface{}) {
	lls.items.PushFront(item)
}

// Pop remove and return an item from top of the stack
func (lls *LinkedListStack) Pop() interface{} {
	return lls.items.PopFront()
}

// Empty check if the stack is empty
func (lls LinkedListStack) Empty() bool {
	return lls.items.Empty()
}
