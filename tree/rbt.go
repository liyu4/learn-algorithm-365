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

// 插入或者删除之后要保持二叉树的性质，也就是左边<=根节点  右边>=根结点
func (r *Rbtree) left_rotate(node *RBNode) {
	var temp *RBNode

	// 这一步的意义是将右边子树的最小关键字变更为node的右子树
	// 其他保持不变，这是因为node的右子树关键字必须大于等于node关键字
	// 总结的说就是  temp的左孩子成为node的右边孩子
	temp = node.right
	node.right = temp.left

	//判断右边子树的左子树是否为哨兵也就是是否为空
	if temp.left != r.sentinel {
		// 不等于更新的好的节点的双亲节点更新为node
		// 如果是相等的情况，那么事情会变的简单一些，那就是一个哨兵占据右边位置
		temp.left.parent = node
	}

	temp.parent = node.parent

	// 以下的三种情况分别是
	// node本身为树的根节点那么temp则直接占据这个位置，不需要再去更新左右的情况
	// 如果node的双亲节点不为空，并且node在双亲的左边，那么temp就更新为node双亲节点的左子树
	// 如果node在双亲的右边子树，那么temp更新为node双亲的右边子树
	// 上述所有的过程都是为了更新temp在其双亲节点中所处的位置
	{
		if node.parent == r.sentinel {
			r.root = temp
		} else if node == node.parent.left {
			temp.parent.left = temp
		} else {
			node.parent.right = temp
		}

	}

	// temp的左子树更新为node
	temp.left = node
	// node的双亲指向为temp
	node.parent = temp
	// 至此完成了左旋转过程
}

func (r *Rbtree) right_rotate(node *RBNode) {

}
