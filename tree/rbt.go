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
	var temp *RBNode

	temp = node.left
	// temp的右孩子变成node的左孩子
	node.left = temp.right

	// 更新temp的双亲节点
	temp.parent = node.parent

	{
		if node.parent == r.sentinel {
			r.root = temp
		} else if node.parent.left == node {
			node.parent.left = temp
		} else {
			node.parent.right = temp
		}
	}

	temp.right = node
	node.parent = temp
}

func (r *Rbtree) rb_insert(element interface{}) *RBNode {
	var (
		insertnode RBNode
		node       = &insertnode
	)

	node.data = element

	// 新增节点的左右子树都是哨兵节点
	node.left = r.sentinel
	node.right = r.sentinel
	// 新插入节点的颜色应该标记为红色，如果为黑色则不能维护红黑树的性质\
	node.color = RED

	// 插入第一个节点，也就是说这是一种最简单和明了的case
	if r.root == r.sentinel {
		node.parent = nil
		// 根节点一点为黑色
		node.color = BLACK
		r.root = node
		// 插入成功返回
		return node
	}

	curpos := tree.root
	for {
		if curpos.key > node.key {
			if curpos.left == tree.sentinel {
				curpos.left = node
				node.parent = curpos
				break
			} else {
				curpos = curpos.left
			}
		} else if curpos.key < node.key {
			if curpos.right == tree.sentinel {
				curpos.right = node
				node.parent = curpos
				break
			} else {
				curpos = curpos.right
			}
		} else {
			curpos.data = node.data
			return node
		}
	}

	var uncle *RBTreeNode
	for node != tree.root && node.color == 'r' && node.parent.color == 'r' {
		if node.parent == node.parent.parent.left {
			uncle = node.parent.parent.right
			if uncle.color == 'r' {
				node.parent.color = 'b'
				uncle.color = 'b'
				node.parent.parent.color = 'r'
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					tree.LeftRotate(node)
				}
				node.parent.color = 'b'
				node.parent.parent.color = 'r'
				tree.RightRotate(node.parent.parent)
			}
		} else {
			uncle = node.parent.parent.left
			if uncle.color == 'r' {
				node.parent.color = 'b'
				uncle.color = 'b'
				node.parent.parent.color = 'r'
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					tree.RightRotate(node)
				}
				node.parent.color = 'b'
				node.parent.parent.color = 'r'
				tree.LeftRotate(node.parent.parent)
			}
		}

	}

	tree.root.color = 'b'

	return node

}
