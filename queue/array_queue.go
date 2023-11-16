package queue

import "github.com/brandencmw/go-data-structures.git/utils"

// Array based implementation of an abstract queue
type ArrayQueue[T comparable] struct {
	contents []T // content[0] is front content[len-1] is rear
	// contents are enqueued to the rear and dequeued from the front
}

func NewArrayQueue[T comparable](contents ...T) *ArrayQueue[T] {
	return &ArrayQueue[T]{contents: contents}
}

func (q *ArrayQueue[T]) Clone() AbstractQueue[T] {
	c := make([]T, len(q.contents))
	copy(c, q.contents)
	return NewArrayQueue[T](c...)
}

func (q *ArrayQueue[T]) Enqueue(val T) {
	q.contents = append(q.contents, val)
}

// Removes and returns front of queue
func (q *ArrayQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return utils.GetZero[T](), &EmptyQueueError{method: "dequeue"}
	}

	val := q.contents[0]
	q.contents = q.contents[1:]
	return val, nil
}

// Like dequeue but doesn't remove the value from front
func (q ArrayQueue[T]) Front() (T, error) {
	if q.IsEmpty() {
		return utils.GetZero[T](), &EmptyQueueError{method: "front"}
	}
	return q.contents[0], nil
}

// Returns the most recently added value to the queue
func (q ArrayQueue[T]) Rear() (T, error) {
	if q.IsEmpty() {
		return utils.GetZero[T](), &EmptyQueueError{method: "front"}
	}
	return q.contents[len(q.contents)-1], nil
}

func (q ArrayQueue[T]) IsEmpty() bool {
	return len(q.contents) == 0
}
