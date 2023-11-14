package list

type listNode[T comparable] struct {
	value T
	next  *listNode[T]
}

type LinkedList[T comparable] struct {
	head *listNode[T]
	Size int
}

func NewLinkedList[T comparable](contents ...T) *LinkedList[T] {
	size := len(contents)
	if size == 0 {
		return &LinkedList[T]{head: nil, Size: size}
	}

	head := newListNode[T](contents[0])
	if size == 1 {
		return &LinkedList[T]{head: head, Size: size}
	}

	curr := head
	for _, item := range contents[1:] {
		curr.next = newListNode[T](item)
		curr = curr.next
	}
	return &LinkedList[T]{head: head, Size: size}
}

func newListNode[T comparable](val T) *listNode[T] {
	return &listNode[T]{
		value: val,
		next:  nil,
	}
}

func (l LinkedList[T]) Head() T {
	return l.head.value
}

func (l *LinkedList[T]) Insert(index int, item T) int {
	if index < 0 || index > l.Size {
		panic("Invalid index")
	} else if index == 0 {
		return l.Prepend(item)
	} else if index == l.Size {
		return l.Append(item)
	}

	current := l.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	newNode := listNode[T]{
		value: item,
		next:  current.next,
	}
	current.next = &newNode
	l.Size++

	return l.Size
}

func (l *LinkedList[T]) Append(item T) int {

	newNode := newListNode[T](item)

	if l.Size == 0 {
		l.head = newNode
	} else {
		current := l.head
		for i := 0; i < l.Size-1; i++ {
			current = current.next
		}
		current.next = newNode
	}

	l.Size++
	return l.Size
}

func (l *LinkedList[T]) Prepend(item T) int {
	var newNodeNext *listNode[T]

	if l.Size == 0 {
		newNodeNext = nil
	} else {
		newNodeNext = l.head
	}

	l.head = &listNode[T]{
		value: item,
		next:  newNodeNext,
	}
	l.Size++

	return l.Size
}

func (l *LinkedList[T]) Remove(index int) T {
	if index < 0 || index > l.Size-1 {
		panic("Invalid index")
	}

	current := l.head
	var previous *listNode[T]
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
	l.Size--

	return returnVal
}

func (l LinkedList[T]) Get(index int) T {
	if index < 0 || index > l.Size-1 {
		panic("Invalid index")
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value
}

func (l *LinkedList[T]) Set(index int, item T) {
	if index < 0 || index > l.Size-1 {
		panic("Invalid index")
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.value = item
}

func (l LinkedList[T]) Equals(listToCompare LinkedList[T]) bool {
	if l.Size != listToCompare.Size {
		return false
	}

	lCurr := l.head
	cCurr := listToCompare.head
	for lCurr != nil && cCurr != nil {
		if lCurr.value != cCurr.value {
			return false
		}
		lCurr = lCurr.next
		cCurr = cCurr.next
	}
	if lCurr != nil || cCurr != nil {
		return false
	}
	return true
}

func (l LinkedList[T]) Clone() *LinkedList[T] {
	newList := &LinkedList[T]{}
	if l.head == nil {
		return newList
	}

	newList.head = &listNode[T]{value: l.head.value, next: nil}
	newList.Size = 1

	lCurr := l.head.next
	newCurr := newList.head
	for lCurr != nil {
		newCurr.next = newListNode[T](lCurr.value)
		newList.Size++
		newCurr = newCurr.next
		lCurr = lCurr.next
	}
	return newList
}

func (l *LinkedList[T]) Contents() []T {
	contents := []T{}
	curr := l.head
	for curr != nil {
		contents = append(contents, curr.value)
		curr = curr.next
	}
	return contents
}
