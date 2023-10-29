package list

type ArrayList[T any] []T

func (l *ArrayList[T]) Append(item T) int {
	return 0
}

func (l *ArrayList[T]) Prepend(item T) int {
	return 0
}

func (l ArrayList[T]) Insert(item T, index int) int {
	return 0
}

func (l ArrayList[T]) Remove(index int) int {
	if len(l) == 0 {
		panic("Can't remove item from empty list")
	} else if len(l) == 1 {
		l = ArrayList[T]{}
	}

	l = append(l[:index], l[index+1:]...)

	return len(l)
}

func (l ArrayList[T]) Get(index int) T {
	if index > len(l)-1 {
		panic("Invalid index")
	}
	return l[index]
}

func (l ArrayList[T]) Set(item T, index int) {
	if index > len(l)-1 {
		panic("Invalid index")
	}
	l[index] = item
}
