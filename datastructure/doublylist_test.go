package datastructure

import (
	"testing"
)

func TestDoubleListI(t *testing.T) {
	dlist := NewDlist()

	dlist.Insert(Value{"light"})

	dlist.Insert(Value{"dylan"})

	n := dlist.Search("light")
	t.Log(n.Name)

	dlist.Insert(Value{"hi"})
	success := dlist.Delete("hi")

	n = dlist.Search("hi")
	t.Log(n)

	t.Log(success)

}
