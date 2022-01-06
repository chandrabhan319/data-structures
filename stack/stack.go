package stack

import "fmt"

type (
	// Stack representation.
	Stack struct {
		top *node
		len int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// New creates a new stack.
func New() *Stack {
	return &Stack{nil, 0}
}

// Len returns the number of items in the stack.
func (s *Stack) Len() int {
	return s.len
}

// Push a value onto the top of the stack.
func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.len++
}

// Peek views the top item on the stack.
func (s *Stack) Peek() (interface{}, error) {
	if s.len <= 0 || s.top == nil {
		return nil, fmt.Errorf("empty stack")
	}

	return s.top.value, nil
}

// Pop return the top item of the stack.
func (s *Stack) Pop() (interface{}, error) {
	if s.len <= 0 || s.top == nil {
		return nil, fmt.Errorf("empty stack")
	}

	n := s.top
	s.top = n.prev
	s.len--
	return n.value, nil
}

func (s *Stack) Parse() {
	n := s.top
	for i := 0; i < s.len; i++ {
		fmt.Println(n.value)
		n = n.prev
	}
}
