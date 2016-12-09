package datastructure

import (
	"testing"
)

func TestRealQueueOverflow(t *testing.T) {
	realqueue := NewRealQueue(3)

	if err := realqueue.Enqueue(&RealNode{1}); err != nil {
		t.Log(err)
	}

	if err := realqueue.Enqueue(&RealNode{2}); err != nil {
		t.Log(err)
	}

	if err := realqueue.Enqueue(&RealNode{3}); err != nil {
		t.Log(err)
	}

	if err := realqueue.Enqueue(&RealNode{4}); err != nil {
		t.Log(err)
	}
}

func TestRealQueueUnderflow(t *testing.T) {
	realqueue := NewRealQueue(3)

	if node, err := realqueue.Dequeuen(); err != nil {
		t.Log(err)
	} else {
		t.Log(node)
	}
}

func TestRealQueue(t *testing.T) {
	realqueue := NewRealQueue(3)

	if err := realqueue.Enqueue(&RealNode{1}); err != nil {
		t.Log(err)
	}
	if err := realqueue.Enqueue(&RealNode{2}); err != nil {
		t.Log(err)
	}

	if node, err := realqueue.Dequeuen(); err != nil {
		t.Log(err)
	} else {
		t.Log(node)
	}

	if err := realqueue.Enqueue(&RealNode{1}); err != nil {
		t.Log(err)
	}
	if err := realqueue.Enqueue(&RealNode{2}); err != nil {
		t.Log(err)
	}
	// 2 2 1
	// for _, v := range realqueue.nodes {
	// 	t.Log(*v)
	// }

	//output: 1
	if node, err := realqueue.Dequeuen(); err != nil {
		t.Log(err)
	} else {
		t.Log(node)
	}

	//optput: 2
	if node, err := realqueue.Dequeuen(); err != nil {
		t.Log(err)
	} else {
		t.Log(node)
	}

	//output: 2
	if node, err := realqueue.Dequeuen(); err != nil {
		t.Log(err)
	} else {
		t.Log(node)
	}

	//output: underflow
	if node, err := realqueue.Dequeuen(); err != nil {
		t.Log(err)
	} else {
		t.Log(node)
	}
}
