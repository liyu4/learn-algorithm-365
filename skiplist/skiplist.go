package skiplist

import ()

// 每个数据节点的信息
type nodeStructure struct {
	key     int
	value   int
	forward []*nodeStructure
}

// 每一个层的信息
type listStructure struct {
	level  int
	header *nodeStructure
}

const (
	MAX_NUMBER_OF_LEVELS = 11
)

func newList() *listStructure {
	l := &listStructure{}

	l.level = 0
	l.header = newNodeOfLevel(MAX_NUMBER_OF_LEVELS)

	var i int
	for i = 0; i < MAX_NUMBER_OF_LEVELS; i++ {
		l.header.forward[i] = nil
	}
	return l
}

func newNodeOfLevel(level int) *nodeStructure {
	nodeArr := make([]*nodeStructure, level)
	return &nodeStructure{forward: nodeArr}
}

func insert(l *listStructure, key, value int) bool {
	var k int
	var update [MAX_NUMBER_OF_LEVELS]*nodeStructure

	var p, q *nodeStructure

	p = l.header
	k = l.level

	for ; k >= 0; k-- {
		q = p.forward[k]

		for q != nil && q.key < key {
			p = q
			q = p.forward[k]
		}
		update[k] = p
	}

}
