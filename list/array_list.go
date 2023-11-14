package list

type ArrayList[T comparable] struct {
	contents []T
}

func NewArrayList[T comparable](contents ...T) *ArrayList[T] {
	return &ArrayList[T]{contents: contents}
}

func (l *ArrayList[T]) Size() int {
	return len(l.contents)
}

func (l *ArrayList[T]) Append(item T) int {
	l.contents = append(l.contents, item)
	return len(l.contents)
}

func (l *ArrayList[T]) Prepend(item T) int {
	l.contents = append([]T{item}, l.contents...)
	return len(l.contents)
}

func (l *ArrayList[T]) Insert(item T, index int) int {
	if index < 0 || index > len(l.contents) {
		panic("Invalid index for insert")
	} else if index == 0 {
		return l.Prepend(item)
	} else if index == len(l.contents) {
		return l.Append(item)
	}

	front := l.contents[:index]
	lContents := make([]T, len(front))
	copy(lContents, front)
	lContents = append(lContents, item)
	l.contents = append(lContents, l.contents[index:]...)

	return l.Size()
}

func (l *ArrayList[T]) Remove(index int) int {
	if len(l.contents) == 0 {
		panic("Can't remove item from empty list")
	} else if len(l.contents) == 1 {
		l.contents = []T{}
	}

	l.contents = append(l.contents[:index], l.contents[index+1:]...)

	return l.Size()
}

func (l ArrayList[T]) Get(index int) T {
	if index > l.Size()-1 {
		panic("Invalid index")
	}
	return l.contents[index]
}

func (l *ArrayList[T]) Set(index int, item T) {
	if index > l.Size()-1 {
		panic("Invalid index")
	}
	l.contents[index] = item
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
