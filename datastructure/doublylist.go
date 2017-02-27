package datastructure

import (
	"errors"
	"fmt"
)

type Value struct {
	Name string
}

type DoubleNode struct {
	Value
	prev *DoubleNode
	next *DoubleNode
}

type doublelist struct {
	head *DoubleNode
	tail *DoubleNode
}

func (n *DoubleNode) Prev() *DoubleNode {
	return n.prev
}

func (n *DoubleNode) Next() *DoubleNode {
	return n.next
}

func (d *doublelist) First() *DoubleNode {
	return d.head
}

func NewDlist() *doublelist {
	return &doublelist{}
}

func (d *doublelist) Insert(v Value) {
	node := &DoubleNode{Value: v}

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

func (d *doublelist) Search(name string) *DoubleNode {
	var (
		found bool        = false
		ret   *DoubleNode = nil
	)

	for n := d.First(); n != nil && !found; n = n.next {
		if n.Value.Name == name {
			found = true
			ret = n
		}
	}
	return ret
}

func (d *doublelist) Delete(name string) bool {
	var (
		success bool = false
		nodedel      = d.Search(name)
	)
	fmt.Println(nodedel)

	if nodedel != nil {
		prev := nodedel.prev
		next := nodedel.next

		// 此种情况是判断删除最后一个元素
		if next == nil {
			prev.next = next
			success = true
		} else {
			prev.next, next.prev = next, prev
			success = true
		}
	}
	return success
}

// 移除最后一个元素。
func (d *doublelist) Remove() (v Value, err error) {
	if d.tail == nil {
		err = errors.New("Double list is empty")
	} else {
		v = d.tail.Value
		d.tail = d.tail.prev
		if d.tail == nil {
			d.head = nil
		}
	}
	return v, err
}
