package arraylist_test

import (
	"testing"

	"github.com/hdnha11/cs/arrays/arraylist"
)

func TestNewArrayList(t *testing.T) {
	al, err := arraylist.New(0)

	if err != nil {
		t.Error(err)
	}

	if al.Len() != 0 {
		t.Error("Length should be zero but got", al.Len())
	}

	if al.Cap() != 0 {
		t.Error("Capacity should be zero but got", al.Cap())
	}
}

func TestGet(t *testing.T) {
	al, _ := arraylist.New(0)

	_, err := al.Get(0)

	if err == nil {
		t.Error("Expected ErrIndexOutOfRange")
	}

	al, _ = arraylist.New(1)

	el, _ := al.Get(0)

	if el != nil {
		t.Error("Expected value = nil but got", el)
	}
}

func TestSet(t *testing.T) {
	al, err := arraylist.New(2)

	if err != nil {
		t.Error(err)
	}

	if al.Len() != 2 {
		t.Error("Expected length = 2 but got", al.Len())
	}

	if al.Cap() != 2 {
		t.Error("Expected capacity = 2 but got", al.Cap())
	}

	al.Set(0, 1)
	al.Set(1, 2)

	el, err := al.Get(0)

	if err != nil {
		t.Error(err)
	}

	if el != 1 {
		t.Error("Expected value = 1 but got", el)
	}

	el, err = al.Get(1)

	if err != nil {
		t.Error(err)
	}

	if el != 2 {
		t.Error("Expected value = 2 but got", el)
	}

	_, err = al.Get(2)

	if err == nil {
		t.Error("Expected ErrIndexOutOfRange")
	}
}

func TestPush(t *testing.T) {
	al, _ := arraylist.New(0)

	al.Push(1)

	if al.Len() != 1 {
		t.Error("Expected length = 1 but got", al.Len())
	}

	if al.Cap() != 2 {
		t.Error("Expected capacity = 2 but got", al.Cap())
	}

	el, err := al.Get(0)

	if err != nil {
		t.Error(err)
	}

	if el != 1 {
		t.Error("Expected value = 1 but got", el)
	}

	al.Push(2)
	al.Push(3)

	if al.Len() != 3 {
		t.Error("Expected length = 3 but got", al.Len())
	}

	if al.Cap() != 4 {
		t.Error("Expected capacity = 4 but got", al.Cap())
	}

	el, err = al.Get(2)

	if err != nil {
		t.Error(err)
	}

	if el != 3 {
		t.Error("Expected value = 3 but got", al.Len())
	}
}

func TestPop(t *testing.T) {
	al, _ := arraylist.New(4)
	al.Set(0, 1)
	al.Set(1, 2)
	al.Set(2, 3)
	al.Set(3, 4)

	el := al.Pop()

	if el != 4 {
		t.Error("Expected value = 4 but got", el)
	}

	if al.Len() != 3 {
		t.Error("Expected length = 3 but got", al.Len())
	}

	if al.Cap() != 4 {
		t.Error("Expected capacity = 4 but got", al.Cap())
	}

	el = al.Pop()

	if el != 3 {
		t.Error("Expected value = 3 but got", el)
	}

	if al.Len() != 2 {
		t.Error("Expected length = 2 but got", al.Len())
	}

	if al.Cap() != 4 {
		t.Error("Expected capacity = 4 but got", al.Cap())
	}

	el = al.Pop()

	if el != 2 {
		t.Error("Expected value = 2 but got", el)
	}

	if al.Len() != 1 {
		t.Error("Expected length = 1 but got", al.Len())
	}

	if al.Cap() != 2 {
		t.Error("Expected capacity = 2 but got", al.Cap())
	}
}
