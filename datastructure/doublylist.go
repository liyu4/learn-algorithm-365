package datastructure

import (
	// "errors"
	"fmt"
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
		d.head = node
	} else {
		// 指针指向同一块区域
		d.tail.next = node
		// 新加入节点的指针指向上一个元素
		// d.tail == d.head  因为它们都指向同一个address
		node.prev = d.tail
	}
	//替换掉tail的值, 新tail就为node的值，d.tail.prev== head
	// d.tail.next == nil
	d.tail = node
}
