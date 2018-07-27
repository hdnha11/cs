package arraylist

import (
	"errors"
)

// ArrayList data structure
type ArrayList struct {
	arr []interface{}
	len int
	cap int
}

// ErrIndexOutOfRange Index out of range error
var ErrIndexOutOfRange = errors.New("Index out of range")

// New init list with length
func New(length int) (*ArrayList, error) {
	if length < 0 {
		return nil, ErrIndexOutOfRange
	}

	al := new(ArrayList)
	al.arr = make([]interface{}, length)
	al.len = length
	al.cap = length

	return al, nil
}

// Len get the length
func (al ArrayList) Len() int {
	return al.len
}

// Cap get the capacity
func (al ArrayList) Cap() int {
	return al.cap
}

// Get get element at index
func (al ArrayList) Get(i int) (interface{}, error) {
	if i < 0 || i >= al.len {
		return nil, ErrIndexOutOfRange
	}

	return al.arr[i], nil
}

// Set set element at index
func (al ArrayList) Set(i int, el interface{}) error {
	if i < 0 || i >= al.len {
		return ErrIndexOutOfRange
	}

	al.arr[i] = el

	return nil
}

// Push push element to the end
func (al *ArrayList) Push(el interface{}) error {
	if al.len >= al.cap {
		cap := al.cap * 2
		if cap == 0 {
			cap = 2
		}
		al.resize(cap)
	}

	al.arr[al.len] = el
	al.len++

	return nil
}

// Pop remove element from the end and return it
func (al *ArrayList) Pop() interface{} {
	if al.len == 0 {
		return nil
	}

	el := al.arr[al.len-1]
	al.len--

	if al.len < al.cap/2 {
		cap := al.cap / 2
		al.resize(cap)
	}

	return el
}

func (al *ArrayList) resize(cap int) {
	arr := make([]interface{}, cap)
	copy(arr, al.arr)
	al.arr = arr
	al.cap = cap
}
