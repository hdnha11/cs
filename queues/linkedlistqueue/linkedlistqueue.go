package linkedlistqueue

import (
	"github.com/hdnha11/cs/arrays/linkedlist"
)

// LinkedListQueue implementation
type LinkedListQueue struct {
	items *linkedlist.LinkedList
}

// New create a new queue
func New() *LinkedListQueue {
	return &LinkedListQueue{linkedlist.New()}
}

// Enqueue adds value at position at tail
func (llq *LinkedListQueue) Enqueue(item interface{}) {
	llq.items.PushBack(item)
}

// Dequeue returns value and removes least recently added element (front)
func (llq *LinkedListQueue) Dequeue() interface{} {
	return llq.items.PopFront()
}

// Empty checks if the queue is empty
func (llq LinkedListQueue) Empty() bool {
	return llq.items.Empty()
}
