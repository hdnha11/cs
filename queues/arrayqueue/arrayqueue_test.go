package arrayqueue_test

import (
	"testing"

	"github.com/hdnha11/cs/queues/arrayqueue"
)

func TestNew(t *testing.T) {
	aq := arrayqueue.New(0)

	if aq == nil {
		t.Error("Should return an not nil queue")
	}

	if !aq.Empty() {
		t.Error("Expected empty queue")
	}
}

func TestEnqueueDequeue(t *testing.T) {
	aq := arrayqueue.New(4)

	if item := aq.Dequeue(); item != nil {
		t.Error("Expected nil value")
	}

	aq.Enqueue(1)

	if aq.Empty() {
		t.Error("Expected the queue is not empty")
	}

	if item := aq.Dequeue(); item != 1 {
		t.Error("Expected item value = 1 but got", item)
	}

	if !aq.Empty() {
		t.Error("Expected an empty queue")
	}

	aq.Enqueue(1)
	aq.Enqueue(2)
	aq.Enqueue(3)
	aq.Enqueue(4)
	err := aq.Enqueue(5)

	if err == nil {
		t.Error("Expected ErrNotEnoughSpace error")
	}

	if item := aq.Dequeue(); item != 1 {
		t.Error("Expected item value = 1 but got", item)
	}

	if aq.Empty() {
		t.Error("Expected the queue is not empty")
	}

	if item := aq.Dequeue(); item != 2 {
		t.Error("Expected item value = 2 but got", item)
	}

	if aq.Empty() {
		t.Error("Expected the queue is not empty")
	}

	if item := aq.Dequeue(); item != 3 {
		t.Error("Expected item value = 3 but got", item)
	}

	if aq.Empty() {
		t.Error("Expected the queue is not empty")
	}

	if item := aq.Dequeue(); item != 4 {
		t.Error("Expected item value = 4 but got", item)
	}

	if !aq.Empty() {
		t.Error("Expected an empty queue")
	}
}
