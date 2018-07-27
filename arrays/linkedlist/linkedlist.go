package linkedlist

import (
	"errors"
)

// LinkedList implementation
type LinkedList struct {
	head *listNode
	tail *listNode
	size int
}

// ErrIndexOutOfRange Index out of range errror
var ErrIndexOutOfRange = errors.New("Index out of range")

// New LinkedList
func New() *LinkedList {
	return new(LinkedList)
}

// Size returns number of data elements in list
func (ll LinkedList) Size() int {
	return ll.size
}

// Empty bool returns true if empty
func (ll LinkedList) Empty() bool {
	return ll.Size() == 0
}

// PushFront adds an item to the front of the list
func (ll *LinkedList) PushFront(item interface{}) {
	node := &listNode{item, ll.head}

	ll.head = node
	ll.size++

	if ll.tail == nil {
		ll.tail = node
	}
}

// PopFront remove front item and return its value
func (ll *LinkedList) PopFront() interface{} {
	if ll.head == nil {
		return nil
	}

	value := ll.head.value

	ll.head = ll.head.next
	ll.size--

	if ll.head == nil {
		ll.tail = nil
	}

	return value
}

// PushBack adds an item at the end
func (ll *LinkedList) PushBack(item interface{}) {
	node := &listNode{item, nil}

	if ll.tail != nil {
		ll.tail.next = node
	}

	ll.tail = node
	ll.size++

	if ll.head == nil {
		ll.head = node
	}
}

// PopBack removes end item and returns its value
func (ll *LinkedList) PopBack() interface{} {
	if ll.tail == nil {
		return nil
	}

	value := ll.tail.value

	node := ll.head.findTo(ll.tail)
	ll.tail = node
	ll.size--

	if node != nil {
		node.next = nil
	}

	if ll.tail == nil {
		ll.head = nil
	}

	return value
}

// Front get value of front item
func (ll LinkedList) Front() interface{} {
	if ll.head == nil {
		return nil
	}

	return ll.head.value
}

// Back get value of end item
func (ll LinkedList) Back() interface{} {
	if ll.tail == nil {
		return nil
	}

	return ll.tail.value
}

// Get returns the value of the nth item (starting at 0 for first)
func (ll LinkedList) Get(i int) (interface{}, error) {
	if i < 0 || i >= ll.Size() {
		return nil, ErrIndexOutOfRange
	}

	if i == 0 {
		return ll.head.value, nil
	}

	node := ll.head.findNth(i)

	return node.value, nil
}

// Insert insert value at index, so current item at that index is pointed to by new item at index
func (ll *LinkedList) Insert(i int, item interface{}) error {
	if i < 0 || i >= ll.Size() {
		return ErrIndexOutOfRange
	}

	if i == 0 {
		ll.PushFront(item)
		return nil
	}

	node := ll.head.findNth(i)
	prevNode := ll.head.findTo(node)
	newNode := &listNode{item, node}
	prevNode.next = newNode
	ll.size++

	return nil
}

// Erase removes node at given index
func (ll *LinkedList) Erase(i int) bool {
	if i < 0 || i >= ll.Size() {
		return false
	}

	if i == 0 {
		ll.PopFront()
		return true
	}

	if i == ll.Size()-1 {
		ll.PopBack()
		return true
	}

	node := ll.head.findNth(i)
	prevNode := ll.head.findTo(node)
	prevNode.next = node.next
	ll.size--

	return true
}

// ValueNthFromEnd returns the value of the node at nth position from the end of the list
func (ll LinkedList) ValueNthFromEnd(i int) interface{} {
	li := ll.Size() - (i + 1)

	if li < 0 || li >= ll.Size() {
		return nil
	}

	return ll.head.findNth(li).value
}

// Reverse reverses the list
func (ll *LinkedList) Reverse() {
	var prev *listNode
	node := ll.head

	for node != nil {
		next := node.next
		node.next = prev
		prev = node
		node = next
	}

	ll.head, ll.tail = ll.tail, ll.head
}

// RemoveValue removes the first item in the list with this value
func (ll *LinkedList) RemoveValue(item interface{}) bool {
	if ll.head == nil {
		return false
	}

	node := ll.head.findByValue(item)

	if node != nil {
		prevNode := ll.head.findTo(node)

		if prevNode == nil {
			ll.head = node.next
		} else {
			prevNode.next = node.next
		}

		if node == ll.tail {
			ll.tail = prevNode
		}

		ll.size--

		return true
	}

	return false
}
