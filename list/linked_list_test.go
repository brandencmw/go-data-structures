package list_test

import (
	"testing"

	"github.com/brandencmw/go-data-structures.git/list"
	"github.com/brandencmw/go-data-structures.git/utils"
)

// func createNodeOfint64() *list.ListNode[int64] {
// 	return &list.ListNode[int64]{
// 		Value: rand.Int63(),
// 		Next:  nil,
// 	}
// }

func createTestLinkedListOfLen(length uint) *list.LinkedList[int64] {
	newListContents := make([]int64, length)
	for i := uint(0); i < length; i++ {
		newListContents[i] = utils.GetRandomValueOfType(testVal).(int64)
	}
	return list.NewLinkedList[int64](newListContents...)
}

func createTestLinkedListOfContent(contents ...int64) *list.LinkedList[int64] {
	return list.NewLinkedList[int64](contents...)
}

func TestLinkedListContentsEqualForEqualLists(t *testing.T) {
	list1 := createTestLinkedListOfLen(0)
	list2 := createTestLinkedListOfContent(list1.Contents()...)

	if !list1.Equals(*list2) {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1.Contents(), list2.Contents())
	}
	if !list2.Equals(*list1) {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1.Contents(), list2.Contents())
	}

	list1 = createTestLinkedListOfLen(5)
	list2 = list1.Clone()
	if !list1.Equals(*list2) {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1.Contents(), list2.Contents())
	}
	if !list2.Equals(*list1) {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1.Contents(), list2.Contents())
	}

}

func TestLinkedListContentsEqualForDifferentLength(t *testing.T) {
	list1 := createTestLinkedListOfLen(5)
	list2 := createTestLinkedListOfLen(4)

	if list1.Equals(*list2) {
		t.Errorf("Wrong output, list contents should not be equal")
	}

	if list2.Equals(*list1) {
		t.Errorf("Wrong output, list contents should not be equal")
	}
}

func TestLinkedListContentsEqualForSameLengthDifferentContents(t *testing.T) {
	list1 := createTestLinkedListOfLen(BASE_LEN)
	list2 := createTestLinkedListOfLen(BASE_LEN)

	if list1.Equals(*list2) {
		t.Errorf("Wrong output, list 1 is %v, list 2 is %v", list1.Contents(), list2.Contents())
	}

	if list2.Equals(*list1) {
		t.Errorf("Wrong output, list contents should not be equal")
	}
}

func TestInsertToFrontOfPopulatedLinkedList(t *testing.T) {
	itemToAdd := getTestItem()

	originalList := createTestLinkedListOfLen(BASE_LEN)

	expectedListContents := append([]int64{itemToAdd}, originalList.Contents()...)
	expectedList := createTestLinkedListOfContent(expectedListContents...)
	resultLen := originalList.Insert(0, itemToAdd)
	if resultLen != expectedList.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size, resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList.Contents(), originalList.Contents())
	}
}

func TestInsertToRearOfPopulatedLinkedList(t *testing.T) {
	itemToAdd := getTestItem()

	originalList := createTestLinkedListOfLen(BASE_LEN)

	expectedList := createTestLinkedListOfContent(append(originalList.Contents(), itemToAdd)...)
	resultLen := originalList.Insert(BASE_LEN, itemToAdd)
	if resultLen != expectedList.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size, resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList.Contents(), originalList.Contents())
	}
}

func TestInsertToMiddleOfPopulatedLinkedList(t *testing.T) {
	originalList := createTestLinkedListOfLen(BASE_LEN)
	originalListContents := originalList.Contents()

	itemToAdd := getTestItem()
	const insertIndex = BASE_LEN / 2

	front := originalListContents[:insertIndex]
	rear := originalListContents[insertIndex:]

	expectedListContents := make([]int64, len(front))
	copy(expectedListContents, front)
	expectedListContents = append(expectedListContents, itemToAdd)
	expectedListContents = append(expectedListContents, rear...)
	expectedList := createTestLinkedListOfContent(expectedListContents...)

	resultLen := originalList.Insert(insertIndex, itemToAdd)
	if resultLen != expectedList.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size, resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList.Contents(), originalList.Contents())
	}
}

func TestInsertToEmptyLinkedList(t *testing.T) {
	itemToAdd := getTestItem()

	originalList := createTestLinkedListOfLen(0)

	expectedList := createTestLinkedListOfContent(itemToAdd)
	resultLen := originalList.Insert(0, itemToAdd)
	if resultLen != expectedList.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size, resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func TestInsertToInvalidIndexLinkedList(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Error("Code should have panicked")
		}
	}()

	itemToAdd := getTestItem()

	originalList := createTestLinkedListOfLen(0)
	initialLen := originalList.Size

	originalList.Insert(initialLen+1, itemToAdd)
}

func TestPrependToPopulatedLinkedList(t *testing.T) {
	itemToAdd := getTestItem()

	originalList := createTestLinkedListOfLen(BASE_LEN)

	expectedList := createTestLinkedListOfContent(append([]int64{itemToAdd}, originalList.Contents()...)...)
	resultLen := originalList.Prepend(itemToAdd)
	if resultLen != expectedList.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size, resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList.Contents(), originalList.Contents())
	}
}

func TestPrependToEmptyLinkedList(t *testing.T) {
	valueToAdd := getTestItem()

	originalList := createTestLinkedListOfLen(0)

	expectedList := createTestLinkedListOfContent(valueToAdd)
	resultLen := originalList.Prepend(valueToAdd)
	if resultLen != expectedList.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size, resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList.Contents(), originalList.Contents())
	}
}

func TestAppendToPopulatedLinkedList(t *testing.T) {
	valueToAdd := getTestItem()
	originalList := createTestLinkedListOfLen(BASE_LEN)

	expectedList := createTestLinkedListOfContent(append(originalList.Contents(), valueToAdd)...)
	resultLen := originalList.Append(valueToAdd)
	if resultLen != expectedList.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size, resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList.Contents(), originalList.Contents())
	}
}

func TestAppendToEmptyLinkedList(t *testing.T) {
	valueToAdd := getTestItem()

	originalList := createTestLinkedListOfLen(0)

	expectedList := createTestLinkedListOfContent(valueToAdd)
	resultLen := originalList.Append(valueToAdd)
	if resultLen != expectedList.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size, resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList.Contents(), originalList.Contents())
	}
}

func TestRemoveFromFrontOfPopulatedLinkedList(t *testing.T) {
	originalList := createTestLinkedListOfLen(BASE_LEN)

	expectedList := createTestLinkedListOfContent(originalList.Contents()[1:]...)
	expectedRemovedItem := originalList.Head()
	removedItem := originalList.Remove(0)
	if removedItem != expectedRemovedItem {
		t.Errorf("Wrong item removed: wanted %v, got %v", expectedRemovedItem, removedItem)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList.Contents(), originalList.Contents())
	}
}

func TestRemoveFromRearOfPopulatedLinkedList(t *testing.T) {
	originalList := createTestLinkedListOfLen(BASE_LEN)
	originalListContents := originalList.Contents()

	expectedList := createTestLinkedListOfContent(originalListContents[:BASE_LEN-1]...)
	expectedRemovedItem := originalListContents[BASE_LEN-1]
	removedItem := originalList.Remove(BASE_LEN - 1)
	if removedItem != expectedRemovedItem {
		t.Errorf("Wrong length: wanted %v, got %v", expectedRemovedItem, removedItem)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList.Contents(), originalList.Contents())
	}
}

func TestRemoveFromMiddleOfPopulatedLinkedList(t *testing.T) {
	const removalIndex = BASE_LEN / 2
	originalList := createTestLinkedListOfLen(BASE_LEN)

	originalListContents := originalList.Contents()
	expectedListContents := make([]int64, len(originalListContents[:removalIndex]))
	copy(expectedListContents, originalListContents[:removalIndex])
	expectedListContents = append(expectedListContents, originalListContents[removalIndex+1:]...)
	expectedList := createTestLinkedListOfContent(expectedListContents...)

	expectedRemovedItem := originalListContents[removalIndex]
	removedItem := originalList.Remove(removalIndex)
	if removedItem != expectedRemovedItem {
		t.Errorf("Wrong item removed: wanted %v, got %v", expectedRemovedItem, removedItem)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList.Contents(), originalList.Contents())
	}
}

func TestRemoveFromEmptyLinkedList(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Code should have panicked")
		}
	}()

	originalList := createTestLinkedListOfLen(0)
	originalList.Remove(0)
}

func TestGetFromPopulatedLinkedList(t *testing.T) {
	testList := createTestLinkedListOfLen(BASE_LEN)
	for index, expectedItem := range testList.Contents() {
		retrievedItem := testList.Get(index)
		if expectedItem != retrievedItem {
			t.Errorf("Didn't retrieve correct item: expected %v, got %v", expectedItem, retrievedItem)
		}
	}
}

func TestGetFromEmptyLinkedList(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Code should have panicked")
		}
	}()

	testList := createTestLinkedListOfLen(0)
	testList.Get(0)
}

func TestSetItemsInPopulatedLinkedList(t *testing.T) {
	testList := createTestLinkedListOfLen(BASE_LEN)
	expectedResultingList := createTestLinkedListOfLen(BASE_LEN)

	for index, expectedItem := range expectedResultingList.Contents() {
		testList.Set(index, expectedItem)
	}

	if !testList.Equals(*expectedResultingList) {
		t.Errorf("Didn't set item(s) correctly: expected %v, got %v", expectedResultingList.Contents(), testList.Contents())
	}

}

func TestSetItemInEmptyLinkedList(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Code should have panicked")
		}
	}()

	testList := createTestLinkedListOfLen(0)
	testList.Set(0, 0)
}
