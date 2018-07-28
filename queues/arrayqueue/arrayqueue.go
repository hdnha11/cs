package arrayqueue

import (
	"errors"
)

// ArrayQueue implementation
type ArrayQueue struct {
	items      []interface{}
	readIndex  int
	writeIndex int
}

// ErrNotEnoughSpace Not enough space error
var ErrNotEnoughSpace = errors.New("Not enough space error")

// New create a new queue
func New(cap int) *ArrayQueue {
	// We buffer 1 element to distinguish between empty and full queue (both case readIndex = writeIndex)
	return &ArrayQueue{make([]interface{}, cap+1), 0, 0}
}

// Enqueue adds item at end of available storage
func (aq *ArrayQueue) Enqueue(item interface{}) error {
	newWriteIndex := aq.safeIndex(aq.writeIndex + 1)

	if newWriteIndex == aq.readIndex {
		return ErrNotEnoughSpace
	}

	aq.items[aq.writeIndex] = item
	aq.writeIndex = newWriteIndex

	return nil
}

// Dequeue returns value and removes least recently added element
func (aq *ArrayQueue) Dequeue() interface{} {
	if aq.Empty() {
		return nil
	}

	item := aq.items[aq.readIndex]
	aq.readIndex = aq.safeIndex(aq.readIndex + 1)

	return item
}

// Empty checks if the queue is empty
func (aq ArrayQueue) Empty() bool {
	return aq.readIndex == aq.writeIndex
}

func (aq ArrayQueue) safeIndex(i int) int {
	return i % cap(aq.items)
}
