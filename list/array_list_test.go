package list_test

import (
	"testing"

	"github.com/brandencmw/go-data-structures.git/list"
)

func compareListContents(expectedList, listToCheck list.ArrayList[int]) {

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
