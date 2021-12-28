package golist

import "testing"

var obj = NewList([]int{1, 2})

func TestNewList(t *testing.T) {
	NewList([]int{1, 2})
	NewList([]int32{1, 2})
	NewList([]int64{1, 2})
	NewList([]float32{1, 2})
	NewList([]float64{1, 2})
	NewList([]string{"hello"})
}

func TestLast(t *testing.T) {

	got, _ := obj.Last()
	expected := 2
	if got != expected {
		t.Errorf("Error [TestLast], Got: %v, Expected: %v.\n", got, expected)
	}
}

func TestAppend(t *testing.T) {
	obj.Append(3)
	got, _ := obj.Last()
	expected := 3
	if got != expected {
		t.Errorf("Error [TestAppend], Got: %v, Expected: %v.\n", got, expected)
	}
}

func TestIndex(t *testing.T) {
	got := obj.Index(2)
	expected := 1
	if got != expected {
		t.Errorf("Error [TestIndex], Got: %v, Expected: %v.\n", got, expected)
	}
}

func TestCount(t *testing.T) {
	got := obj.Count(2)
	expected := 1
	if got != expected {
		t.Errorf("Error [TestCount], Got: %v, Expected: %v.\n", got, expected)
	}
}

func TestClear(t *testing.T) {
	obj.Clear()
	got, err := obj.Last()
	if err == nil {
		t.Errorf("Error [TestClear], List not empty {%v} | got: %v\n", err, got)
	}
}

func TestSortInt(t *testing.T) {
	obj := NewList([]int{2, 1, 4})
	sorted := obj.Sort()
	sortObj := sorted.(*[]int)
	expected := []int{1, 2, 4}
	for i, v := range *sortObj {
		if v != expected[i] {
			t.Errorf("Error [TestSortInt], Got: %v, Expected: %v.\n", sortObj, expected)
		}
	}

}

func TestSortInt32(t *testing.T) {
	obj := NewList([]int32{2, 1, 4})
	sorted := obj.Sort()
	sortObj := sorted.(*[]int32)
	expected := []int32{1, 2, 4}
	for i, v := range *sortObj {
		if v != expected[i] {
			t.Errorf("Error [TestSortInt32], Got: %v, Expected: %v.\n", sortObj, expected)
		}
	}

}

func TestPop(t *testing.T) {
	obj := NewList([]int{2, 3, 4})
	expected := []int{2, 4}
	popped := obj.Pop(1)
	expected_popped := 3
	if popped != expected_popped {
		t.Errorf("Error [TestPop], Got: %v, Expected: %v.\n", popped, expected_popped)
	}
	got := obj.list.([]int)
	for i, v := range got {
		if v != expected[i] {
			t.Errorf("Error [TestPop], Got: %v, Expected: %v.\n", got, expected)
		}
	}

}
