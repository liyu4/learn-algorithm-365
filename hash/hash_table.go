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
