package list

import (
	"github.com/brandencmw/go-data-structures.git/utils"
)

type ArrayList[T comparable] struct {
	contents []T
}

func NewArrayList[T comparable](contents ...T) *ArrayList[T] {
	return &ArrayList[T]{contents: contents}
}

func (l *ArrayList[T]) Size() int {
	return len(l.contents)
}

func (l *ArrayList[T]) Append(item T) {
	l.contents = append(l.contents, item)
}

func (l *ArrayList[T]) Prepend(item T) {
	l.contents = append([]T{item}, l.contents...)
}

func (l *ArrayList[T]) Insert(item T, idx int) error {
	if idx < 0 || idx > len(l.contents) {
		return &InvalidIndexError{l.Size(), idx}
	} else if idx == 0 {
		l.Prepend(item)
		return nil
	} else if idx == len(l.contents) {
		l.Append(item)
		return nil
	}

	front := l.contents[:idx]
	lContents := make([]T, len(front))
	copy(lContents, front)
	lContents = append(lContents, item)
	l.contents = append(lContents, l.contents[idx:]...)

	return nil
}

func (l *ArrayList[T]) Remove(index int) error {
	if len(l.contents) == 0 {
		return ErrEmptyList
	} else if len(l.contents) == 1 {
		l.contents = []T{}
	}

	l.contents = append(l.contents[:index], l.contents[index+1:]...)

	return nil
}

func (l ArrayList[T]) Get(idx int) (T, error) {
	if idx > l.Size()-1 {
		return utils.GetZero[T](), &InvalidIndexError{max: l.Size(), idx: idx}
	}
	return l.contents[idx], nil
}

func (l *ArrayList[T]) Set(idx int, item T) error {
	if idx > l.Size()-1 {
		return &InvalidIndexError{max: l.Size(), idx: idx}
	}
	l.contents[idx] = item
	return nil
}

func (l ArrayList[T]) Equals(listToCompare ArrayList[T]) bool {
	if l.Size() != listToCompare.Size() {
		return false
	}

	for i := 0; i < l.Size(); i++ {
		if l.contents[i] != listToCompare.contents[i] {
			return false
		}
	}
	return true
}

func (l ArrayList[T]) Clone() *ArrayList[T] {
	newListContents := make([]T, l.Size())
	for i := 0; i < l.Size(); i++ {
		newListContents[i] = l.contents[i]
	}
	return &ArrayList[T]{
		contents: newListContents,
	}
}

func (l ArrayList[T]) Contents() []T {
	returnContents := make([]T, l.Size())
	copy(returnContents, l.contents)
	return returnContents
}
