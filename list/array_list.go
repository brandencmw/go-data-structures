package list

type ArrayList[T comparable] []T

func (l ArrayList[T]) Append(item T) int {
	l = append(l, item)
	return len(l)
}

func (l ArrayList[T]) Prepend(item T) int {
	l = append([]T{item}, l...)
	return len(l)
}

func (l ArrayList[T]) Insert(item T, index int) int {
	if index < 0 || index > len(l) {
		panic("Invalid index for insert")
	} else if index == 0 {
		return l.Prepend(item)
	} else if index == len(l) {
		return l.Append(item)
	}

	front := append(l[:index], item)
	l = append(front, l[index:]...)

	return len(l)
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

func (l ArrayList[T]) ContentsEqualTo(listToCompare ArrayList[T]) bool {
	if len(l) != len(listToCompare) {
		return false
	}

	for i := 0; i < len(l); i++ {
		if l[i] != listToCompare[i] {
			return false
		}
	}
	return true
}
