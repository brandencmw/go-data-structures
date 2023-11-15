package stack_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/brandencmw/go-data-structures.git/stack"
	"github.com/brandencmw/go-data-structures.git/utils"
)

const BASE_LEN = 5

var testVal int64

func createTestArrayStackOfLen(len uint) *stack.ArrayStack[int64] {
	content := make([]int64, len)
	for i := uint(0); i < len; i++ {
		content[i] = utils.GetRandomValueOfType(testVal).(int64)
	}
	fmt.Printf("Content: %v\n", content)
	return stack.NewArrayStack[int64](content...)
}

func createTestArrayStackOfContent(content ...int64) *stack.ArrayStack[int64] {
	return stack.NewArrayStack[int64](content...)
}

func getTestItem() int64 {
	return utils.GetRandomValueOfType(testVal).(int64)
}

func TestPeekOnEmptyArrayStack(t *testing.T) {
	s := createTestArrayStackOfLen(0)
	_, err := s.Peek()

	var emptyError *stack.EmptyStackError
	if !errors.As(err, &emptyError) {
		t.Errorf("Should have got empty stack error")
	}
}

func TestPeekOnPopulatedArrayStack(t *testing.T) {
	expectedTop := getTestItem()
	s := createTestArrayStackOfContent(expectedTop)
	top, _ := s.Peek()
	if expectedTop != top {
		t.Errorf("Top is wrong. Expected %v, got %v", expectedTop, top)
	}
}

func TestPushToEmptyArrayStack(t *testing.T) {
	item := getTestItem()
	s1 := createTestArrayStackOfLen(0)
	s2 := createTestArrayStackOfContent(item)

	s1.Push(item)

	if !stack.Equal[int64](s1, s2) {
		t.Errorf("Stack has wrong content. Expected %v, got %v", stack.AsSlice(s1), stack.AsSlice(s2))
	}
}

func TestPushToPopulatedArrayStack(t *testing.T) {
	item := getTestItem()
	s1 := createTestArrayStackOfLen(BASE_LEN)
	s2Content := append([]int64{item}, stack.AsSlice(s1)...)
	s2 := createTestArrayStackOfContent(s2Content...)

	s1.Push(item)

	if !stack.Equal(s1, s2) {
		t.Errorf("Stack has wrong content. Expected %v, got %v", stack.AsSlice(s1), stack.AsSlice(s2))
	}
}

func TestPopFromEmptyArrayStack(t *testing.T) {
	s := createTestArrayStackOfLen(0)
	_, err := s.Pop()

	var emptyError *stack.EmptyStackError
	if !errors.As(err, &emptyError) {
		t.Errorf("Should have got empty stack error")
	}
}

func TestPopFromAlmostEmptyArrayStack(t *testing.T) {
	s1 := createTestArrayStackOfLen(1)
	s2 := createTestArrayStackOfLen(0)

	s1.Pop()

	if !stack.Equal(s1, s2) {
		t.Errorf("Stack should be empty, instead was %v", stack.AsSlice(s1))
	}
}

func TestPopFromPopulatedArrayStack(t *testing.T) {
	s1 := createTestArrayStackOfLen(BASE_LEN)
	s2 := createTestArrayStackOfContent(stack.AsSlice(s1)[1:]...)

	s1.Pop()

	if !stack.Equal(s1, s2) {
		t.Errorf("Contents not equal, expected %v, got %v", stack.AsSlice(s1), stack.AsSlice(s2))
	}
}

func TestArrayCloneCreatesSeparateMemory(t *testing.T) {
	s1 := createTestArrayStackOfLen(BASE_LEN)
	s2 := s1.Clone()

	if s1 == s2 {
		t.Errorf("Both stack reference same memory: %v", s1)
	}
}
