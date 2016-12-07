package datastructure

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue(2)

	queue.push(&Node{1})
	queue.push(&Node{2})
	queue.push(&Node{3})
}
