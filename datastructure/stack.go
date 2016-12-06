package datastructure

import (
	"errors"
)

//栈的数组实现
type stack struct {
	arr []int
}

func New() *stack {
	return &stack{}
}

func (s *stack) empty() bool {
	if len(s.arr) == 0 {
		return true
	}
	return false
}

func (s *stack) push(key int) {
	s.arr = append(s.arr, key)
}

func (s *stack) pop() (int, error) {
	var top int
	topsize := len(s.arr)
	if s.empty() {
		return -1, errors.New("underflow")
	} else {
		if topsize == 1 {
			top = s.arr[topsize-1]
			s.clean()
		} else {
			top = s.arr[topsize-1]
			s.arr = s.arr[:topsize-1]
		}
	}
	return top, nil
}

func (s *stack) clean() {
	s.arr = s.arr[:0:0]
}
