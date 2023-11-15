package list_test

import (
	"errors"
	"testing"

	"github.com/brandencmw/go-data-structures.git/list"
	"github.com/brandencmw/go-data-structures.git/utils"
)

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
	l1 := createTestLinkedListOfLen(0)
	l2 := l1.Clone()

	if !l1.Equals(*l2) {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", l1.Contents(), l2.Contents())
	}

	l1 = createTestLinkedListOfLen(5)
	l2 = l1.Clone()
	if !l1.Equals(*l2) {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", l1.Contents(), l2.Contents())
	}
}

func TestLinkedListContentsEqualForDifferentLength(t *testing.T) {
	l1 := createTestLinkedListOfLen(BASE_LEN)
	l2 := createTestLinkedListOfContent(l1.Contents()[:BASE_LEN-1]...)

	if l1.Equals(*l2) {
		t.Errorf("Wrong output, list contents should not be equal")
	}
}

func TestLinkedListContentsEqualForSameLengthDifferentContents(t *testing.T) {
	l1 := createTestLinkedListOfLen(BASE_LEN)
	l2 := createTestLinkedListOfLen(BASE_LEN)

	if l1.Equals(*l2) {
		t.Errorf("Wrong output, list 1 is %v, list 2 is %v", l1.Contents(), l2.Contents())
	}
}

func TestInsertToFrontOfPopulatedLinkedList(t *testing.T) {
	item := getTestItem()

	original := createTestLinkedListOfLen(BASE_LEN)

	expectedContents := append([]int64{item}, original.Contents()...)
	expected := createTestLinkedListOfContent(expectedContents...)
	original.Insert(0, item)
	if original.Size != expected.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size, original.Size)
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestInsertToRearOfPopulatedLinkedList(t *testing.T) {
	item := getTestItem()

	original := createTestLinkedListOfLen(BASE_LEN)

	expected := createTestLinkedListOfContent(append(original.Contents(), item)...)
	original.Insert(BASE_LEN, item)
	if original.Size != expected.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size, original.Size)
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestInsertToMiddleOfPopulatedLinkedList(t *testing.T) {
	original := createTestLinkedListOfLen(BASE_LEN)
	originalContents := original.Contents()

	item := getTestItem()
	const idx = BASE_LEN / 2

	front := originalContents[:idx]
	rear := originalContents[idx:]

	expectedContents := make([]int64, len(front))
	copy(expectedContents, front)
	expectedContents = append(expectedContents, item)
	expectedContents = append(expectedContents, rear...)
	expected := createTestLinkedListOfContent(expectedContents...)

	original.Insert(idx, item)
	if original.Size != expected.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size, original.Size)
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestInsertToEmptyLinkedList(t *testing.T) {
	item := getTestItem()
	original := createTestLinkedListOfLen(0)
	expected := createTestLinkedListOfContent(item)

	original.Insert(0, item)
	if original.Size != expected.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size, original.Size)
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestInsertToInvalidIndexLinkedList(t *testing.T) {
	item := getTestItem()
	original := createTestLinkedListOfLen(BASE_LEN)

	var indexError *list.InvalidIndexError
	err := original.Insert(BASE_LEN+1, item)
	if !errors.As(err, &indexError) {
		t.Error("Should have got invalid index error")
	}
}

func TestPrependToPopulatedLinkedList(t *testing.T) {
	item := getTestItem()

	original := createTestLinkedListOfLen(BASE_LEN)

	expected := createTestLinkedListOfContent(append([]int64{item}, original.Contents()...)...)
	original.Prepend(item)
	if original.Size != expected.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size, original.Size)
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestPrependToEmptyLinkedList(t *testing.T) {
	item := getTestItem()
	original := createTestLinkedListOfLen(0)
	expected := createTestLinkedListOfContent(item)

	original.Prepend(item)
	if original.Size != expected.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size, original.Size)
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestAppendToPopulatedLinkedList(t *testing.T) {
	item := getTestItem()
	original := createTestLinkedListOfLen(BASE_LEN)

	expected := createTestLinkedListOfContent(append(original.Contents(), item)...)
	original.Append(item)
	if original.Size != expected.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size, original.Size)
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestAppendToEmptyLinkedList(t *testing.T) {
	item := getTestItem()
	original := createTestLinkedListOfLen(0)
	expected := createTestLinkedListOfContent(item)

	original.Append(item)
	if original.Size != expected.Size {
		t.Errorf("Wrong length: wanted %v, got %v", expected.Size, original.Size)
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestRemoveFromFrontOfPopulatedLinkedList(t *testing.T) {
	original := createTestLinkedListOfLen(BASE_LEN)
	expected := createTestLinkedListOfContent(original.Contents()[1:]...)

	expectedItem := original.Head()
	removedItem, _ := original.Remove(0)
	if removedItem != expectedItem {
		t.Errorf("Wrong item removed: wanted %v, got %v", expectedItem, removedItem)
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestRemoveFromRearOfPopulatedLinkedList(t *testing.T) {
	original := createTestLinkedListOfLen(BASE_LEN)
	originalContents := original.Contents()

	expected := createTestLinkedListOfContent(originalContents[:BASE_LEN-1]...)
	expectedItem := originalContents[BASE_LEN-1]
	removedItem, _ := original.Remove(BASE_LEN - 1)
	if removedItem != expectedItem {
		t.Errorf("Wrong length: wanted %v, got %v", expectedItem, removedItem)
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestRemoveFromMiddleOfPopulatedLinkedList(t *testing.T) {
	const idx = BASE_LEN / 2
	original := createTestLinkedListOfLen(BASE_LEN)

	originalContents := original.Contents()
	expectedContents := make([]int64, len(originalContents[:idx]))
	copy(expectedContents, originalContents[:idx])
	expectedContents = append(expectedContents, originalContents[idx+1:]...)
	expected := createTestLinkedListOfContent(expectedContents...)

	expectedItem := originalContents[idx]
	removedItem, _ := original.Remove(idx)
	if removedItem != expectedItem {
		t.Errorf("Wrong item removed: wanted %v, got %v", expectedItem, removedItem)
	}

	if !original.Equals(*expected) {
		t.Errorf("Wrong list contents: expected %v, got %v", expected.Contents(), original.Contents())
	}
}

func TestRemoveFromEmptyLinkedList(t *testing.T) {
	original := createTestLinkedListOfLen(0)
	_, err := original.Remove(0)
	var indexError *list.InvalidIndexError
	if !errors.As(err, &indexError) {
		t.Error("Should have got invalid index error")
	}
}

func TestGetFromPopulatedLinkedList(t *testing.T) {
	l := createTestLinkedListOfLen(BASE_LEN)
	for i, item := range l.Contents() {
		retrieved, _ := l.Get(i)
		if item != retrieved {
			t.Errorf("Didn't retrieve correct item: expected %v, got %v", item, retrieved)
		}
	}
}

func TestGetFromEmptyLinkedList(t *testing.T) {
	testList := createTestLinkedListOfLen(0)

	var indexError *list.InvalidIndexError
	_, err := testList.Get(0)
	if !errors.As(err, &indexError) {
		t.Errorf("Should have got invalid index error")
	}
}

func TestSetItemsInPopulatedLinkedList(t *testing.T) {
	l := createTestLinkedListOfLen(BASE_LEN)
	expected := createTestLinkedListOfLen(BASE_LEN)

	for i, item := range expected.Contents() {
		l.Set(i, item)
	}

	if !l.Equals(*expected) {
		t.Errorf("Didn't set item(s) correctly: expected %v, got %v", expected.Contents(), l.Contents())
	}

}

func TestSetItemInEmptyLinkedList(t *testing.T) {
	testList := createTestLinkedListOfLen(0)

	var indexError *list.InvalidIndexError
	err := testList.Set(0, 0)
	if !errors.As(err, &indexError) {
		t.Error("Should have got invalid index error")
	}
}
