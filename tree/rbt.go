// tree provides an implement of a red-black tree
package tree

import (
	"fmt"
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

func (r *Rbtree) Init() {
	var sentinel RBNode
	sentinel.color = 2
	r.root = &sentinel     // 根结点为一个黑色的空哨兵
	r.sentinel = &sentinel // 哨兵节点
}
