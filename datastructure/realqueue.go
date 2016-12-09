package datastructure

import (
	"errors"
	"fmt"
)

type RealNode struct {
	value int
}

type RealQueue struct {
	nodes []*RealNode
	head  int
	tail  int
}

func NewRealQueue(size int) *RealQueue {
	return &RealQueue{nodes: make([]*RealNode, size)}
}

func (r *RealQueue) Enqueue(key *RealNode) error {
	if r.head == r.tail%len(r.nodes) && r.tail != 0 {
		return errors.New("overflow")
	}

	r.nodes[r.tail%len(r.nodes)] = key
	fmt.Println("before tail:", r.tail)

	// 如果tail指向的是slice之外的下标，那么tail就开始指向slice开始的地方
	if r.tail%(len(r.nodes)-1) == 0 && r.tail != 0 {
		r.tail = 0
	} else {
		r.tail = r.tail + 1
	}
	return nil
}

func (r *RealQueue) Dequeuen() (*RealNode, error) {
	if r.head == r.tail && r.nodes[r.head] == nil {
		return nil, errors.New("underflow")
	}

	node := r.nodes[r.head]
	//增加哨兵
	r.nodes[r.head] = nil
	if r.head == (len(r.nodes) - 1) {
		r.head = 0
	} else {
		r.head = r.head + 1
	}

	return node, nil
}
