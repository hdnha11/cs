package linkedlist

type listNode struct {
	value interface{}
	next  *listNode
}

func (ln *listNode) findNth(i int) *listNode {
	if i == 0 {
		return ln
	}

	if i < 0 || ln.next == nil {
		return nil
	}

	return ln.next.findNth(i - 1)
}

func (ln *listNode) findTo(tn *listNode) *listNode {
	if ln.next == nil {
		return nil
	}

	if ln.next == tn {
		return ln
	}

	return ln.next.findTo(tn)
}

func (ln *listNode) findByValue(item interface{}) *listNode {
	if ln == nil {
		return nil
	}

	if ln.value == item {
		return ln
	}

	return ln.next.findByValue(item)
}
