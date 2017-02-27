package datastructure

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	if !stack.empty() {
		t.Fatal("stack should not be empty")
	}

	stack.push(1)
	stack.push(2000)

	top, err := stack.pop()
	if err != nil {
		t.Fatalf("Pop failed %s", err.Error())
	}

	if top != 2000 {
		t.Fatalf("Pop got %s instead of 2000", top)
	}
}
