package list_test

import (
	"errors"
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
	original := createTestArrayListOfLen(BASE_LEN)
	cloned := original.Clone()

	if original == cloned {
		t.Errorf("Cloned list should not reference same memory. Both reference %v", original)
	}
}

func TestArrayListCloneGeneratesSameContents(t *testing.T) {
	original := createTestArrayListOfLen(BASE_LEN)
	cloned := original.Clone()
	if original.Size() != cloned.Size() {
		t.Errorf("Lists not same size. Original is %v, cloned is %v", original.Size(), cloned.Size())
	}

	originalContents := original.Contents()
	clonedContents := cloned.Contents()
	for idx := range originalContents {
		if originalContents[idx] != clonedContents[idx] {
			t.Errorf("Item not cloned. Expected %v, got %v", originalContents[idx], clonedContents[idx])
		}
	}
}

func TestArrayListContentsEqualForEqualLists(t *testing.T) {
	l1 := createTestArrayListOfLen(0)
	l2 := l1.Clone()
	l1Equalsl2 := l1.Equals(*l2)
	if !l1Equalsl2 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", l1.Contents(), l2.Contents())
	}

	l1 = createTestArrayListOfLen(5)
	l2 = l1.Clone()
	l1Equalsl2 = l1.Equals(*l2)
	if !l1Equalsl2 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", l1.Contents(), l2.Contents())
	}
}

func TestArrayListContentsEqualForDifferentLength(t *testing.T) {

	l1 := createTestArrayListOfLen(BASE_LEN)
	l2 := createTestArrayListOfContent(l1.Contents()[:BASE_LEN-1]...)

	l1Equalsl2 := l1.Equals(*l2)
	if l1Equalsl2 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", l1.Contents(), l2.Contents())
	}

	l2Equalsl1 := l2.Equals(*l1)
	if l2Equalsl1 {
		t.Errorf("Wrong output, l2 is %v and l1 is %v", l2.Contents(), l1.Contents())
	}
}

func TestArrayListContentsEqualForSameLengthDifferentContents(t *testing.T) {

	l1 := createTestArrayListOfLen(BASE_LEN)
	l2 := createTestArrayListOfLen(BASE_LEN)

	l1Equalsl2 := l1.Equals(*l2)
	if l1Equalsl2 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", l1.Contents(), l2.Contents())
	}

	l2Equalsl1 := l2.Equals(*l1)
	if l2Equalsl1 {
		t.Errorf("Wrong output, l2 is %v and l1 is %v", l2.Contents(), l2.Contents())
	}
}

func TestInsertToFrontOfPopulatedArrayList(t *testing.T) {
	itemToAdd := getTestItem()
	original := createTestArrayListOfLen(BASE_LEN)
	expected := append([]int64{itemToAdd}, original.Contents()...)

	original.Insert(itemToAdd, 0)
	if original.Size() != len(expected) {
		t.Errorf("Wrong length: wanted %v, got %v", len(expected), original.Size())
	}

	if !original.Equals(*createTestArrayListOfContent(expected...)) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected, original.Contents())
	}
}

func TestInsertToRearOfPopulatedArrayList(t *testing.T) {
	item := getTestItem()
	original := createTestArrayListOfLen(BASE_LEN)
	expected := append(original.Contents(), item)

	original.Insert(item, BASE_LEN)
	if original.Size() != len(expected) {
		t.Errorf("Wrong length: wanted %v, got %v", len(expected), original.Size())
	}

	if !original.Equals(*createTestArrayListOfContent(expected...)) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected, original.Contents())
	}
}

func TestInsertToMiddleOfPopulatedArrayList(t *testing.T) {
	item := getTestItem()
	original := createTestArrayListOfLen(BASE_LEN)
	idx := BASE_LEN / 2

	originalContents := original.Contents()
	front := originalContents[:idx]
	expectedContents := make([]int64, len(front))
	copy(expectedContents, front)
	expectedContents = append(expectedContents, item)
	expectedContents = append(expectedContents, originalContents[idx:]...)

	original.Insert(item, idx)
	if original.Size() != len(expectedContents) {
		t.Errorf("Wrong length: wanted %v, got %v", len(expectedContents), original.Size())
	}

	if !original.Equals(*createTestArrayListOfContent(expectedContents...)) {
		t.Errorf("Wrong list contents: expected %v, got %v", expectedContents, original.Contents())
	}
}

func TestInsertToEmptyArrayList(t *testing.T) {
	item := getTestItem()
	original := createTestArrayListOfLen(0)
	expected := createTestArrayListOfContent(item)

	original.Insert(item, 0)
	if original.Size() != expected.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size(), original.Size())
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestInsertToInvalidIndex(t *testing.T) {
	item := getTestItem()
	original := createTestArrayListOfLen(BASE_LEN)
	err := original.Insert(item, BASE_LEN+1)

	var indexError *list.InvalidIndexError
	if !errors.As(err, &indexError) {
		t.Error("Should have gotten invalid index error")
	}
}

func TestPrependToPopulatedArrayList(t *testing.T) {
	item := getTestItem()
	original := createTestArrayListOfLen(BASE_LEN)

	expectedContents := append([]int64{item}, original.Contents()...)
	expected := createTestArrayListOfContent(expectedContents...)

	original.Prepend(item)
	if original.Size() != expected.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size(), original.Size())
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestPrependToEmptyArrayList(t *testing.T) {
	item := getTestItem()
	original := createTestArrayListOfLen(0)

	expected := createTestArrayListOfContent(item)
	original.Prepend(item)
	if original.Size() != expected.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size(), original.Size())
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected, original)
	}
}

func TestAppendToPopulatedArrayList(t *testing.T) {
	item := getTestItem()
	original := createTestArrayListOfLen(BASE_LEN)

	expectedContents := append(original.Contents(), item)
	expected := createTestArrayListOfContent(expectedContents...)

	original.Append(item)
	if original.Size() != expected.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size(), original.Size())
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestAppendToEmptyArrayList(t *testing.T) {
	item := getTestItem()
	original := createTestArrayListOfLen(0)

	expected := createTestArrayListOfContent(item)
	original.Append(item)
	if original.Size() != expected.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size(), original.Size())
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected, original)
	}
}

func TestRemoveFromFrontOfPopulatedArrayList(t *testing.T) {
	original := createTestArrayListOfLen(BASE_LEN)
	expected := createTestArrayListOfContent(original.Contents()[1:]...)

	original.Remove(0)
	if original.Size() != expected.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size(), original.Size())
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected, original)
	}
}

func TestRemoveFromRearOfPopulatedArrayList(t *testing.T) {
	original := createTestArrayListOfLen(BASE_LEN)
	expected := createTestArrayListOfContent(original.Contents()[:BASE_LEN-1]...)

	original.Remove(BASE_LEN - 1)
	if original.Size() != expected.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size(), original.Size())
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected, original)
	}
}

func TestRemoveFromMiddleOfPopulatedArrayList(t *testing.T) {
	original := createTestArrayListOfLen(BASE_LEN)
	idx := BASE_LEN / 2

	originalContents := original.Contents()
	expectedContents := make([]int64, len(originalContents[:idx]))
	copy(expectedContents, originalContents[:idx])
	expectedContents = append(expectedContents, originalContents[idx+1:]...)
	expected := createTestArrayListOfContent(expectedContents...)

	original.Remove(idx)
	if original.Size() != expected.Size() {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size(), original.Size())
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestRemoveFromEmptyArrayList(t *testing.T) {
	original := createTestArrayListOfLen(0)
	err := original.Remove(0)
	if !errors.Is(err, list.ErrEmptyList) {
		t.Errorf("Should have got empty list error")
	}
}

func TestGetFromPopulatedArrayList(t *testing.T) {
	l := createTestArrayListOfLen(BASE_LEN)
	for i, item := range l.Contents() {
		retrieved, err := l.Get(i)
		if err != nil {
			t.Fatalf("Received error in get: %v", err.Error())
		}
		if item != retrieved {
			t.Errorf("Didn't retrieve correct item: expected %v, got %v", item, retrieved)
		}
	}
}

func TestGetFromEmptyArrayList(t *testing.T) {
	l := createTestArrayListOfLen(0)
	_, err := l.Get(0)

	var indexError *list.InvalidIndexError
	if !errors.As(err, &indexError) {
		t.Error("Should have got invalid index error")
	}
}

func TestSetItemsInPopulatedArrayList(t *testing.T) {
	l := createTestArrayListOfLen(BASE_LEN)
	expected := createTestArrayListOfLen(BASE_LEN)

	for i, item := range expected.Contents() {
		l.Set(i, item)
	}

	if !l.Equals(*expected) {
		t.Errorf("Didn't set items correctly: expected %v, got %v", expected.Contents(), l.Contents())
	}

}

func TestSetItemInEmptyArrayList(t *testing.T) {
	testList := createTestArrayListOfLen(0)
	err := testList.Set(0, 0)

	var indexError *list.InvalidIndexError
	if !errors.As(err, &indexError) {
		t.Error("Should have got invalid index error")
	}
}
