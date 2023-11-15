package list

import (
	"errors"
	"fmt"
)

type AbstractList[T comparable] interface {
	Insert(T, int) int
	Append(T) int
	Prepend(T) int
	Remove(int) T
	Get(int) T
	Set(int, T)
	Equals(AbstractList[T]) bool
	Clone() *AbstractList[T]
	Contents() []T
}

type InvalidIndexError struct {
	max int
	idx int
}

func (e *InvalidIndexError) Error() string {
	return fmt.Sprintf("Index %v out of bounds for list of size %v", e.idx, e.max)
}

var ErrEmptyList = errors.New("Cannot remove element from empty list")
