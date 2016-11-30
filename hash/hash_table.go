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

type hashtable struct {
	header []*Node
}

func New() *hashtable {

	h := &hashtable{}
	h.header = make([]*Node, HASHSIZE)

	for i := 0; i < HASHSIZE; i++ {
		h.header[i] = nil
	}
	return h
}

func hash(key string) {

}
