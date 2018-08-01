package hashmap

// Key type
type Key interface {
	hashCode() uint
}

// IntKey integer key
type IntKey uint

// StringKey string key
type StringKey string

func (ik IntKey) hashCode() uint {
	return uint(ik)
}

func (sk StringKey) hashCode() uint {
	var h uint

	for _, c := range []rune(string(sk)) {
		h += uint(c)
	}

	return h
}
