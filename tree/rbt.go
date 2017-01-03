// tree provides an implement of a red-black tree
package tree

import (
// "fmt"
)

type Color byte

const RED, BLACK Color = 1, 2

type comparer2 func(interface{}, interface{}) bool

// red-black node
type RBNode struct {
	key     int64
	element interface{} // data
	left    *RBNode     // left child
	right   *RBNode     // right child
	parent  *RBNode     // parent
	color   Color       // color
}

type Rbtree struct {
	compare  comparer2
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

func initNode(key int64, element interface{}) *RBNode {

	// 新插入节点的颜色应该标记为红色，如果为黑色则不能维护红黑树的性质
	return &RBNode{key: key, element: element, color: RED}
}

func (r *Rbtree) rb_insert(key int64, element interface{}) *RBNode {

	node := initNode(key, element)

	// 新增节点的左右子树都是哨兵节点
	node.left = r.sentinel
	node.right = r.sentinel

	// 插入第一个节点，也就是说这是一种最简单和明了的case
	if r.root == r.sentinel {
		node.parent = nil
		// 根节点一点为黑色
		node.color = BLACK
		r.root = node
		// 插入成功返回
		return node
	}

	// 当前位置,从root开始遍历，可以在O(logn)时间内完成这个过程
	curpos := r.root
	for {
		// 举例子n1 < n2 为真则走左， 为假则走右
		if r.compare(curpos.key, key) {
			// 当前的左节点为哨兵
			if curpos.left == r.sentinel {
				// 左边节点赋值为node
				curpos.left = node
				// 更新的node的双亲节点
				node.parent = curpos
				// 结束插入
				break
			} else {
				// 如果左边节点不为空，则往下更新当前节点
				curpos = curpos.left
			}
		} else {
			// 说明上面的问题，那就是往右子树方向寻找需要插入的地方
			if curpos.right == r.sentinel {
				// 如果有边节点为空
				// 这更新当前位置节点的右子树为node
				curpos.right = node
				// 更新node的双亲节点为当前位置节点
				node.parent = curpos
				// 结束
				break
			} else {
				// 从右边往下更新当前位置节点
				curpos = curpos.right
			}
		}

		// 这一步的意思就是更新某一个节点，当然如果你清楚的知道你要更新的节点，这一段可以抽出去

		if key == curpos.key {
			// 已存在的节点更新自身的element，然后返回不需要维护红黑树的性质
			curpos.element = node.element
			return node
		}
	}

	r.rb_insert_fixup(node)
	r.root.color = BLACK
	return node
}

func (r *Rbtree) rb_insert_fixup(node *RBNode) {
	var uncle *RBNode
	for node.parent.color == RED {
		if node.parent == node.parent.parent.left {
			uncle = node.parent.parent.right
			if uncle.color == RED {
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else if node == node.parent.right {
				node = node.parent
				r.left_rotate(node)
			}
			node.parent.color = BLACK
			node.parent.parent.color = RED
			r.right_rotate(node)
		} else {
			uncle = node.parent.parent.left
			if uncle.color == RED {
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else if node == node.parent.left {
				node = node.parent
				r.right_rotate(node)
			}
			node.parent.color = BLACK
			node.parent.parent.color = RED
			r.left_rotate(node)
		}
	}
}

func (r *Rbtree) Delete(key int64) {
	// 寻找需要删除的节点
	// 然后开始维护这颗红黑树
	node := r.search(key)
	// 如果是空树则直接返回
	if node == nil {
		return
	}

	var temp *RBNode
	var subst *RBNode
	var w *RBNode
	if node.left == r.sentinel {
		temp = node.right
		subst = node
	} else if node.right == r.sentinel {
		temp = node.right
		subst = node
	} else {
		subst = r.min()
		if subst.left != r.sentinel {
			temp = subst.left
		} else {
			temp = subst.right
		}
	}

	if subst == r.root {
		r.root = temp
		temp.color = BLACK

		node.parent = nil
		node.left = nil
		node.right = nil

		return
	}

	color := subst.color

	if subst == subst.parent.left {
		subst.parent.left = temp
	} else {
		subst.parent.right = temp
	}

	if subst == node {
		temp.parent = subst.parent
	} else {
		if subst.parent == node {
			temp.parent = subst
		} else {
			temp.parent = subst.parent
		}

		subst.left = node.left
		subst.right = node.right
		subst.parent = node.parent
		subst.color = node.color

		if node == r.root {
			r.root = subst
		} else {
			if node == node.parent.left {
				node.parent.left = subst
			} else {
				node.parent.right = subst
			}
		}

		if subst.left != r.sentinel {
			subst.left.parent = subst
		}
		if subst.right != r.sentinel {
			subst.right.parent = subst
		}
	}

	node.parent = nil
	node.left = nil
	node.right = nil

	if color == RED {
		return
	}

	for temp != r.root && temp.color == BLACK {
		if temp == temp.parent.left {
			w = temp.parent.right
			if w.color == RED {
				w.color = BLACK
				temp.parent.color = RED
				r.left_rotate(temp.parent)
				w = temp.parent.right
			}

			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = BLACK
				temp = temp.parent
			} else {
				if w.right.color == BLACK {
					w.left.color = BLACK
					w.color = BLACK
					r.right_rotate(w)
					w = temp.parent.right
				}

				w.color = temp.parent.color
				temp.parent.color = BLACK
				w.right.color = BLACK
				r.left_rotate(temp.parent)
				temp = r.root
			}
		} else {
			w = temp.parent.left
			if w.color == RED {
				w.color = BLACK
				temp.parent.color = RED
				r.right_rotate(temp.parent)
				w = temp.parent.left
			}

			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				temp = temp.parent
			} else {
				if w.left.color == BLACK {
					w.right.color = BLACK
					w.color = 'r'
					r.left_rotate(w)
					w = temp.parent.left
				}

				w.color = temp.parent.color
				temp.parent.color = BLACK
				w.left.color = BLACK
				r.right_rotate(temp.parent)
				temp = r.root
			}
		}
	}

	temp.color = BLACK
}

func (r *Rbtree) search(key int64) *RBNode {
	curpos := r.root

	for {
		if curpos == nil {
			return nil
		}

		if curpos.key == key {
			return curpos
		}

		// 举例子n1 < n2 为真则走左， 为假则走右
		if r.compare(curpos.key, key) {
			curpos = curpos.left
		} else {
			curpos = curpos.right
		}
	}
}

func (r *Rbtree) min() *RBNode {
	curpos := r.root
	for {
		if curpos.left != r.sentinel {
			curpos = curpos.left
		} else {
			return curpos
		}
	}
}

func (r *Rbtree) max() *RBNode {
	curpos := r.root
	for {
		if curpos.right != r.sentinel {
			curpos = curpos.right
		} else {
			return curpos
		}
	}
}

func (r *Rbtree) IsEmpty() bool {
	if r.root == r.sentinel {
		return true
	} else {
		return false
	}
}
