package list

type ListNode[T any] struct {
	value T
	next  *ListNode[T]
}

type LinkedList[T any] struct {
	head *ListNode[T]
	size int
}

func (l LinkedList[T]) Insert(item T, index int) int {
	if index < 0 || index > l.size {
		panic("Invalid index")
	} else if index == 0 {
		return l.Prepend(item)
	} else if index == l.size {
		return l.Append(item)
	}

	current := l.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	newNode := ListNode[T]{
		value: item,
		next:  current.next,
	}
	current.next = &newNode
	l.size++

	return l.size
}

func (l LinkedList[T]) Append(item T) int {

	newNode := ListNode[T]{
		value: item,
		next:  nil,
	}

	if l.size == 0 {
		l.head = &newNode
	} else {
		current := l.head
		for i := 0; i < l.size-1; i++ {
			current = current.next
		}
		current.next = &newNode
	}

	l.size++
	return l.size
}

func (l LinkedList[T]) Prepend(item T) int {
	var newNodeNext *ListNode[T]

	if l.size == 0 {
		newNodeNext = nil
	} else {
		newNodeNext = l.head
	}

	l.head = &ListNode[T]{
		value: item,
		next:  newNodeNext,
	}
	l.size++

	return l.size
}

func (l LinkedList[T]) Remove(index int) T {
	if index < 0 || index > l.size-1 {
		panic("Invalid index")
	}

	current := l.head
	var previous *ListNode[T]
	for i := 0; i < index; i++ {
		previous = current
		current = current.next
	}

	if current == l.head {
		l.head = l.head.next
	} else {
		previous.next = current.next
	}
	returnVal := current.value
	current = nil
	l.size--

	return returnVal
}

func (l LinkedList[T]) Get(index int) T {
	if index < 0 || index > l.size-1 {
		panic("Invalid index")
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value
}

func (l LinkedList[T]) Set(item T, index int) {
	if index < 0 || index > l.size-1 {
		panic("Invalid index")
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.value = item
}
