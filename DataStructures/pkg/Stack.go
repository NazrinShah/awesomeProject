package pkg

import "errors"

type Stack struct {
	top  *node
	size int
}

func (s *Stack) Count() int {
	return s.size
}

func (s *Stack) Peek() (int, error) {
	if s.size == 0 {
		return 0, errors.New("stack is empty")
	}

	return s.top.value, nil
}

func (s *Stack) Push(v int) {
	newNode := &node{
		value: v,
	}

	if s.top == nil {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}

	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return 0, errors.New("stack is empty")
	}

	v := s.top.value
	s.top = s.top.next
	s.size--

	if s.size == 0 {
		s.top = nil
	}

	return v, nil
}
