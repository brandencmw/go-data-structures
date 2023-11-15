package list

import "github.com/brandencmw/go-data-structures.git/utils"

type listNode[T comparable] struct {
	value T
	next  *listNode[T]
}

type LinkedList[T comparable] struct {
	head *listNode[T]
	Size int
}

func NewLinkedList[T comparable](contents ...T) *LinkedList[T] {
	s := len(contents)
	if s == 0 {
		return &LinkedList[T]{head: nil, Size: s}
	}

	h := newListNode[T](contents[0])
	curr := h
	for _, item := range contents[1:] {
		curr.next = newListNode[T](item)
		curr = curr.next
	}
	return &LinkedList[T]{head: h, Size: s}
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

func (l *LinkedList[T]) Insert(idx int, item T) error {
	if idx < 0 || idx > l.Size {
		return &InvalidIndexError{max: l.Size, idx: idx}
	} else if idx == 0 {
		l.Prepend(item)
		return nil
	} else if idx == l.Size {
		l.Append(item)
		return nil
	}

	curr := l.head
	for i := 0; i < idx-1; i++ {
		curr = curr.next
	}

	newNode := listNode[T]{
		value: item,
		next:  curr.next,
	}
	curr.next = &newNode
	l.Size++

	return nil
}

func (l *LinkedList[T]) Append(item T) {

	node := newListNode[T](item)

	if l.Size == 0 {
		l.head = node
	} else {
		curr := l.head
		for i := 0; i < l.Size-1; i++ {
			curr = curr.next
		}
		curr.next = node
	}
	l.Size++
}

func (l *LinkedList[T]) Prepend(item T) {
	var next *listNode[T]

	if l.Size == 0 {
		next = nil
	} else {
		next = l.head
	}

	l.head = &listNode[T]{
		value: item,
		next:  next,
	}
	l.Size++
}

func (l *LinkedList[T]) Remove(idx int) (T, error) {
	if idx < 0 || idx > l.Size-1 {
		return utils.GetZero[T](), &InvalidIndexError{max: l.Size, idx: idx}
	}

	curr := l.head
	var prev *listNode[T]
	for i := 0; i < idx; i++ {
		prev = curr
		curr = curr.next
	}

	if curr == l.head {
		l.head = l.head.next
	} else {
		prev.next = curr.next
	}
	val := curr.value
	curr = nil
	l.Size--

	return val, nil
}

func (l LinkedList[T]) Get(idx int) (T, error) {
	if idx < 0 || idx > l.Size-1 {
		return utils.GetZero[T](), &InvalidIndexError{max: l.Size, idx: idx}
	}

	curr := l.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	return curr.value, nil
}

func (l *LinkedList[T]) Set(idx int, item T) error {
	if idx < 0 || idx > l.Size-1 {
		return &InvalidIndexError{max: l.Size, idx: idx}
	}

	curr := l.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	curr.value = item

	return nil
}

func (l LinkedList[T]) Equals(c LinkedList[T]) bool {
	if l.Size != c.Size {
		return false
	}

	lCurr := l.head
	cCurr := c.head
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
