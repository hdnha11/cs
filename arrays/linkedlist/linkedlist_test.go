package linkedlist_test

import (
	"testing"

	"github.com/hdnha11/cs/arrays/linkedlist"
)

func TestNewLinkedList(t *testing.T) {
	ll := linkedlist.New()

	if ll == nil {
		t.Error("Expected a LinkedList was returned")
	}

	if ll.Size() != 0 {
		t.Error("Expected new LinkedList size = 0 but got", ll.Size())
	}
}

func TestEmpty(t *testing.T) {
	ll := linkedlist.New()

	if !ll.Empty() {
		t.Error("Expected new LinkedList empty")
	}
}

func TestPushFront(t *testing.T) {
	ll := linkedlist.New()

	ll.PushFront(1)
	ll.PushFront(2)
	ll.PushFront(3)

	if ll.Size() != 3 {
		t.Error("Expected list size = 3 but got", ll.Size())
	}

	item, _ := ll.Get(0)

	if item != 3 {
		t.Error("Expected item value at 0 = 3 but got", item)
	}

	item, _ = ll.Get(1)

	if item != 2 {
		t.Error("Expected item value at 0 = 2 but got", item)
	}

	item, _ = ll.Get(2)

	if item != 1 {
		t.Error("Expected item value at 0 = 1 but got", item)
	}

	_, err := ll.Get(3)

	if err == nil {
		t.Error("Expected ErrIndexOutOfRange error")
	}
}

func TestPopFront(t *testing.T) {
	ll := linkedlist.New()

	item := ll.PopFront()

	if item != nil {
		t.Error("Expected item value = nil", item)
	}

	ll.PushFront(1)
	ll.PushFront(2)
	ll.PushFront(3)

	item = ll.PopFront()

	if ll.Size() != 2 {
		t.Error("Expected size = 2 but got", ll.Size())
	}

	if item != 3 {
		t.Error("Expected item value = 3 but got", item)
	}

	item = ll.PopFront()

	if ll.Size() != 1 {
		t.Error("Expected size = 1 but got", ll.Size())
	}

	if item != 2 {
		t.Error("Expected item value = 2 but got", item)
	}

	item = ll.PopFront()

	if ll.Size() != 0 {
		t.Error("Expected size = 0 but got", ll.Size())
	}

	if item != 1 {
		t.Error("Expected item value = 1 but got", item)
	}
}

func TestPushBack(t *testing.T) {
	ll := linkedlist.New()

	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)

	if ll.Size() != 3 {
		t.Error("Expected list size = 3 but got", ll.Size())
	}

	item, _ := ll.Get(0)

	if item != 1 {
		t.Error("Expected item value at 0 = 1 but got", item)
	}

	item, _ = ll.Get(1)

	if item != 2 {
		t.Error("Expected item value at 0 = 2 but got", item)
	}

	item, _ = ll.Get(2)

	if item != 3 {
		t.Error("Expected item value at 0 = 3 but got", item)
	}

	_, err := ll.Get(3)

	if err == nil {
		t.Error("Expected ErrIndexOutOfRange error")
	}
}

func TestPopBack(t *testing.T) {
	ll := linkedlist.New()

	item := ll.PopBack()

	if item != nil {
		t.Error("Expected item value = nil", item)
	}

	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)

	item = ll.PopBack()

	if ll.Size() != 2 {
		t.Error("Expected size = 2 but got", ll.Size())
	}

	if item != 3 {
		t.Error("Expected item value = 3 but got", item)
	}

	item = ll.PopBack()

	if ll.Size() != 1 {
		t.Error("Expected size = 1 but got", ll.Size())
	}

	if item != 2 {
		t.Error("Expected item value = 2 but got", item)
	}

	item = ll.PopBack()

	if ll.Size() != 0 {
		t.Error("Expected size = 0 but got", ll.Size())
	}

	if item != 1 {
		t.Error("Expected item value = 1 but got", item)
	}
}

func TestPushPop(t *testing.T) {
	ll := linkedlist.New()

	ll.PushFront(1)

	item := ll.PopBack()

	if item != 1 {
		t.Error("Expected item value = 1 but got", item)
	}

	ll.PushBack(2)

	item = ll.PopFront()

	if item != 2 {
		t.Error("Expected item value = 2 but got", item)
	}

	if !ll.Empty() {
		t.Error("Expected empty list")
	}
}

func TestFrontBack(t *testing.T) {
	ll := linkedlist.New()

	if ll.Front() != nil {
		t.Error("Expected nil")
	}

	if ll.Back() != nil {
		t.Error("Expected nil")
	}

	ll.PushBack(1)

	if ll.Front() != 1 {
		t.Error("Expected value = 1 but got", ll.Front())
	}

	if ll.Back() != 1 {
		t.Error("Expected value = 1 but got", ll.Back())
	}

	ll.PushBack(2)

	if ll.Front() != 1 {
		t.Error("Expected value = 1 but got", ll.Front())
	}

	if ll.Back() != 2 {
		t.Error("Expected value = 2 but got", ll.Back())
	}

	ll.PushFront(0)

	if ll.Front() != 0 {
		t.Error("Expected value = 0 but got", ll.Front())
	}

	if ll.Back() != 2 {
		t.Error("Expected value = 2 but got", ll.Back())
	}
}

func TestInsert(t *testing.T) {
	ll := linkedlist.New()

	err := ll.Insert(0, 1)

	if err == nil {
		t.Error("Expected ErrIndexOutOfRange error")
	}

	ll.PushFront(0)
	ll.Insert(0, 1)

	if item, _ := ll.Get(0); item != 1 {
		t.Error("Expected item value = 1 but got", item)
	}

	ll.Insert(1, 2)

	if item, _ := ll.Get(1); item != 2 {
		t.Error("Expected item value = 2 but got", item)
	}

	ll.Insert(1, 3)

	if ll.Size() != 4 {
		t.Error("Expected size = 4 but got", ll.Size())
	}

	if item, _ := ll.Get(0); item != 1 {
		t.Error("Expected item value = 1 but got", item)
	}

	if item, _ := ll.Get(1); item != 3 {
		t.Error("Expected item value = 3 but got", item)
	}

	if item, _ := ll.Get(2); item != 2 {
		t.Error("Expected item value = 2 but got", item)
	}

	if item := ll.Back(); item != 0 {
		t.Error("Expected item value = 0 but got", item)
	}
}

func TestErase(t *testing.T) {
	ll := linkedlist.New()

	if ll.Erase(0) {
		t.Error("Should not erase empty list")
	}

	ll.PushFront(1)

	if !ll.Erase(0) {
		t.Error("Should erase successfully")
	}

	if !ll.Empty() {
		t.Error("Expected empty list")
	}

	ll.PushFront(0)
	ll.PushBack(2)
	ll.Insert(1, 1)
	ll.PushBack(3)

	if !ll.Erase(3) {
		t.Error("Expected erase successfully")
	}

	if ll.Size() != 3 {
		t.Error("Expected size = 3 but got", ll.Size())
	}

	if item, _ := ll.Get(2); item != 2 {
		t.Error("Expected item value = 2 but got", item)
	}

	if !ll.Erase(1) {
		t.Error("Should erase successfully")
	}

	if ll.Size() != 2 {
		t.Error("Expected size = 2 but got", ll.Size())
	}

	if ll.Front() != 0 {
		t.Error("Expected item value = 0 but got", ll.Front())
	}

	if ll.Back() != 2 {
		t.Error("Expected item value = 2 but got", ll.Back())
	}
}

func TestValueNthFromEnd(t *testing.T) {
	ll := linkedlist.New()

	if item := ll.ValueNthFromEnd(0); item != nil {
		t.Error("Expected nil")
	}

	ll.PushBack(1)

	if item := ll.ValueNthFromEnd(0); item != 1 {
		t.Error("Expected 1 but got", item)
	}

	ll.PushBack(2)

	if item := ll.ValueNthFromEnd(0); item != 2 {
		t.Error("Expected 2 but got", item)
	}

	if item := ll.ValueNthFromEnd(1); item != 1 {
		t.Error("Expected 1 but got", item)
	}

	ll.PushBack(3)
	ll.PushBack(4)

	if item := ll.ValueNthFromEnd(0); item != 4 {
		t.Error("Expected 4 but got", item)
	}

	if item := ll.ValueNthFromEnd(1); item != 3 {
		t.Error("Expected 3 but got", item)
	}

	if item := ll.ValueNthFromEnd(2); item != 2 {
		t.Error("Expected 2 but got", item)
	}

	if item := ll.ValueNthFromEnd(3); item != 1 {
		t.Error("Expected 1 but got", item)
	}
}

func TestReverse(t *testing.T) {
	ll := linkedlist.New()

	ll.Reverse() // Reverse empty list

	ll.PushBack(1)
	ll.PushBack(2)

	ll.Reverse()

	if ll.Front() != 2 {
		t.Error("Expected value = 2 but got", ll.Front())
	}

	if ll.Back() != 1 {
		t.Error("Expected value = 1 but got", ll.Back())
	}

	ll.PushFront(3)
	ll.PushFront(4)

	ll.Reverse()

	if ll.Front() != 1 {
		t.Error("Expected value = 1 but got", ll.Front())
	}

	if ll.Back() != 4 {
		t.Error("Expected value = 4 but got", ll.Back())
	}

	if item, _ := ll.Get(1); item != 2 {
		t.Error("Expected value = 2 but got", item)
	}

	if item, _ := ll.Get(2); item != 3 {
		t.Error("Expected value = 3 but got", item)
	}
}

func TestRemoveValue(t *testing.T) {
	ll := linkedlist.New()

	if ll.RemoveValue(1) {
		t.Error("Should not remove on empty list")
	}

	ll.PushBack(1)

	if !ll.RemoveValue(1) {
		t.Error("Should remove successfully")
	}

	if !ll.Empty() {
		t.Error("Expected empty list")
	}

	ll.PushBack(2)
	ll.PushFront(1)

	if !ll.RemoveValue(2) {
		t.Error("Should remove successfully")
	}

	if ll.Size() != 1 {
		t.Error("Expected size = 1 but got", ll.Size())
	}

	if ll.Front() != ll.Back() {
		t.Error("Expected", ll.Front(), "=", ll.Back())
	}

	ll.PushBack(2)
	ll.PushBack(3)
	ll.PushBack(3)

	if !ll.RemoveValue(3) {
		t.Error("Should remove successfully")
	}

	if ll.Size() != 3 {
		t.Error("Expected size = 3 but got", ll.Size())
	}

	if ll.Front() != 1 {
		t.Error("Expected value = 1 but got", ll.Front())
	}

	if ll.Back() != 3 {
		t.Error("Expected value = 3 but got", ll.Back())
	}

	if item, _ := ll.Get(1); item != 2 {
		t.Error("Expected  value = 2 but got", item)
	}
}
