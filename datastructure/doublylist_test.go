package datastructure

import (
	"testing"
)

func TestDoubleListI(t *testing.T) {
	dlist := NewDlist()

	dlist.Insert(Value{"light"})

	dlist.Insert(Value{"dylan"})

	t.Log(*dlist.head)
	t.Log(*dlist.tail)
}
