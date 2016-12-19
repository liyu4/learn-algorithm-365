// tree provides an implement of a red-black tree
package tree

import (
// "fmt"
)

type Color byte

const RED, BLACK Color = 1, 2

// red-black node
type RBNode struct {
	element interface{} // data
	left    *RBNode     // left child
	right   *RBNode     // right child
	parent  *RBNode     // parent
	color   Color       // color
}

type Rbtree struct {
	root     *RBNode // root node
	sentinel *RBNode // a sentry or guard
}

//  初始化树
func (r *Rbtree) Init() {
	var sentinel RBNode
	sentinel.color = BLACK
	r.root = &sentinel     // 根结点为一个黑色的空哨兵
	r.sentinel = &sentinel // 哨兵节点
}

func NewRbtree() *Rbtree {
	var sentinel RBNode
	sentinel.color = BLACK
	return &Rbtree{root: &sentinel, sentinel: &sentinel}
}

func (r *Rbtree) left_rotate(node *RBNode) {
	var temp *RBNode

	// 此处元素为node节点的右子树
	temp = node.right
	// 转换temp的左子树成为node的左子树
	node.right = temp.left

	if temp.left != r.sentinel {
		temp.left.parent = node
	}

	temp.parent = node.parent

	if node.parent == r.sentinel {
		r.root = temp
	} else if node == node.parent.left {
		temp.parent.left = temp
	} else {
		node.parent.right = temp
	}

	temp.left = node
	node.parent = temp
}

func (r *Rbtree) right_rotate(node *RBNode) {

}
