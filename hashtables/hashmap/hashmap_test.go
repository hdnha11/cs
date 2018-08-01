package hashmap_test

import (
	"testing"

	"github.com/hdnha11/cs/hashtables/hashmap"
)

func TestNew(t *testing.T) {
	hm, err := hashmap.New(0)

	if err == nil {
		t.Error("Expected error")
	}

	hm, _ = hashmap.New(1)

	if hm == nil {
		t.Error("Expected not nil map")
	}
}

func TestAdd(t *testing.T) {
	hm, _ := hashmap.New(3)

	hm.Add(hashmap.IntKey(1), 1)
	hm.Add(hashmap.IntKey(2), 2)
	hm.Add(hashmap.IntKey(3), 3)

	if !hm.Exists(hashmap.IntKey(1)) {
		t.Error("Expected key 1 existed")
	}

	if !hm.Exists(hashmap.IntKey(2)) {
		t.Error("Expected key 2 existed")
	}

	if !hm.Exists(hashmap.IntKey(3)) {
		t.Error("Expected key 3 existed")
	}

	if v := hm.Get(hashmap.IntKey(1)); v != 1 {
		t.Error("Expected value = 1 but got", v)
	}

	if v := hm.Get(hashmap.IntKey(2)); v != 2 {
		t.Error("Expected value = 2 but got", v)
	}

	if v := hm.Get(hashmap.IntKey(3)); v != 3 {
		t.Error("Expected value = 3 but got", v)
	}

	hm.Add(hashmap.IntKey(4), 4)

	if !hm.Exists(hashmap.IntKey(4)) {
		t.Error("Expected key 4 existed")
	}

	if v := hm.Get(hashmap.IntKey(4)); v != 4 {
		t.Error("Expected value = 4 but got", v)
	}
}

func TestRemove(t *testing.T) {
	hm, _ := hashmap.New(3)

	hm.Add(hashmap.IntKey(1), 1)
	hm.Add(hashmap.IntKey(2), 2)
	hm.Add(hashmap.IntKey(3), 3)

	hm.Remove(hashmap.IntKey(1))

	if hm.Exists(hashmap.IntKey(1)) {
		t.Error("Expected key 1 not existed")
	}

	// Check remove no existed key
	hm.Remove(hashmap.IntKey(1))

	hm.Remove(hashmap.IntKey(2))

	if hm.Exists(hashmap.IntKey(2)) {
		t.Error("Expected key 2 not existed")
	}

	hm.Remove(hashmap.IntKey(3))

	if hm.Exists(hashmap.IntKey(3)) {
		t.Error("Expected key 3 not existed")
	}
}

func TestResize(t *testing.T) {
	n := 128
	hm, _ := hashmap.New(n / 2)

	for i := 0; i < n; i++ {
		hm.Add(hashmap.IntKey(i), i)
	}

	if hm.Size() != n {
		t.Error("Expected size = ", n, " but got", hm.Size())
	}

	if hm.Cap() != n*2 {
		t.Error("Expected capacity = ", n*2, " but got", hm.Cap())
	}

	for i := 0; i < n; i++ {
		if !hm.Exists(hashmap.IntKey(i)) {
			t.Error("Expected key =", i, "existed")
		}
	}

	for i := 0; i < n; i++ {
		hm.Remove(hashmap.IntKey(i))
	}

	if hm.Size() != 0 {
		t.Error("Expected size = 0 but got", hm.Size())
	}

	if hm.Cap() != 2 {
		t.Error("Expected capacity = 2 but got", hm.Cap())
	}

	for i := 0; i < n; i++ {
		if hm.Exists(hashmap.IntKey(i)) {
			t.Error("Expected key =", i, "not existed")
		}
	}
}

func TestMixedKeyTypes(t *testing.T) {
	hm, _ := hashmap.New(3)

	hm.Add(hashmap.IntKey(1), 1)
	hm.Add(hashmap.StringKey("two"), 2)
	hm.Add(hashmap.StringKey("3"), 3)

	if !hm.Exists(hashmap.IntKey(1)) {
		t.Error("Expected key 1 existed")
	}

	if !hm.Exists(hashmap.StringKey("two")) {
		t.Error("Expected key \"two\" existed")
	}

	if !hm.Exists(hashmap.StringKey("3")) {
		t.Error("Expected key \"3\" existed")
	}

	if v := hm.Get(hashmap.IntKey(1)); v != 1 {
		t.Error("Expected value = 1 but got", v)
	}

	if v := hm.Get(hashmap.StringKey("two")); v != 2 {
		t.Error("Expected value = 2 but got", v)
	}

	if v := hm.Get(hashmap.StringKey("3")); v != 3 {
		t.Error("Expected value = 3 but got", v)
	}
}
