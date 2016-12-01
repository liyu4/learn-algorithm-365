package hash

import (
	"fmt"
)

const (
	HASHSIZE = 101
)

type Node struct {
	key   string
	value string
	next  *Node
}

type Hashtable struct {
	header []*Node
}

func New() *Hashtable {

	h := &Hashtable{}
	h.header = make([]*Node, HASHSIZE)

	for i := 0; i < HASHSIZE; i++ {
		h.header[i] = nil
	}
	return h
}

// Hash is hash function
func (h *Hashtable) hash(key string) uint {
	var hk uint

	ascii := []byte(key)

	for _, v := range ascii {
		hk = uint(v) + hk*31
	}
	return hk % HASHSIZE
}

func (h *Hashtable) lookup(key string) *Node {
	hk := h.hash(key)

	np := h.header[hk]

	for ; np != nil; np = np.next {
		//冲突指的是key不同但是hash到了相同的slot上
		if np.key == key {
			return np
		}
	}
	return nil
}

func (h *Hashtable) Get(key string) string {
	np := h.lookup(key)

	if np != nil {
		return np.value
	}
	return ""
}

func (h *Hashtable) Put(key, value string) int {
	var (
		hk uint
		np *Node
	)
	np = h.lookup(key)

	//如果这个key没有被插入，那么执行插入操作,当然冲突的元素也会扔进来。
	if np == nil {
		hk = h.hash(key)
		// 实例化np
		np = &Node{}
		np.key = key

		// 先前的元素往后移动了
		np.next = h.header[hk]
		h.header[hk] = np
	}
	np.value = value
	return 1
}

// 下面是一个漂亮的但是无用的调试函数
func (h *Hashtable) display() {
	var i int
	var t *Node
	for i = 0; i < HASHSIZE; i++ {
		t = h.header[i]
		if t == nil {
			continue
		} else {
			for ; t != nil; t = t.next {
				fmt.Printf("key is: %s,value is: %s", t.key, t.value)
				fmt.Println()
			}
		}
	}
}
