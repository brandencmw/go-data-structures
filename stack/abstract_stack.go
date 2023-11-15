package stack

import "fmt"

type AbstractStack[T comparable] interface {
	Push(T)
	Pop() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Clone() AbstractStack[T]
}

// Converts copy of stack contents into a slice
func AsSlice[T comparable](s AbstractStack[T]) []T {
	l := []T{}
	c := s.Clone() // Operate on clone to avoid modifying original object
	for !c.IsEmpty() {
		val, err := c.Pop()
		if err != nil {
			panic(err.Error())
		}
		l = append(l, val)
	}
	return l
}

func Equal[T comparable](s1, s2 AbstractStack[T]) bool {
	// Operate on clones to avoid modifying original objects
	c1 := s1.Clone()
	c2 := s2.Clone()
	for !c1.IsEmpty() && !c2.IsEmpty() {
		v1, err := c1.Pop()
		if err != nil {
			panic(err.Error())
		}
		v2, err := c2.Pop()
		if err != nil {
			panic(err.Error())
		}
		if v1 != v2 {
			return false
		}
	}

	return c1.IsEmpty() && c2.IsEmpty()
}

type EmptyStackError struct {
	method string
}

func (e *EmptyStackError) Error() string {
	return fmt.Sprintf("Cannot use %v on an empty stack", e.method)
}
