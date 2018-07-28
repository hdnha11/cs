package linkedlistqueue_test

import (
	"testing"

	"github.com/hdnha11/cs/queues/linkedlistqueue"
)

func TestNew(t *testing.T) {
	llq := linkedlistqueue.New()

	if llq == nil {
		t.Error("Should return an not nil queue")
	}

	if !llq.Empty() {
		t.Error("Expected empty queue")
	}
}

func TestEnqueueDequeue(t *testing.T) {
	llq := linkedlistqueue.New()

	if item := llq.Dequeue(); item != nil {
		t.Error("Expected nil value")
	}

	llq.Enqueue(1)

	if llq.Empty() {
		t.Error("Expected the queue is not empty")
	}

	if item := llq.Dequeue(); item != 1 {
		t.Error("Expected item value = 1 but got", item)
	}

	if !llq.Empty() {
		t.Error("Expected an empty queue")
	}

	llq.Enqueue(2)
	llq.Enqueue(3)
	llq.Enqueue(4)

	if item := llq.Dequeue(); item != 2 {
		t.Error("Expected item value = 2 but got", item)
	}

	if llq.Empty() {
		t.Error("Expected the queue is not empty")
	}

	if item := llq.Dequeue(); item != 3 {
		t.Error("Expected item value = 3 but got", item)
	}

	if llq.Empty() {
		t.Error("Expected the queue is not empty")
	}

	if item := llq.Dequeue(); item != 4 {
		t.Error("Expected item value = 4 but got", item)
	}

	if !llq.Empty() {
		t.Error("Expected an empty queue")
	}
}
