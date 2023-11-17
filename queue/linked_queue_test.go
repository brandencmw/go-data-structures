package queue_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/brandencmw/go-data-structures.git/queue"
)

func createTestLinkedQueueOfLen(l uint) *queue.LinkedQueue[int64] {
	c := make([]int64, l)
	for i := uint(0); i < l; i++ {
		c[i] = getTestItem()
	}
	return queue.NewLinkedQueue[int64](c...)
}

func createTestLinkedQueueOfContent(c ...int64) *queue.LinkedQueue[int64] {
	return queue.NewLinkedQueue[int64](c...)
}

func TestEnqueueToEmptyLinkedQueue(t *testing.T) {
	item := getTestItem()
	q1 := createTestLinkedQueueOfLen(0)
	q2 := createTestLinkedQueueOfContent(item)

	q1.Enqueue(item)
	if !queue.Equal(q1, q2) {
		t.Errorf("Wrong content. Expected %v, got %v", queue.AsSlice(q1), queue.AsSlice(q2))
	}
}

func TestEnqueueToPopulatedLinkedQueue(t *testing.T) {
	item := getTestItem()
	q1 := createTestLinkedQueueOfLen(BASE_LEN)
	q2 := createTestLinkedQueueOfContent(append(queue.AsSlice(q1), item)...)

	fmt.Printf("Q1: %v\n", queue.AsSlice(q1))

	q1.Enqueue(item)
	if !queue.Equal(q1, q2) {
		t.Errorf("Wrong content. Expected %v, got %v", queue.AsSlice(q2), queue.AsSlice(q1))
	}
}

func TestDequeueFromEmptyLinkedQueue(t *testing.T) {
	q := createTestLinkedQueueOfLen(0)

	_, err := q.Dequeue()
	var emptyError *queue.EmptyQueueError
	if !errors.As(err, &emptyError) {
		t.Errorf("Should have got empty queue error")
	}
}

func TestDequeueFromPopulatedLinkedQueue(t *testing.T) {
	q1 := createTestLinkedQueueOfLen(BASE_LEN)
	q2 := createTestLinkedQueueOfContent(queue.AsSlice(q1)[1:]...)

	q1.Dequeue()

	if !queue.Equal(q1, q2) {
		t.Errorf("Wrong content. Expected %v, got %v", queue.AsSlice(q2), queue.AsSlice(q1))
	}
}

func TestFrontForEmptyLinkedQueue(t *testing.T) {
	q := createTestLinkedQueueOfLen(0)

	_, err := q.Front()
	var emptyError *queue.EmptyQueueError
	if !errors.As(err, &emptyError) {
		t.Errorf("Should have got empty queue error")
	}
}

func TestFrontForPopulatedLinkedQueue(t *testing.T) {
	f := getTestItem()
	c := append([]int64{f}, queue.AsSlice(createTestLinkedQueueOfLen(BASE_LEN))...)
	q := createTestLinkedQueueOfContent(c...)

	retrieved, _ := q.Front()
	if f != retrieved {
		t.Errorf("Wrong front. Expected %v, got %v", f, retrieved)
	}
}

func TestRearForEmptyLinkedQueue(t *testing.T) {
	q := createTestLinkedQueueOfLen(0)

	_, err := q.Rear()
	var emptyError *queue.EmptyQueueError
	if !errors.As(err, &emptyError) {
		t.Error("Should have got empty queue error")
	}
}

func TestRearForPopulatedLinkedQueue(t *testing.T) {
	r := getTestItem()
	c := append(queue.AsSlice(createTestLinkedQueueOfLen(BASE_LEN)), r)
	q := createTestLinkedQueueOfContent(c...)

	retrieved, _ := q.Rear()
	if r != retrieved {
		t.Errorf("Wrong rear. Expected %v, got %v", r, retrieved)
	}
}

func TestIsEmptyOnEmptyLinkedQueue(t *testing.T) {
	q := createTestLinkedQueueOfLen(0)
	if !q.IsEmpty() {
		t.Errorf("Wrong output. Queue is empty")
	}
}

func TestIsEmptyOnPopulatedLinkedQueue(t *testing.T) {
	q := createTestLinkedQueueOfLen(BASE_LEN)
	if q.IsEmpty() {
		t.Errorf("Wrong output. Queue is not empty")
	}
}

func TestLinkedQueueCloneAllocatesSeparateMemory(t *testing.T) {
	q1 := createTestLinkedQueueOfLen(BASE_LEN)
	q2 := q1.Clone()
	if q1 == q2 {
		t.Errorf("Clone should allocate separate memory. Both point to %v", q1)
	}
}

func TestLinkedQueueCloneHasSameElements(t *testing.T) {
	c := make([]int64, BASE_LEN)
	for i := 0; i < BASE_LEN; i++ {
		c[i] = getTestItem()
	}
	q1 := createTestLinkedQueueOfContent(c...)
	expected := createTestLinkedQueueOfContent(c...)
	q2 := q1.Clone()
	if !queue.Equal(q2, expected) {
		t.Errorf("Contents not cloned. Expected %v, got %v", expected, q2)
	}
}
