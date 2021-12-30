package golist

import (
	"testing"
)

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
	reverse := false
	sorted := obj.Sort(reverse)
	sortObj := sorted.([]int)
	expected := []int{1, 2, 4}
	for i, v := range sortObj {
		if v != expected[i] {
			t.Errorf("Error [TestSortInt], Got: %v, Expected: %v.\n", sortObj, expected)
		}
	}

}

func TestSortInt32(t *testing.T) {
	obj := NewList([]int32{2, 1, 4})
	reverse := false
	sorted := obj.Sort(reverse)
	sortObj := sorted.([]int32)
	expected := []int32{1, 2, 4}
	for i, v := range sortObj {
		if v != expected[i] {
			t.Errorf("Error [TestSortInt32 ], Got: %v, Expected: %v.\n", sortObj, expected)
		}
	}

}

func TestReverseSort(t *testing.T) {
	reverse := true
	testCases := []struct {
		Obj      List
		expected interface{}
	}{
		{
			Obj:      *NewList([]int{2, 3, 4}),
			expected: []int{4, 3, 2},
		},
		{
			Obj:      *NewList([]int32{2, 3, 4}),
			expected: []int32{4, 3, 2},
		},
		{
			Obj:      *NewList([]int64{2, 3, 4}),
			expected: []int64{4, 3, 2},
		},
		{
			Obj:      *NewList([]float32{2, 3, 4}),
			expected: []float32{4, 3, 2},
		},
		{
			Obj:      *NewList([]float64{2, 3, 4}),
			expected: []float64{4, 3, 2},
		},
		{
			Obj:      *NewList([]string{"2", "3", "4"}),
			expected: []string{"4", "3", "2"},
		},
	}
	for _, tC := range testCases {
		tC.Obj.Sort(reverse)
		got, err := tC.Obj.Last()
		if err != nil {
			t.Errorf("[Getting Last Error] %v", err)
		}
		expected, err := NewList(tC.expected).Last()
		if err != nil {
			t.Errorf("[Getting Last Error] %v", err)
		}
		if got != expected {
			t.Errorf("Error [TestReverseSort], Got: %v, Expected: %v | %v, %v.\n", got, expected, tC.Obj.list, tC.expected)
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

func TestLength(t *testing.T) {
	testCases := []struct {
		Obj      List
		expected int
	}{
		{
			Obj:      *NewList([]int{2, 3, 4}),
			expected: 3,
		},
		{
			Obj:      *NewList([]int32{2, 3, 4}),
			expected: 3,
		},
		{
			Obj:      *NewList([]int64{2, 3, 4}),
			expected: 3,
		},
		{
			Obj:      *NewList([]float32{2, 3, 4}),
			expected: 3,
		},
		{
			Obj:      *NewList([]float64{2, 3, 4}),
			expected: 3,
		},
		{
			Obj:      *NewList([]string{"2", "3", "4"}),
			expected: 3,
		},
	}
	for _, tC := range testCases {
		got := tC.Obj.Len()
		if got != tC.expected {
			t.Errorf("Error [TestLength], Got: %v, Expected: %v.\n", got, tC.expected)
		}

	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		Obj      List
		expected int
		index    int
	}{
		{
			Obj:      *NewList([]int{2, 3, 4}),
			expected: 2,
			index:    0,
		},
		{
			Obj:      *NewList([]int{2, 3, 4}),
			expected: 3,
			index:    1,
		},
		{
			Obj:      *NewList([]int{2, 3, 4}),
			expected: 4,
			index:    2,
		},
	}
	for _, tC := range testCases {
		gott := tC.Obj.Get(tC.index)
		got := gott.(int)
		if got != tC.expected {
			t.Errorf("Error [TestGet], Got: %v, Expected: %v.\n", got, tC.expected)
		}

	}
}

func TestInsert(t *testing.T) {
	var vint int = 5
	var vint32 int32 = 5
	var vint64 int64 = 5
	var vfloat32 float32 = 5
	var vfloat64 float64 = 5
	var vstring string = "5"

	testCases := []struct {
		Obj      List
		expected interface{}
		insert   interface{}
		index    int
	}{
		{
			Obj:      *NewList([]int{2, 3, 4}),
			expected: []int{2, vint, 3, 4},
			insert:   vint,
			index:    1,
		},
		{
			Obj:      *NewList([]int32{2, 3, 4}),
			expected: []int32{2, 3, vint32, 4},
			insert:   vint32,
			index:    2,
		},
		{
			Obj:      *NewList([]int64{2, 3, 4}),
			expected: []int64{2, 3, 4, vint64},
			insert:   vint64,
			index:    3,
		},
		{
			Obj:      *NewList([]float32{2, 3, 4}),
			expected: []float32{vfloat32, 2, 3, 4},
			insert:   vfloat32,
			index:    0,
		},
		{
			Obj:      *NewList([]float64{2, 3, 4}),
			expected: []float64{vfloat64, 2, 3, 4},
			insert:   vfloat64,
			index:    0,
		},
		{
			Obj:      *NewList([]string{"2", "3", "4"}),
			expected: []string{vstring, "2", "3", "4"},
			insert:   vstring,
			index:    0,
		},
	}
	for _, tC := range testCases {
		err := tC.Obj.Insert(tC.insert, tC.index)
		if err != nil {
			t.Errorf("[Error Inserting] : %v", err)
		}

		switch tC.Obj.list.(type) {
		case []int:
			list := tC.Obj.list.([]int)
			compare := tC.expected.([]int)
			for i, v := range list {
				if v != compare[i] {
					t.Errorf("[Error Inserting | not equal] : Got: %v, Expected: %v.\n", list, tC.expected)
				}
			}

		case []int32:
			list := tC.Obj.list.([]int32)
			compare := tC.expected.([]int32)
			for i, v := range list {
				if v != compare[i] {
					t.Errorf("[Error Inserting | not equal] : Got: %v, Expected: %v.\n", list, tC.expected)
				}
			}

		case []int64:
			list := tC.Obj.list.([]int64)
			compare := tC.expected.([]int64)
			for i, v := range list {
				if v != compare[i] {
					t.Errorf("[Test Failed | not equal] : Got: %v, Expected: %v.\n", list, tC.expected)
				}
			}

		case []float32:
			list := tC.Obj.list.([]float32)
			compare := tC.expected.([]float32)
			for i, v := range list {
				if v != compare[i] {
					t.Errorf("[Test Failed  | not equal] : Got: %v, Expected: %v.\n", list, tC.expected)
				}
			}

		case []float64:
			list := tC.Obj.list.([]float64)
			compare := tC.expected.([]float64)
			for i, v := range list {
				if v != compare[i] {
					t.Errorf("[Test Failed  | not equal] : Got: %v, Expected: %v.\n", list, tC.expected)
				}
			}

		case []string:
			list := tC.Obj.list.([]string)
			compare := tC.expected.([]string)
			for i, v := range list {
				if v != compare[i] {
					t.Errorf("[Error Inserting | not equal] : Got: %v, Expected: %v.\n", list, tC.expected)
				}
			}

		default:
			return
		}

	}
}
