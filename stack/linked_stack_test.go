package stack_test

import (
	"errors"
	"testing"

	"github.com/brandencmw/go-data-structures.git/stack"
	"github.com/brandencmw/go-data-structures.git/utils"
)

func createTestLinkedStackOfLen(len uint) *stack.LinkedStack[int64] {
	content := make([]int64, len)
	for i := uint(0); i < len; i++ {
		content[i] = utils.GetRandomValueOfType(testVal).(int64)
	}

	return stack.NewLinkedStack[int64](content...)
}

func createTestLinkedStackOfContent(content ...int64) *stack.LinkedStack[int64] {
	return stack.NewLinkedStack[int64](content...)
}

func TestPeekOnEmptyLinkedStack(t *testing.T) {
	s := createTestLinkedStackOfLen(0)
	_, err := s.Peek()

	var emptyError *stack.EmptyStackError
	if !errors.As(err, &emptyError) {
		t.Errorf("Should have got empty stack error")
	}
}

func TestPeekOnPopulatedLinkedStack(t *testing.T) {
	expectedTop := getTestItem()
	s := createTestLinkedStackOfContent(expectedTop)
	top, _ := s.Peek()
	if expectedTop != top {
		t.Errorf("Top is wrong. Expected %v, got %v", expectedTop, top)
	}
}

func TestPushToEmptyLinkedStack(t *testing.T) {
	item := getTestItem()
	s1 := createTestLinkedStackOfLen(0)
	s2 := createTestLinkedStackOfContent(item)

	s1.Push(item)

	if !stack.Equal[int64](s1, s2) {
		t.Errorf("Stack has wrong content. Expected %v, got %v", stack.AsSlice(s1), stack.AsSlice(s2))
	}
}

func TestPushToPopulatedLinkedStack(t *testing.T) {
	item := getTestItem()
	s1 := createTestLinkedStackOfLen(BASE_LEN)
	s2Content := append([]int64{item}, stack.AsSlice(s1)...)
	s2 := createTestLinkedStackOfContent(s2Content...)

	s1.Push(item)

	if !stack.Equal(s1, s2) {
		t.Errorf("Stack has wrong content. Expected %v, got %v", stack.AsSlice(s1), stack.AsSlice(s2))
	}
}

func TestPopFromEmptyLinkedStack(t *testing.T) {
	s := createTestLinkedStackOfLen(0)
	_, err := s.Pop()

	var emptyError *stack.EmptyStackError
	if !errors.As(err, &emptyError) {
		t.Errorf("Should have got empty stack error")
	}
}

func TestPopFromAlmostEmptyLinkedStack(t *testing.T) {
	s1 := createTestArrayStackOfLen(1)
	s2 := createTestArrayStackOfLen(0)

	s1.Pop()

	if !stack.Equal(s1, s2) {
		t.Errorf("Stack should be empty, instead was %v", stack.AsSlice(s1))
	}
}

func TestPopFromPopulatedLinkedStack(t *testing.T) {
	s1 := createTestArrayStackOfLen(BASE_LEN)
	s2 := createTestArrayStackOfContent(stack.AsSlice(s1)[1:]...)

	s1.Pop()

	if !stack.Equal(s1, s2) {
		t.Errorf("Contents not equal, expected %v, got %v", stack.AsSlice(s1), stack.AsSlice(s2))
	}
}

func TestLinkedCloneCreatesSeparateMemory(t *testing.T) {
	s1 := createTestLinkedStackOfLen(BASE_LEN)
	s2 := s1.Clone()

	if s1 == s2 {
		t.Errorf("Both stack reference same memory: %v", s1)
	}
}
