package hashmap

import "errors"

// HashMap implementation
type HashMap struct {
	arr  []*mapNode
	size int
}

type mapNode struct {
	k Key
	v interface{}
}

var deleted = &mapNode{}

// New create a new hashmap
func New(cap int) (*HashMap, error) {
	if cap <= 0 {
		return nil, errors.New("Capacity should be greater than 0")
	}

	return &HashMap{make([]*mapNode, cap), 0}, nil
}

// Hash calculate hash
func (hm HashMap) Hash(k Key) int {
	return int(k.hashCode()) % cap(hm.arr)
}

// Size get the table size
func (hm HashMap) Size() int {
	return hm.size
}

// Cap get the capacity
func (hm HashMap) Cap() int {
	return cap(hm.arr)
}

// Add if key already exists, update value
func (hm *HashMap) Add(k Key, v interface{}) {
	// Double table size if 50% full
	if cap(hm.arr) <= hm.size*2 {
		hm.resize(cap(hm.arr) * 2)
	}

	h := hm.Hash(k)

	// Find empty slot
	for hm.arr[h] != nil && hm.arr[h] != deleted && hm.arr[h].k != k {
		h = (h + 1) % cap(hm.arr)
	}

	if hm.arr[h] == nil || hm.arr[h] == deleted {
		hm.size++
	}

	node := &mapNode{k, v}
	hm.arr[h] = node
}

// Exists check if key exists
func (hm HashMap) Exists(k Key) bool {
	return hm.Get(k) != nil
}

// Get value with key
func (hm HashMap) Get(k Key) interface{} {
	h := hm.Hash(k)

	for hm.arr[h] != nil {
		if hm.arr[h].k == k {
			return hm.arr[h].v
		}
		h = (h + 1) % cap(hm.arr)
	}

	return nil
}

// Remove key
func (hm *HashMap) Remove(k Key) {
	h := hm.Hash(k)

	for hm.arr[h] != nil {
		if hm.arr[h].k == k {
			hm.arr[h] = deleted
			hm.size--
			break
		}
		h = (h + 1) % cap(hm.arr)
	}

	// Halves size of the array if it's 12.5% full or less
	if cap(hm.arr) >= hm.size*8 {
		hm.resize(cap(hm.arr) / 2)
	}
}

func (hm *HashMap) resize(cap int) {
	newHm, err := New(cap)

	if err != nil {
		return
	}

	for _, n := range hm.arr {
		if n != nil && n != deleted {
			newHm.Add(n.k, n.v)
		}
	}

	hm.arr = newHm.arr
	hm.size = newHm.size
}
