package list_test

import (
	"testing"

	"github.com/brandencmw/go-data-structures.git/list"
	"github.com/brandencmw/go-data-structures.git/utils"
)

const BASE_LEN = 5

var testVal int64 //Used for generating test values of consistent type

func createTestArrayListOfLen(length uint) *list.ArrayList[int64] {
	newListContents := make([]int64, length)
	for i := uint(0); i < length; i++ {
		newListContents[i] = utils.GetRandomValueOfType(testVal).(int64)
	}
	return list.NewArrayList[int64](newListContents...)
}

func createTestArrayListOfContent(contents ...int64) *list.ArrayList[int64] {
	return list.NewArrayList[int64](contents...)
}

func getTestItem() int64 {
	return utils.GetRandomValueOfType(testVal).(int64)
}

func TestArrayListCloneAllocatesDifferentMemory(t *testing.T) {
	originalList := createTestArrayListOfLen(BASE_LEN)
	clonedList := originalList.Clone()

	if originalList == clonedList {
		t.Errorf("Cloned list should not reference same memory, both reference %v", originalList)
	}
}

func TestArrayListCloneGeneratesSameContents(t *testing.T) {
	originalList := createTestArrayListOfLen(BASE_LEN)
	clonedList := originalList.Clone()

	originalContents := originalList.Contents()
	clonedContents := clonedList.Contents()
	for idx := range originalContents {
		if originalContents[idx] != clonedContents[idx] {
			t.Errorf("Item not cloned. Expected %v, got %v", originalContents[idx], clonedContents[idx])
		}
	}
}

func TestArrayListContentsEqualForEqualLists(t *testing.T) {
	list1 := createTestArrayListOfLen(BASE_LEN)
	list2 := createTestArrayListOfContent(list1.Contents()...)

	l1Equalsl2 := list1.Equals(*list2)
	if !l1Equalsl2 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1, list2)
	}
	l2Equalsl1 := list2.Equals(*list1)
	if !l2Equalsl1 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1, list2)
	}

	list1 = createTestArrayListOfLen(5)
	list2 = list1.Clone()
	l1Equalsl2 = list1.Equals(*list2)
	if !l1Equalsl2 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1, list2)
	}
	l2Equalsl1 = list2.Equals(*list1)
	if !l2Equalsl1 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1, list2)
	}

}

func TestArrayListContentsEqualForDifferentLength(t *testing.T) {

	list1 := createTestArrayListOfLen(BASE_LEN)
	list2 := createTestArrayListOfContent(list1.Contents()[:BASE_LEN-1]...)

	l1Equalsl2 := list1.Equals(*list2)
	if l1Equalsl2 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1.Contents(), list2.Contents())
	}

	l2Equalsl1 := list2.Equals(*list1)
	if l2Equalsl1 {
		t.Errorf("Wrong output, l2 is %v and l1 is %v", list2.Contents(), list1.Contents())
	}
}

func TestArrayListContentsEqualForSameLengthDifferentContents(t *testing.T) {

	list1 := createTestArrayListOfLen(BASE_LEN)
	list2 := createTestArrayListOfLen(BASE_LEN)

	l1Equalsl2 := list1.Equals(*list2)
	if l1Equalsl2 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1.Contents(), list2.Contents())
	}

	l2Equalsl1 := list2.Equals(*list1)
	if l2Equalsl1 {
		t.Errorf("Wrong output, l2 is %v and l1 is %v", list2.Contents(), list2.Contents())
	}
}

func TestInsertToFrontOfPopulatedArrayList(t *testing.T) {
	itemToAdd := getTestItem()

	originalList := createTestArrayListOfLen(BASE_LEN)

	expectedListContents := append([]int64{itemToAdd}, originalList.Contents()...)
	resultLen := originalList.Insert(itemToAdd, 0)
	if resultLen != BASE_LEN+1 {
		t.Errorf("Wrong length: wanted %v, got %v", BASE_LEN+1, resultLen)
	}

	if !originalList.Equals(*createTestArrayListOfContent(expectedListContents...)) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedListContents, originalList.Contents())
	}
}

func TestInsertToRearOfPopulatedArrayList(t *testing.T) {
	itemToAdd := getTestItem()

	originalList := createTestArrayListOfLen(BASE_LEN)

	expectedListContents := append(originalList.Contents(), itemToAdd)
	resultLen := originalList.Insert(itemToAdd, BASE_LEN)
	if resultLen != BASE_LEN+1 {
		t.Errorf("Wrong length: wanted %v, got %v", BASE_LEN+1, resultLen)
	}

	if !originalList.Equals(*createTestArrayListOfContent(expectedListContents...)) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedListContents, originalList.Contents())
	}
}

func TestInsertToMiddleOfPopulatedArrayList(t *testing.T) {
	itemToAdd := getTestItem()

	originalList := createTestArrayListOfLen(BASE_LEN)

	insertIndex := BASE_LEN / 2

	originalListContents := originalList.Contents()
	front := originalListContents[:insertIndex]
	expectedListContents := make([]int64, len(front))
	copy(expectedListContents, front)
	expectedListContents = append(expectedListContents, itemToAdd)
	expectedListContents = append(expectedListContents, originalListContents[insertIndex:]...)

	resultLen := originalList.Insert(itemToAdd, insertIndex)
	if resultLen != BASE_LEN+1 {
		t.Errorf("Wrong length: wanted %v, got %v", BASE_LEN+1, resultLen)
	}

	if !originalList.Equals(*createTestArrayListOfContent(expectedListContents...)) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedListContents, originalList.Contents())
	}
}

func TestInsertToEmptyArrayList(t *testing.T) {
	itemToAdd := getTestItem()

	originalList := createTestArrayListOfLen(0)

	expectedList := createTestArrayListOfContent(itemToAdd)
	resultLen := originalList.Insert(itemToAdd, 0)
	if resultLen != expectedList.Size() {
		t.Errorf("Wrong length: wanted 1, got %v", resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func TestInsertToInvalidIndex(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Error("Code should have panicked")
		}
	}()

	itemToAdd := getTestItem()
	originalList := createTestArrayListOfLen(BASE_LEN)
	originalList.Insert(itemToAdd, BASE_LEN+1)
}

func TestPrependToPopulatedArrayList(t *testing.T) {
	itemToAdd := getTestItem()
	originalList := createTestArrayListOfLen(BASE_LEN)

	expectedListContents := append([]int64{itemToAdd}, originalList.Contents()...)
	expectedList := createTestArrayListOfContent(expectedListContents...)

	resultLen := originalList.Prepend(itemToAdd)
	if resultLen != expectedList.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size(), resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList.Contents(), originalList.Contents())
	}
}

func TestPrependToEmptyArrayList(t *testing.T) {
	itemToAdd := getTestItem()
	originalList := createTestArrayListOfLen(0)

	expectedList := createTestArrayListOfContent(itemToAdd)
	resultLen := originalList.Prepend(itemToAdd)
	if resultLen != expectedList.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size(), resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func TestAppendToPopulatedArrayList(t *testing.T) {
	itemToAdd := getTestItem()
	originalList := createTestArrayListOfLen(BASE_LEN)

	expectedListContents := append(originalList.Contents(), itemToAdd)
	expectedList := createTestArrayListOfContent(expectedListContents...)

	resultLen := originalList.Append(itemToAdd)
	if resultLen != expectedList.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size(), resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func TestAppendToEmptyArrayList(t *testing.T) {
	itemToAdd := getTestItem()
	originalList := createTestArrayListOfLen(0)

	expectedList := createTestArrayListOfContent(itemToAdd)
	resultLen := originalList.Append(itemToAdd)
	if resultLen != expectedList.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size(), resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func TestRemoveFromFrontOfPopulatedArrayList(t *testing.T) {
	originalList := createTestArrayListOfLen(BASE_LEN)

	expectedList := createTestArrayListOfContent(originalList.Contents()[1:]...)
	resultLen := originalList.Remove(0)
	if resultLen != expectedList.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size(), resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func TestRemoveFromRearOfPopulatedArrayList(t *testing.T) {
	originalList := createTestArrayListOfLen(BASE_LEN)

	expectedList := createTestArrayListOfContent(originalList.Contents()[:BASE_LEN-1]...)
	resultLen := originalList.Remove(BASE_LEN - 1)
	if resultLen != expectedList.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size(), resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func TestRemoveFromMiddleOfPopulatedArrayList(t *testing.T) {
	originalList := createTestArrayListOfLen(BASE_LEN)

	removalIndex := BASE_LEN / 2

	originalListContents := originalList.Contents()
	expectedListContents := make([]int64, len(originalListContents[:removalIndex]))
	copy(expectedListContents, originalListContents[:removalIndex])
	expectedListContents = append(expectedListContents, originalListContents[removalIndex+1:]...)
	expectedList := createTestArrayListOfContent(expectedListContents...)

	resultLen := originalList.Remove(removalIndex)
	if resultLen != expectedList.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expectedList.Size(), resultLen)
	}

	if !originalList.Equals(*expectedList) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedList, originalList)
	}
}

func TestRemoveFromEmptyArrayList(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Code should have panicked")
		}
	}()

	originalList := createTestArrayListOfLen(0)
	originalList.Remove(0)
}

func TestGetFromPopulatedArrayList(t *testing.T) {
	testList := createTestArrayListOfLen(BASE_LEN)
	for index, expectedItem := range testList.Contents() {
		retrievedItem := testList.Get(index)
		if expectedItem != retrievedItem {
			t.Errorf("Didn't retrieve correct item: expected %v, got %v", expectedItem, retrievedItem)
		}
	}
}

func TestGetFromEmptyArrayList(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Code should have panicked")
		}
	}()

	testList := createTestArrayListOfLen(0)
	testList.Get(0)
}

func TestSetItemsInPopulatedArrayList(t *testing.T) {
	testList := createTestArrayListOfLen(BASE_LEN)
	expectedResultingList := createTestArrayListOfLen(BASE_LEN)

	for index, expectedItem := range expectedResultingList.Contents() {
		testList.Set(index, expectedItem)
	}

	if !testList.Equals(*expectedResultingList) {
		t.Errorf("Didn't set items correctly: expected %v, got %v", expectedResultingList.Contents(), testList.Contents())
	}

}

func TestSetItemInEmptyArrayList(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Code should have panicked")
		}
	}()

	testList := createTestArrayListOfLen(0)
	testList.Set(0, 0)
}
