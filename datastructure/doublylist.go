package datastructure

import (
	"errors"
)

type Value struct {
	Name string
}

type Node struct {
	Value
	prev *Node
	next *Node
}

type doublelist struct {
	head *Node
	tail *Node
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Next() *Node {
	return n.next
}

func (d *doublelist) First() *Node {
	return d.head
}

func NewDlist() *doublelist {
	return &doublelist{}
}

func (d *doublelist) Insert(v Value) {
	node := &Node{Value: v}

	if d.head == nil {
		// 第一个节点数据
		d.head == node
	} else {
		d.tail.next = node
		node.prev = d.tail
	}
	d.tail = node
}
