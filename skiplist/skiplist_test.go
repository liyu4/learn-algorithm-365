package skiplist

import (
	"testing"
)

func TestSkiplistInsert(t *testing.T) {
	l := newList()
	insert(l, 31, 5)
	insert(l, 1, 3)
	insert(l, 31, 6)
	insert(l, 20, 7)
	insert(l, 20, 2)

	for _, v := range l.header.forward {
		if v != nil {
			t.Log(v.key)
			t.Log(v.value)
		}
	}

	ret := search(l, 20)

	if ret != -1 {
		t.Log(ret)
	} else {
		t.Failed()
	}
}
