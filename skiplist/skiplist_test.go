package skiplist

import (
	"testing"
)

func TestSkiplistInsert(t *testing.T) {
	l := newList()
	insert(l, 31, 5)
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

	bl := delete(l, 20)

	if bl {
		t.Log("key=20的元素已经被删除了")
	}

	ret = search(l, 20)

	t.Log(ret)
}
