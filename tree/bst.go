// tree provides an implement of binary search trees
package tree

import (
	"fmt"
)

// 树中每一个节点包含的属性
type Node struct {
	// 可以存储任意类型的数据
	element interface{}
	// 双亲
	parent *Node
	// 左孩子
	left *Node
	// 右孩子
	right *Node
}

//比较器,判断是走左子树还是右子树
type comparer func(interface{}, interface{}) bool

type Bst struct {
	compare comparer
	root    *Node
}

// 创建一个空的二叉树
func (b *Bst) New(compare comparer) *Bst {
	return &Bst{compare: compare}
}

// 中序遍历
// 子树根的关键字位于左子树关键字和右子树关键字的中间
func inorder_tree_walk(tree *Node) {
	if tree == nil {
		return
	}

	inorder_tree_walk(tree.left)
	fmt.Println(tree.element)
	inorder_tree_walk(tree.right)
}

// 查找一个具有给定关键字的节点,运行时间最坏情况为这颗树的高h, O(h)
func (b *Bst) tree_search(tree *Node, element interface{}) *Node {
	if tree == nil {
		return nil
	}

	if tree.element == element {
		return tree
	}

	// 为真则走左，为假则走右
	if b.compare(element, tree.element) {
		return b.tree_search(tree.right, element)
	} else {
		return b.tree_search(tree.left, element)
	}
}

// 迭代法，避免递归的性能消耗， 你懂的
func (b *Bst) iterative_tree_search(tree *Node, element interface{}) *Node {
	if tree == nil {
		return nil
	}

	for element != tree.element {
		if b.compare(element, tree.element) {
			tree = tree.right
		} else {
			tree = tree.left
		}
	}
	return tree
}

// 最小关键元素
func (b *Bst) tree_minimum(tree *Node) *Node {
	for tree != nil {
		tree = tree.left
	}
	return tree
}

// 最大关键元素, 显然最大关键元素和最小关键元素都是由二叉树的性质决定的，并且它们的代码也是对称的。
func (b *Bst) tree_maximum(tree *Node) *Node {
	for tree != nil {
		tree = tree.right
	}
	return tree
}

// 后继
func (b *Bst) tree_successor(tree *Node) *Node {
	if tree.right != nil {
		return b.tree_minimum(tree.right)
	}

	y := tree.parent

	for y != nil && tree == y.right {
		tree = y
		y = y.parent
	}
	return y
}

// 前继
func (b *Bst) tree_predecessor(tree *Node) *Node {
	if tree.left != nil {
		return b.tree_maximum(tree.left)
	}

	y := tree.parent

	for y != nil && tree == y.right {
		tree = y
		y = y.parent
	}
	return y
}

// 插入
func (b *Bst) tree_insert(tree *Node, element interface{}) *Node {
	y := &Node{}
	for tree != nil {
		// y为双亲节点
		y = tree
		// 举例子n1 < n2 为真则走左， 为假则走右
		if b.compare(element, tree.element) {
			tree = tree.left
		} else {
			tree = tree.right
		}
	}

	var node *Node = &Node{element, nil, nil, nil}
	node.parent = y

	// 说明是一颗空二叉搜索树
	if y == nil {
		b.root = node
		return node
	}

	// 寻找左边还是右边需要插入
	if b.compare(element, y.element) {
		y.left = node
	} else {
		y.right = node
	}
	return y
}

// 转移, 用一棵子树替换一棵子树并成为其双亲的孩子结点。
func (b *Bst) transplant(to, from *Node) {
	// 如果to 没有双亲则代表to就是这棵树的根
	if to.parent == nil {
		b.root = from
		// to是否是双亲的左子树
	} else if to == to.parent.left {
		// 否则to的双亲节点的左子树就用from代替
		to.parent.left = from
	} else {
		// 否则to的双亲节点的右子树就用from代替
		to.parent.right = from
	}

	// 允许from为空
	if from != nil {
		from.parent = to.parent
	}
}

// 删除
func (b *Bst) tree_delete(tree *Node, element interface{}) *Node {
	// 寻找当前节点的位置
	target := b.tree_search(tree, element)

	// 如果target的左子树为空
	if target.left == nil {
		b.transplant(target, target.right)
	} else if target.right == nil {
		b.transplant(target, target.left)
	} else {
		// target的后继元素，同时这里是处理target的左右子树都非空，所以情况稍微复杂一点
		// 啰嗦一点y就是后继元素
		y := b.tree_minimum(target.right)
		if y.parent != target {
			// 更新后继节点右子树的位置
			b.transplant(y, y.right)
			// 更新后继节点的位置
			y.right = target.right
			y.right.parent = y
		}

		// 如果y是target的右孩子，那么用y替换target，并且留下y的右孩子
		// 更新这棵树的左右情况
		b.transplant(target, y)
		y.left = target.left
		y.left.parent = y
	}
	return tree
}
