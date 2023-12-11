package stack

import "fmt"

type Stack[T any] interface {
	Push(elem T)
	Pop() (T, error)
	Empty() bool
	Peek() (T, error)
}

type stack[T any] struct {
	elems []T
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{
		elems: []T{},
	}
}

// Empty checks if stack is empty
func (s *stack[T]) Empty() bool {
	return len(s.elems) == 0
}

// Peek returns the value of element on top of the stack without removing it
func (s *stack[T]) Peek() (T, error) {
	if s.Empty() {
		var t T
		return t, fmt.Errorf("empty stack")
	}
	return s.elems[len(s.elems)-1], nil
}

// Pop returns the value of element on top of the stack and removes it from
// stack
func (s *stack[T]) Pop() (T, error) {
	if s.Empty() {
		var t T
		return t, fmt.Errorf("empty stack")
	}

	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return elem, nil
}

// Push implements Stack.
func (s *stack[T]) Push(elem T) {
	s.elems = append(s.elems, elem)
}
