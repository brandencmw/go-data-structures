package queue

import "fmt"

type AbstractQueue[T comparable] interface {
	Enqueue(T)
	Dequeue() (T, error)
	Front() (T, error)
	Rear() (T, error)
	IsEmpty() bool
	Clone() AbstractQueue[T]
}

// Returns copy of queue contents as a slice
// slice[0] is front, slice[len(slice)-1] is rear
func AsSlice[T comparable](q AbstractQueue[T]) []T {
	l := make([]T, 0)
	c := q.Clone() // Operate on clone to not modify original queue
	for !c.IsEmpty() {
		val, err := c.Dequeue()
		if err != nil {
			panic(err.Error())
		}
		l = append(l, val)
	}
	return l
}

func Equal[T comparable](q1, q2 AbstractQueue[T]) bool {
	c1 := q1.Clone() //Operate on clones to not modify originals
	c2 := q2.Clone()

	for !c1.IsEmpty() && !c2.IsEmpty() {
		v1, err := c1.Dequeue()
		if err != nil {
			panic(err.Error())
		}
		v2, err := c2.Dequeue()
		if err != nil {
			panic(err.Error())
		}
		if v1 != v2 {
			return false
		}
	}
	return c1.IsEmpty() && c2.IsEmpty()
}

type EmptyQueueError struct {
	method string
}

func (e *EmptyQueueError) Error() string {
	return fmt.Sprintf("Cannot use %v on an empty queue", e.method)
}
