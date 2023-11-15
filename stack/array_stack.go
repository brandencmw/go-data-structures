package stack

import (
	"fmt"

	"github.com/brandencmw/go-data-structures.git/utils"
)

// Head of stack is contents[0]
// By not exporting contents from the package we ensure that you can only interact
// with the stack through its methods, abstracting the implementation
type ArrayStack[T comparable] struct {
	contents []T
}

// Creates array based stack. First value passed in will be head of stack
// Last value passed will be bottom of stack
func NewArrayStack[T comparable](contents ...T) *ArrayStack[T] {
	fmt.Printf("Contents are: %v\n", contents)
	return &ArrayStack[T]{contents: contents}
}

func (s *ArrayStack[T]) Clone() AbstractStack[T] {
	newContent := make([]T, len(s.contents))
	copy(newContent, s.contents)
	return &ArrayStack[T]{contents: newContent}
}

// Like pop but returns top value without removing it
func (s ArrayStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return utils.GetZero[T](), &EmptyStackError{method: "peek"}
	}

	val := s.contents[0]
	return val, nil
}

// Prepends value to stack contents to create new head
func (s *ArrayStack[T]) Push(val T) {
	s.contents = append([]T{val}, s.contents...)
}

// Removes top value and returns it
func (s *ArrayStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return utils.GetZero[T](), &EmptyStackError{method: "pop"}
	}
	val := s.contents[0]
	s.contents = s.contents[1:]
	return val, nil
}

func (s ArrayStack[T]) IsEmpty() bool {
	return len(s.contents) == 0
}
