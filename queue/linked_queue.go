package queue

import (
	"fmt"

	"github.com/brandencmw/go-data-structures.git/utils"
)

type QueueNode[T comparable] struct {
	val  T
	next *QueueNode[T]
	prev *QueueNode[T]
}

type LinkedQueue[T comparable] struct {
	front *QueueNode[T]
	rear  *QueueNode[T]
}

func NewLinkedQueue[T comparable](c ...T) *LinkedQueue[T] {
	if len(c) == 0 {
		return &LinkedQueue[T]{front: nil, rear: nil}
	}

	fmt.Printf("Content: %v\n", c)
	front := &QueueNode[T]{val: c[0], next: nil, prev: nil}
	q := &LinkedQueue[T]{front: front, rear: nil}
	curr := front
	for _, item := range c[1:] {
		n := &QueueNode[T]{val: item, next: curr, prev: nil}
		curr.prev = n
		curr = curr.prev
	}
	q.rear = curr
	return q
}

func (q *LinkedQueue[T]) Clone() AbstractQueue[T] {
	c := &LinkedQueue[T]{front: nil, rear: nil}
	if q.IsEmpty() {
		return c
	}

	qCurr := q.front
	cCurr := &QueueNode[T]{next: nil, prev: nil, val: qCurr.val}
	c.front = cCurr
	qCurr = qCurr.prev
	for qCurr != nil {
		cCurr.prev = &QueueNode[T]{next: cCurr, prev: nil, val: qCurr.val}
		cCurr = cCurr.prev
		qCurr = qCurr.prev
	}
	c.rear = cCurr
	return c
}

func (q LinkedQueue[T]) IsEmpty() bool {
	return q.front == nil
}

func (q *LinkedQueue[T]) Enqueue(val T) {
	n := &QueueNode[T]{val: val, next: q.rear, prev: nil}
	if q.IsEmpty() {
		q.front = n // First element in queue will be both front and rear
	} else {
		q.rear.prev = n
	}
	q.rear = n // New node will always be enqueued to rear
}

func (q *LinkedQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return utils.GetZero[T](), &EmptyQueueError{method: "dequeue"}
	}

	v := q.front.val
	if q.front == q.rear {
		q.rear = nil
	}
	q.front = q.front.prev
	return v, nil
}

func (q LinkedQueue[T]) Front() (T, error) {
	if q.IsEmpty() {
		return utils.GetZero[T](), &EmptyQueueError{method: "front"}
	}
	return q.front.val, nil
}

func (q LinkedQueue[T]) Rear() (T, error) {
	if q.IsEmpty() {
		return utils.GetZero[T](), &EmptyQueueError{method: "rear"}
	}
	return q.rear.val, nil
}
