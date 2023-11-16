package queue_test

import (
	"errors"
	"testing"

	"github.com/brandencmw/go-data-structures.git/queue"
	"github.com/brandencmw/go-data-structures.git/utils"
)

const BASE_LEN = 5

var testVal int64

func createTestArrayQueueOfLen(l uint) *queue.ArrayQueue[int64] {
	c := make([]int64, l)
	for i := uint(0); i < l; i++ {
		c[i] = getTestItem()
	}
	return queue.NewArrayQueue[int64](c...)
}

func createTestArrayQueueOfContent(c ...int64) *queue.ArrayQueue[int64] {
	return queue.NewArrayQueue[int64](c...)
}

func getTestItem() int64 {
	return utils.GetRandomValueOfType(testVal).(int64)
}
func TestEnqueueToEmptyArrayQueue(t *testing.T) {
	item := getTestItem()
	q1 := createTestArrayQueueOfLen(0)
	q2 := createTestArrayQueueOfContent(item)

	q1.Enqueue(item)
	if !queue.Equal(q1, q2) {
		t.Errorf("Wrong content. Expected %v, got %v", queue.AsSlice(q1), queue.AsSlice(q2))
	}
}

func TestEnqueueToPopulatedArrayQueue(t *testing.T) {
	item := getTestItem()
	q1 := createTestArrayQueueOfLen(BASE_LEN)
	q2 := createTestArrayQueueOfContent(append(queue.AsSlice(q1), item)...)

	q1.Enqueue(item)
	if !queue.Equal(q1, q2) {
		t.Errorf("Wrong content. Expected %v, got %v", queue.AsSlice(q2), queue.AsSlice(q1))
	}
}

func TestDequeueFromEmptyArrayQueue(t *testing.T) {
	q := createTestArrayQueueOfLen(0)

	_, err := q.Dequeue()
	var emptyError *queue.EmptyQueueError
	if !errors.As(err, &emptyError) {
		t.Errorf("Should have got empty queue error")
	}
}

func TestDequeueFromPopulatedArrayQueue(t *testing.T) {
	q1 := createTestArrayQueueOfLen(BASE_LEN)
	q2 := createTestArrayQueueOfContent(queue.AsSlice(q1)[1:]...)

	q1.Dequeue()

	if !queue.Equal(q1, q2) {
		t.Errorf("Wrong content. Expected %v, got %v", queue.AsSlice(q2), queue.AsSlice(q1))
	}
}

func TestFrontForEmptyArrayQueue(t *testing.T) {
	q := createTestArrayQueueOfLen(0)

	_, err := q.Front()
	var emptyError *queue.EmptyQueueError
	if !errors.As(err, &emptyError) {
		t.Errorf("Should have got empty queue error")
	}
}

func TestFrontForPopulatedArrayQueue(t *testing.T) {
	f := getTestItem()
	c := append([]int64{f}, queue.AsSlice(createTestArrayQueueOfLen(BASE_LEN))...)
	q := createTestArrayQueueOfContent(c...)

	retrieved, _ := q.Front()
	if f != retrieved {
		t.Errorf("Wrong front. Expected %v, got %v", f, retrieved)
	}
}

func TestRearForEmptyArrayQueue(t *testing.T) {
	q := createTestArrayQueueOfLen(0)

	_, err := q.Rear()
	var emptyError *queue.EmptyQueueError
	if !errors.As(err, &emptyError) {
		t.Error("Should have got empty queue error")
	}
}

func TestRearForPopulatedArrayQueue(t *testing.T) {
	r := getTestItem()
	c := append(queue.AsSlice(createTestArrayQueueOfLen(BASE_LEN)), r)
	q := createTestArrayQueueOfContent(c...)

	retrieved, _ := q.Rear()
	if r != retrieved {
		t.Errorf("Wrong rear. Expected %v, got %v", r, retrieved)
	}
}

func TestIsEmptyOnEmptyArrayQueue(t *testing.T) {
	q := createTestArrayQueueOfLen(0)
	if !q.IsEmpty() {
		t.Errorf("Wrong output. Queue is empty")
	}
}

func TestIsEmptyOnPopulatedArrayQueue(t *testing.T) {
	q := createTestArrayQueueOfLen(BASE_LEN)
	if q.IsEmpty() {
		t.Errorf("Wrong output. Queue is not empty")
	}
}

func TestArrayQueueCloneAllocatesSeparateMemory(t *testing.T) {
	q1 := createTestArrayQueueOfLen(BASE_LEN)
	q2 := q1.Clone()
	if q1 == q2 {
		t.Errorf("Clone should allocate separate memory. Both point to %v", q1)
	}
}

func TestArrayQueueCloneHasSameElements(t *testing.T) {
	c := make([]int64, BASE_LEN)
	for i := 0; i < BASE_LEN; i++ {
		c[i] = getTestItem()
	}
	q1 := createTestArrayQueueOfContent(c...)
	expected := createTestArrayQueueOfContent(c...)
	q2 := q1.Clone()
	if !queue.Equal(q2, expected) {
		t.Errorf("Contents not cloned. Expected %v, got %v", expected, q2)
	}
}
