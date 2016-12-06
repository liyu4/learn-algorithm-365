package datastructure

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := New()

	if stack.empty() {
		t.Log("stack is empty")
	}
	stack.push(1)
	stack.push(2000)

	if top, err := stack.pop(); err != nil {
		t.Log(err)
	} else {
		t.Log(top)
	}

	t.Log(stack)
}
