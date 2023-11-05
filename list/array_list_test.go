package list_test

import (
	"testing"

	"github.com/brandencmw/go-data-structures.git/list"
)

type testType int64

func TestListContentsEqualForEqualLists(t *testing.T) {
	list1 := list.ArrayList[testType]{}
	list2 := list.ArrayList[testType]{}

	l1Equalsl2 := list1.ContentsEqualTo(list2)
	if !l1Equalsl2 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1, list2)
	}
	l2Equalsl1 := list2.ContentsEqualTo(list1)
	if !l2Equalsl1 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1, list2)
	}

	list1 = list.ArrayList[testType]{1, 2, 3, 4, 5}
	list2 = list.ArrayList[testType]{1, 2, 3, 4, 5}
	l1Equalsl2 = list1.ContentsEqualTo(list2)
	if !l1Equalsl2 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1, list2)
	}
	l2Equalsl1 = list2.ContentsEqualTo(list1)
	if !l2Equalsl1 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1, list2)
	}

}

func TestListContentsEqualForDifferentLength(t *testing.T) {
	list1 := list.ArrayList[testType]{1, 2, 3, 4, 5}
	list2 := list.ArrayList[testType]{1, 2, 3, 4}

	l1Equalsl2 := list1.ContentsEqualTo(list2)
	if l1Equalsl2 {
		t.Errorf("Wrong output, list contents are not equal")
	}

	l2Equalsl1 := list2.ContentsEqualTo(list1)
	if l2Equalsl1 {
		t.Errorf("Wrong output, list contents are not equal")
	}
}

func TestListContentsEqualForSameLengthDifferentContents(t *testing.T) {
	list1 := list.ArrayList[testType]{1, 2, 3, 4, 5}
	list2 := list.ArrayList[testType]{1, 2, 3, 4, 6}

	l1Equalsl2 := list1.ContentsEqualTo(list2)
	if l1Equalsl2 {
		t.Errorf("Wrong output, list contents are not equal")
	}

	l2Equalsl1 := list2.ContentsEqualTo(list1)
	if l2Equalsl1 {
		t.Errorf("Wrong output, list contents are not equal")
	}
}

func TestInsertToFrontOfPopulatedList(t *testing.T) {
	const itemToAdd testType = 5

	originalList := list.ArrayList[testType]{1, 4, 5, 7, 5}
	initialLen := len(originalList)

	expectedList := list.ArrayList[testType]{itemToAdd, 1, 4, 5, 7, 5}
	resultLen := originalList.Insert(itemToAdd, 0)
	if resultLen != initialLen+1 {
		t.Errorf("Wrong length: wanted %v, got %v", initialLen+1, resultLen)
	}

	if !originalList.ContentsEqualTo(expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func TestInsertToRearOfPopulatedList(t *testing.T) {
	const itemToAdd testType = 5

	originalList := list.ArrayList[testType]{1, 4, 5, 7, 5}
	initialLen := len(originalList)

	expectedList := list.ArrayList[testType]{1, 4, 5, 7, 5, itemToAdd}
	resultLen := originalList.Insert(itemToAdd, initialLen)
	if resultLen != initialLen+1 {
		t.Errorf("Wrong length: wanted %v, got %v", initialLen+1, resultLen)
	}

	if !originalList.ContentsEqualTo(expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func TestInsertToMiddleOfPopulatedList(t *testing.T) {
	const itemToAdd testType = 5

	originalList := list.ArrayList[testType]{1, 4, 5, 7, 5}
	initialLen := len(originalList)
	insertIndex := initialLen / 2

	expectedList := append(originalList[:insertIndex], itemToAdd)
	expectedList = append(expectedList, originalList[insertIndex:]...)

	resultLen := originalList.Insert(itemToAdd, insertIndex)
	if resultLen != initialLen+1 {
		t.Errorf("Wrong length: wanted %v, got %v", initialLen+1, resultLen)
	}

	if !originalList.ContentsEqualTo(expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func TestInsertToEmptyList(t *testing.T) {
	const itemToAdd testType = 5

	originalList := list.ArrayList[testType]{}
	initialLen := len(originalList)

	expectedList := list.ArrayList[testType]{itemToAdd}
	resultLen := originalList.Insert(itemToAdd, 0)
	if resultLen != initialLen+1 {
		t.Errorf("Wrong length: wanted %v, got %v", initialLen+1, resultLen)
	}

	if !originalList.ContentsEqualTo(expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func checkForPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("Code should have panicked")
	}
}

func TestInsertToInvalidIndex(t *testing.T) {

	defer checkForPanic(t)

	const itemToAdd testType = 5

	originalList := list.ArrayList[testType]{}
	initialLen := len(originalList)

	originalList.Insert(itemToAdd, initialLen+1)
}

func TestPrependToPopulatedList(t *testing.T) {
	const itemToAdd testType = 0
	originalList := list.ArrayList[testType]{1, 2, 3, 4, 5}
	initialLen := len(originalList)

	expectedList := append(list.ArrayList[testType]{itemToAdd}, originalList...)
	resultLen := originalList.Prepend(itemToAdd)
	if resultLen != initialLen+1 {
		t.Errorf("Wrong length: wanted %v, got %v", initialLen+1, resultLen)
	}

	if !originalList.ContentsEqualTo(expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func TestPrependToEmptyList(t *testing.T) {
	const itemToAdd testType = 0
	originalList := list.ArrayList[testType]{}
	initialLen := len(originalList)

	expectedList := list.ArrayList[testType]{itemToAdd}
	resultLen := originalList.Prepend(itemToAdd)
	if resultLen != initialLen+1 {
		t.Errorf("Wrong length: wanted %v, got %v", initialLen+1, resultLen)
	}

	if !originalList.ContentsEqualTo(expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}
