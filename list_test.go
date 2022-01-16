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

func TestSorted(t *testing.T) {
	testCases := []struct {
		Obj      *List
		expected *List
		reverse  bool
	}{
		{
			Obj:      NewList([]int{2, 3, 4}),
			expected: NewList([]int{4, 3, 2}),
			reverse:  true,
		},
		{
			Obj:      NewList([]int32{2, 3, 4}),
			expected: NewList([]int32{4, 3, 2}),
			reverse:  true,
		},
		{
			Obj:      NewList([]int64{2, 3, 4}),
			expected: NewList([]int64{4, 3, 2}),
			reverse:  true,
		},
		{
			Obj:      NewList([]float32{2, 3, 4}),
			expected: NewList([]float32{4, 3, 2}),
			reverse:  true,
		},
		{
			Obj:      NewList([]float64{4, 3, 2}),
			expected: NewList([]float64{2, 3, 4}),
			reverse:  false,
		},
		{
			Obj:      NewList([]string{"2", "3", "4"}),
			expected: NewList([]string{"4", "3", "2"}),
			reverse:  true,
		},
	}
	for _, tC := range testCases {
		got := tC.Obj.Sorted(tC.reverse)

		for i := 0; i < got.Len(); i++ {
			Gotitem := got.Get(i)
			Expecteditem := tC.expected.Get(i)
			if Gotitem != Expecteditem {
				t.Errorf("Error [TestSorted] Got: %v Expected: %v \n.", got, tC.expected)
			}
		}

	}
}

func TestMax(t *testing.T) {
	var vint int = 5
	var vint32 int32 = 5
	var vint64 int64 = 5
	var vfloat32 float32 = 5
	var vfloat64 float64 = 5

	testCases := []struct {
		Obj      List
		expected interface{}
	}{
		{
			Obj:      *NewList([]int{2, 3, 5}),
			expected: vint,
		},
		{
			Obj:      *NewList([]int32{2, 3, 5}),
			expected: vint32,
		},
		{
			Obj:      *NewList([]int64{2, 3, 5}),
			expected: vint64,
		},
		{
			Obj:      *NewList([]float32{2, 3, 5}),
			expected: vfloat32,
		},
		{
			Obj:      *NewList([]float64{2, 3, 5}),
			expected: vfloat64,
		},
		{
			Obj:      *NewList([]string{"Hello", "world"}),
			expected: "world",
		},
	}
	for _, tC := range testCases {
		got, err := tC.Obj.Max()
		if err != nil {
			t.Errorf("Error TextMax: %v", err)
		}

		if got != tC.expected {
			t.Errorf("[Failed TestMax] : Got: %v, Expected: %v.\n", got, tC.expected)
		}

	}
}

func TestMin(t *testing.T) {
	var vint int = 2
	var vint32 int32 = 2
	var vint64 int64 = 2
	var vfloat32 float32 = 2
	var vfloat64 float64 = 2

	testCases := []struct {
		Obj      List
		expected interface{}
	}{
		{
			Obj:      *NewList([]int{2, 3, 5}),
			expected: vint,
		},
		{
			Obj:      *NewList([]int32{2, 3, 5}),
			expected: vint32,
		},
		{
			Obj:      *NewList([]int64{2, 3, 5}),
			expected: vint64,
		},
		{
			Obj:      *NewList([]float32{2, 3, 5}),
			expected: vfloat32,
		},
		{
			Obj:      *NewList([]float64{2, 3, 5}),
			expected: vfloat64,
		},
		{
			Obj:      *NewList([]string{"Hello", "world"}),
			expected: "Hello",
		},
	}
	for _, tC := range testCases {
		got, err := tC.Obj.Min()
		if err != nil {
			t.Errorf("Error TextMin: %v", err)
		}

		if got != tC.expected {
			t.Errorf("[Failed TestMin] : Got: %v, Expected: %v.\n", got, tC.expected)
		}

	}
}

func TestExists(t *testing.T) {
	var vint int = 2
	var vint32 int32 = 2
	var vint64 int64 = 2
	var vfloat32 float32 = 2
	var vfloat64 float64 = 2

	testCases := []struct {
		Obj      List
		element  interface{}
		expected bool
	}{
		{
			Obj:      *NewList([]int{2, 3, 5}),
			element:  vint,
			expected: true,
		},
		{
			Obj:      *NewList([]int32{2, 3, 5}),
			element:  vint32,
			expected: true,
		},
		{
			Obj:      *NewList([]int64{2, 3, 5}),
			element:  vint64,
			expected: true,
		},
		{
			Obj:      *NewList([]float32{2, 3, 5}),
			element:  vfloat32,
			expected: true,
		},
		{
			Obj:      *NewList([]float64{2, 3, 5}),
			element:  vfloat64,
			expected: true,
		},
		{
			Obj:      *NewList([]string{"Hello", "world"}),
			element:  "Hello",
			expected: true,
		},
		{
			Obj:      *NewList([]string{"Hello", "world"}),
			element:  "hell",
			expected: false,
		},
	}
	for _, tC := range testCases {
		got := tC.Obj.Contains(tC.element)

		if got != tC.expected {
			t.Errorf("[Failed TestExists] : Got: %v, Expected: %v.\n", got, tC.expected)
		}

	}
}

func TestCombinations(t *testing.T) {

	testCases := []struct {
		Obj      List
		no       int
		expected List
	}{
		{
			Obj:      *NewList([]string{"a", "b", "c"}),
			no:       2,
			expected: *NewList([]string{"ab", "ac", "bc"}),
		},
		{
			Obj:      *NewList([]string{"a", "b", "c"}),
			no:       3,
			expected: *NewList([]string{"abc"}),
		},
	}
	for _, tC := range testCases {
		got, _ := tC.Obj.Combinations(tC.no, "")

		for i := 0; i < got.Len(); i++ {
			if got.Get(i) != tC.expected.Get(i) {
				t.Errorf("[Failed TestCombinations] : Got: %v, Expected: %v.\n", got, &tC.expected)
			}
		}

	}
}

func TestIsEqual(t *testing.T) {

	testCases := []struct {
		Obj      *List
		other    *List
		expected bool
	}{
		{
			Obj:      NewList([]int{2, 3}),
			other:    NewList([]int{2, 3}),
			expected: true,
		},
		{
			Obj:      NewList([]int32{2, 3}),
			other:    NewList([]int32{5, 3}),
			expected: false,
		},
		{
			Obj:      NewList([]int64{2, 3}),
			other:    NewList([]int64{2, 5}),
			expected: false,
		},
		{
			Obj:      NewList([]float32{2, 3}),
			other:    NewList([]float32{2, 3}),
			expected: true,
		},
		{
			Obj:      NewList([]float64{2, 3}),
			other:    NewList([]float64{2, 5}),
			expected: false,
		},
		{
			Obj:      NewList([]string{"Hello", "world"}),
			other:    NewList([]string{"Hello", "world", "gang"}),
			expected: false,
		},
	}
	for _, tC := range testCases {
		got := tC.Obj.IsEqual(tC.other)

		if got != tC.expected {
			t.Errorf("[Error TestIsEqual] : Got: %v, Expected: %v.\n", got, tC.expected)
		}
	}
}

func TestSet(t *testing.T) {

	testCases := []struct {
		Obj      *List
		expected *List
	}{
		{
			Obj:      NewList([]int{2, 3, 2, 3}),
			expected: NewList([]int{2, 3}),
		},
		{
			Obj:      NewList([]int32{2, 3, 5, 5}),
			expected: NewList([]int32{2, 3, 5}),
		},
		{
			Obj:      NewList([]int64{2, 3}),
			expected: NewList([]int64{2, 3}),
		},
		{
			Obj:      NewList([]float32{2, 3, 1, 2, 3, 1}),
			expected: NewList([]float32{2, 3, 1}),
		},
		{
			Obj:      NewList([]float64{2, 3, 3}),
			expected: NewList([]float64{2, 3}),
		},
		{
			Obj:      NewList([]string{"Hello", "world", "hello", "world"}),
			expected: NewList([]string{"Hello", "world", "hello"}),
		},
	}
	for _, tC := range testCases {
		got, _ := tC.Obj.Set()

		if !got.IsEqual(tC.expected) {
			t.Errorf("[Error TestSet] : Got: %v, Expected: %v.\n", got, tC.expected)
		}
	}
}
