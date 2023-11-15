package stack

import (
	"fmt"

	"github.com/brandencmw/go-data-structures.git/utils"
)

type stackNode[T comparable] struct {
	value T
	next  *stackNode[T] // The element "below" in the stack
}

type LinkedStack[T comparable] struct {
	top  *stackNode[T]
	Size uint
}

// First argument becomes top of stack, last argument becomes bottom
func NewLinkedStack[T comparable](content ...T) *LinkedStack[T] {
	s := uint(len(content))
	if s == 0 {
		return &LinkedStack[T]{top: nil, Size: s}
	}

	t := &stackNode[T]{value: content[0], next: nil}
	curr := t
	for _, val := range content[1:] {
		curr.next = &stackNode[T]{value: val, next: nil}
		curr = curr.next
	}

	return &LinkedStack[T]{top: t, Size: s}
}

func (s *LinkedStack[T]) Clone() AbstractStack[T] {
	c := make([]T, 0)

	// Unload stack into slice to create new stack
	for !s.IsEmpty() {
		val, _ := s.Pop()
		c = append(c, val)
	}

	fmt.Printf("Clone contents: %v\n", c)
	r := NewLinkedStack[T](c...)

	// Load contents back into stack in same order
	// Start at back because c[0] should be top
	for i := len(c) - 1; i >= 0; i-- {
		s.Push(c[i])
	}
	return r
}

func (s *LinkedStack[T]) Push(val T) {
	t := &stackNode[T]{
		value: val,
		next:  s.top,
	}
	s.top = t
	s.Size++
}

func (s *LinkedStack[T]) Pop() (T, error) {
	if s.Size == 0 {
		return utils.GetZero[T](), &EmptyStackError{method: "pop"}
	}
	val := s.top.value
	s.top = s.top.next
	return val, nil
}

func (s LinkedStack[T]) Peek() (T, error) {
	if s.Size == 0 {
		return utils.GetZero[T](), &EmptyStackError{method: "peek"}
	}
	return s.top.value, nil
}

func (s LinkedStack[T]) IsEmpty() bool {
	return s.top == nil
}
