package list_test

import (
	"testing"

	"github.com/brandencmw/go-data-structures.git/list"
)

func TestListContentsEqualForEqualLists(t *testing.T) {
	list1 := list.ArrayList[int64]{}
	list2 := list.ArrayList[int64]{}

	l1Equalsl2 := list1.ContentsEqualTo(list2)
	if !l1Equalsl2 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1, list2)
	}
	l2Equalsl1 := list2.ContentsEqualTo(list1)
	if !l2Equalsl1 {
		t.Errorf("Wrong output, l1 is %v and l2 is %v", list1, list2)
	}

	list1 = list.ArrayList[int64]{1, 2, 3, 4, 5}
	list2 = list.ArrayList[int64]{1, 2, 3, 4, 5}
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
	list1 := list.ArrayList[int64]{1, 2, 3, 4, 5}
	list2 := list.ArrayList[int64]{1, 2, 3, 4}

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
	list1 := list.ArrayList[int64]{1, 2, 3, 4, 5}
	list2 := list.ArrayList[int64]{1, 2, 3, 4, 6}

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
	list := list.ArrayList[int64]{1, 4, 5, 7, 5}
	initialLen := len(list)

	result := list.Insert(5, 0)
	if result != initialLen+1 {
		t.Errorf("Wrong length: wanted %v, got %v", initialLen+1, result)
	}

}

func TestInsertToRearOfPopulatedList(t *testing.T) {
	list := list.ArrayList[int64]{1, 4, 5, 7, 5}
	initialLen := len(list)

	result := list.Insert(5, initialLen)
	if result != initialLen+1 {
		t.Errorf("Wrong length: wanted %v, got %v", initialLen+1, result)
	}
}
