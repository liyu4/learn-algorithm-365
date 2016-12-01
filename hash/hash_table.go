package hash

import ()

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

	//如果这个key没有被插入，那么执行插入操作
	if np == nil {
		hk = h.hash(key)
		// 实例化np
		np = &Node{}
		np.key = key
		np.next = h.header[hk]
		h.header[hk] = np
	}
	np.value = value
	return 1
}
