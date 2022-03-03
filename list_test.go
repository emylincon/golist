package golist_test

import (
	"testing"

	"github.com/emylincon/golist"
)

func TestNewList(t *testing.T) {
	golist.NewList([]int{1, 2})
	golist.NewList([]int32{1, 2})
	golist.NewList([]int64{1, 2})
	golist.NewList([]float32{1, 2})
	golist.NewList([]float64{1, 2})
	golist.NewList([]string{"hello"})
}

func TestLast(t *testing.T) {
	obj := golist.NewList([]int{1, 2})
	got, _ := obj.Last()
	expected := 2
	if got != expected {
		t.Errorf("Error [TestLast], Got: %v, Expected: %v.\n", got, expected)
	}
}

func TestAppend(t *testing.T) {
	obj := golist.NewList([]int{1, 2})
	obj.Append(3)
	got, _ := obj.Last()
	expected := 3
	if got != expected {
		t.Errorf("Error [TestAppend], Got: %v, Expected: %v.\n", got, expected)
	}
}

func TestIndex(t *testing.T) {
	obj := golist.NewList([]int{1, 2})
	got := obj.Index(2)
	expected := 1
	if got != expected {
		t.Errorf("Error [TestIndex], Got: %v, Expected: %v.\n", got, expected)
	}
}

func TestCount(t *testing.T) {
	obj := golist.NewList([]int{1, 2})
	got := obj.Count(2)
	expected := 1
	if got != expected {
		t.Errorf("Error [TestCount], Got: %v, Expected: %v.\n", got, expected)
	}
}

func TestClear(t *testing.T) {
	obj := golist.NewList([]int{1, 2})
	obj.Clear()
	got, err := obj.Last()
	if err == nil {
		t.Errorf("Error [TestClear], List not empty {%v} | got: %v\n", err, got)
	}
}

func TestSortInt(t *testing.T) {
	obj := golist.NewList([]int{2, 1, 4})
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
	obj := golist.NewList([]int32{2, 1, 4})
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
		Obj      golist.List
		expected golist.List
	}{
		{
			Obj:      *golist.NewList([]int{2, 3, 4}),
			expected: *golist.NewList([]int{4, 3, 2}),
		},
		{
			Obj:      *golist.NewList([]int32{2, 3, 4}),
			expected: *golist.NewList([]int32{4, 3, 2}),
		},
		{
			Obj:      *golist.NewList([]int64{2, 3, 4}),
			expected: *golist.NewList([]int64{4, 3, 2}),
		},
		{
			Obj:      *golist.NewList([]float32{2, 3, 4}),
			expected: *golist.NewList([]float32{4, 3, 2}),
		},
		{
			Obj:      *golist.NewList([]float64{2, 3, 4}),
			expected: *golist.NewList([]float64{4, 3, 2}),
		},
		{
			Obj:      *golist.NewList([]string{"2", "3", "4"}),
			expected: *golist.NewList([]string{"4", "3", "2"}),
		},
	}
	for _, tC := range testCases {
		tC.Obj.Sort(reverse)
		got := tC.Obj
		if !got.IsEqual(&tC.expected) {
			t.Errorf("Error [TestReverseSort], Got: %v, Expected: %v.\n", &got, &tC.expected)
		}

	}
}

func TestPop(t *testing.T) {
	obj := golist.NewList([]int{2, 3, 4})
	expected := []int{2, 4}
	popped := obj.Pop(1)
	expectedPopped := 3
	if popped != expectedPopped {
		t.Errorf("Error [TestPop], Got: %v, Expected: %v.\n", popped, expectedPopped)
	}
	got := obj.List().([]int)
	for i, v := range got {
		if v != expected[i] {
			t.Errorf("Error [TestPop], Got: %v, Expected: %v.\n", got, expected)
		}
	}

}

func TestLength(t *testing.T) {
	testCases := []struct {
		Obj      golist.List
		expected int
	}{
		{
			Obj:      *golist.NewList([]int{2, 3, 4}),
			expected: 3,
		},
		{
			Obj:      *golist.NewList([]int32{2, 3, 4}),
			expected: 3,
		},
		{
			Obj:      *golist.NewList([]int64{2, 3, 4}),
			expected: 3,
		},
		{
			Obj:      *golist.NewList([]float32{2, 3, 4}),
			expected: 3,
		},
		{
			Obj:      *golist.NewList([]float64{2, 3, 4}),
			expected: 3,
		},
		{
			Obj:      *golist.NewList([]string{"2", "3", "4"}),
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
		Obj      golist.List
		expected int
		index    int
	}{
		{
			Obj:      *golist.NewList([]int{2, 3, 4}),
			expected: 2,
			index:    0,
		},
		{
			Obj:      *golist.NewList([]int{2, 3, 4}),
			expected: 3,
			index:    1,
		},
		{
			Obj:      *golist.NewList([]int{2, 3, 4}),
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
		Obj      golist.List
		expected golist.List
		insert   interface{}
		index    int
	}{
		{
			Obj:      *golist.NewList([]int{2, 3, 4}),
			expected: *golist.NewList([]int{2, vint, 3, 4}),
			insert:   vint,
			index:    1,
		},
		{
			Obj:      *golist.NewList([]int32{2, 3, 4}),
			expected: *golist.NewList([]int32{2, 3, vint32, 4}),
			insert:   vint32,
			index:    2,
		},
		{
			Obj:      *golist.NewList([]int64{2, 3, 4}),
			expected: *golist.NewList([]int64{2, 3, 4, vint64}),
			insert:   vint64,
			index:    3,
		},
		{
			Obj:      *golist.NewList([]float32{2, 3, 4}),
			expected: *golist.NewList([]float32{vfloat32, 2, 3, 4}),
			insert:   vfloat32,
			index:    0,
		},
		{
			Obj:      *golist.NewList([]float64{2, 3, 4}),
			expected: *golist.NewList([]float64{vfloat64, 2, 3, 4}),
			insert:   vfloat64,
			index:    0,
		},
		{
			Obj:      *golist.NewList([]string{"2", "3", "4"}),
			expected: *golist.NewList([]string{vstring, "2", "3", "4"}),
			insert:   vstring,
			index:    0,
		},
	}
	for _, tC := range testCases {
		err := tC.Obj.Insert(tC.insert, tC.index)
		if err != nil {
			t.Errorf("[Error Inserting] : %v", err)
		}
		got := tC.Obj
		if !got.IsEqual(&tC.expected) {
			t.Errorf("[Error Inserting | not equal] : Got: %v, Expected: %v.\n", &got, &tC.expected)
		}

	}
}

func TestSorted(t *testing.T) {
	testCases := []struct {
		Obj      *golist.List
		expected *golist.List
		reverse  bool
	}{
		{
			Obj:      golist.NewList([]int{2, 3, 4}),
			expected: golist.NewList([]int{4, 3, 2}),
			reverse:  true,
		},
		{
			Obj:      golist.NewList([]int32{2, 3, 4}),
			expected: golist.NewList([]int32{4, 3, 2}),
			reverse:  true,
		},
		{
			Obj:      golist.NewList([]int64{2, 3, 4}),
			expected: golist.NewList([]int64{4, 3, 2}),
			reverse:  true,
		},
		{
			Obj:      golist.NewList([]float32{2, 3, 4}),
			expected: golist.NewList([]float32{4, 3, 2}),
			reverse:  true,
		},
		{
			Obj:      golist.NewList([]float64{4, 3, 2}),
			expected: golist.NewList([]float64{2, 3, 4}),
			reverse:  false,
		},
		{
			Obj:      golist.NewList([]string{"2", "3", "4"}),
			expected: golist.NewList([]string{"4", "3", "2"}),
			reverse:  true,
		},
	}
	for _, tC := range testCases {
		got := tC.Obj.Sorted(tC.reverse)

		if !got.IsEqual(tC.expected) {
			t.Errorf("Error [TestSorted] Got: %v Expected: %v \n.", got, tC.expected)
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
		Obj      golist.List
		expected interface{}
	}{
		{
			Obj:      *golist.NewList([]int{2, 3, 5}),
			expected: vint,
		},
		{
			Obj:      *golist.NewList([]int32{2, 3, 5}),
			expected: vint32,
		},
		{
			Obj:      *golist.NewList([]int64{2, 3, 5}),
			expected: vint64,
		},
		{
			Obj:      *golist.NewList([]float32{2, 3, 5}),
			expected: vfloat32,
		},
		{
			Obj:      *golist.NewList([]float64{2, 3, 5}),
			expected: vfloat64,
		},
		{
			Obj:      *golist.NewList([]string{"Hello", "world"}),
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
		Obj      golist.List
		expected interface{}
	}{
		{
			Obj:      *golist.NewList([]int{2, 3, 5}),
			expected: vint,
		},
		{
			Obj:      *golist.NewList([]int32{2, 3, 5}),
			expected: vint32,
		},
		{
			Obj:      *golist.NewList([]int64{2, 3, 5}),
			expected: vint64,
		},
		{
			Obj:      *golist.NewList([]float32{2, 3, 5}),
			expected: vfloat32,
		},
		{
			Obj:      *golist.NewList([]float64{2, 3, 5}),
			expected: vfloat64,
		},
		{
			Obj:      *golist.NewList([]string{"Hello", "world"}),
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
		Obj      golist.List
		element  interface{}
		expected bool
	}{
		{
			Obj:      *golist.NewList([]int{2, 3, 5}),
			element:  vint,
			expected: true,
		},
		{
			Obj:      *golist.NewList([]int32{2, 3, 5}),
			element:  vint32,
			expected: true,
		},
		{
			Obj:      *golist.NewList([]int64{2, 3, 5}),
			element:  vint64,
			expected: true,
		},
		{
			Obj:      *golist.NewList([]float32{2, 3, 5}),
			element:  vfloat32,
			expected: true,
		},
		{
			Obj:      *golist.NewList([]float64{2, 3, 5}),
			element:  vfloat64,
			expected: true,
		},
		{
			Obj:      *golist.NewList([]string{"Hello", "world"}),
			element:  "Hello",
			expected: true,
		},
		{
			Obj:      *golist.NewList([]string{"Hello", "world"}),
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
		Obj      golist.List
		no       int
		expected golist.List
	}{
		{
			Obj:      *golist.NewList([]string{"a", "b", "c"}),
			no:       2,
			expected: *golist.NewList([]string{"ab", "ac", "bc"}),
		},
		{
			Obj:      *golist.NewList([]string{"a", "b", "c"}),
			no:       3,
			expected: *golist.NewList([]string{"abc"}),
		},
	}
	for _, tC := range testCases {
		got, _ := tC.Obj.Combinations(tC.no, "")

		if !got.IsEqual(&tC.expected) {
			t.Errorf("[Failed TestCombinations] : Got: %v, Expected: %v.\n", got, &tC.expected)
		}

	}
}

func TestCombinationsMax(t *testing.T) {

	testCases := []struct {
		Obj      golist.List
		no       int
		expected golist.List
	}{
		{
			Obj:      *golist.NewList([]string{"a", "b", "c"}),
			no:       2,
			expected: *golist.NewList([]string{"a", "b", "c", "ab", "ac", "bc"}),
		},
		{
			Obj:      *golist.NewList([]string{"a", "b", "c"}),
			no:       3,
			expected: *golist.NewList([]string{"a", "b", "c", "ab", "ac", "bc", "abc"}),
		},
	}
	for _, tC := range testCases {
		got, _ := tC.Obj.CombinationsMax(tC.no, "")

		if !got.Sorted(false).IsEqual(tC.expected.Sorted(false)) {
			t.Errorf("[Failed TestCombinationsMax] : Got: %v, Expected: %v.\n", got, &tC.expected)
		}

	}
}

func TestIsEqual(t *testing.T) {

	testCases := []struct {
		Obj      *golist.List
		other    *golist.List
		expected bool
	}{
		{
			Obj:      golist.NewList([]int{2, 3}),
			other:    golist.NewList([]int{2, 3}),
			expected: true,
		},
		{
			Obj:      golist.NewList([]int32{2, 3}),
			other:    golist.NewList([]int32{5, 3}),
			expected: false,
		},
		{
			Obj:      golist.NewList([]int64{2, 3}),
			other:    golist.NewList([]int64{2, 5}),
			expected: false,
		},
		{
			Obj:      golist.NewList([]float32{2, 3}),
			other:    golist.NewList([]float32{2, 3}),
			expected: true,
		},
		{
			Obj:      golist.NewList([]float64{2, 3}),
			other:    golist.NewList([]float64{2, 5}),
			expected: false,
		},
		{
			Obj:      golist.NewList([]string{"Hello", "world"}),
			other:    golist.NewList([]string{"Hello", "world", "gang"}),
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
		Obj      *golist.List
		expected *golist.List
	}{
		{
			Obj:      golist.NewList([]int{2, 3, 2, 3}),
			expected: golist.NewList([]int{2, 3}),
		},
		{
			Obj:      golist.NewList([]int32{2, 3, 5, 5}),
			expected: golist.NewList([]int32{2, 3, 5}),
		},
		{
			Obj:      golist.NewList([]int64{2, 3}),
			expected: golist.NewList([]int64{2, 3}),
		},
		{
			Obj:      golist.NewList([]float32{2, 3, 1, 2, 3, 1}),
			expected: golist.NewList([]float32{2, 3, 1}),
		},
		{
			Obj:      golist.NewList([]float64{2, 3, 3}),
			expected: golist.NewList([]float64{2, 3}),
		},
		{
			Obj:      golist.NewList([]string{"Hello", "world", "hello", "world"}),
			expected: golist.NewList([]string{"Hello", "world", "hello"}),
		},
	}
	for _, tC := range testCases {
		got, _ := tC.Obj.Set()

		if !got.IsEqual(tC.expected) {
			t.Errorf("[Error TestSet] : Got: %v, Expected: %v.\n", got, tC.expected)
		}
	}
}

func TestExtend(t *testing.T) {

	testCases := []struct {
		Obj      *golist.List
		other    interface{}
		expected *golist.List
	}{
		{
			Obj:      golist.NewList([]int{2, 3}),
			other:    []int{4, 5},
			expected: golist.NewList([]int{2, 3, 4, 5}),
		},
		{
			Obj:      golist.NewList([]int32{2, 5}),
			other:    []int32{4, 5},
			expected: golist.NewList([]int32{2, 5, 4, 5}),
		},
		{
			Obj:      golist.NewList([]int64{2, 3}),
			other:    []int64{4, 5},
			expected: golist.NewList([]int64{2, 3, 4, 5}),
		},
		{
			Obj:      golist.NewList([]float32{2, 1, 2, 1}),
			other:    []float32{4, 5},
			expected: golist.NewList([]float32{2, 1, 2, 1, 4, 5}),
		},
		{
			Obj:      golist.NewList([]float64{2, 3, 3}),
			other:    []float64{1, 1},
			expected: golist.NewList([]float64{2, 3, 3, 1, 1}),
		},
		{
			Obj:      golist.NewList([]string{"Hello"}),
			other:    []string{"world"},
			expected: golist.NewList([]string{"Hello", "world"}),
		},
	}
	for _, tC := range testCases {
		err := tC.Obj.Extend(tC.other)
		if err != nil {
			t.Errorf("[Error TestExtend]: %v", err)
		}
		got := tC.Obj
		if !got.IsEqual(tC.expected) {
			t.Errorf("[Error TestExtend] : Got: %v, Expected: %v.\n", got, tC.expected)
		}
	}
}

func TestAdd(t *testing.T) {

	testCases := []struct {
		Obj      *golist.List
		other    *golist.List
		expected *golist.List
	}{
		{
			Obj:      golist.NewList([]int{2, 3}),
			other:    golist.NewList([]int{4, 5}),
			expected: golist.NewList([]int{2, 3, 4, 5}),
		},
		{
			Obj:      golist.NewList([]int32{2, 5}),
			other:    golist.NewList([]int32{4, 5}),
			expected: golist.NewList([]int32{2, 5, 4, 5}),
		},
		{
			Obj:      golist.NewList([]int64{2, 3}),
			other:    golist.NewList([]int64{4, 5}),
			expected: golist.NewList([]int64{2, 3, 4, 5}),
		},
		{
			Obj:      golist.NewList([]float32{2, 1, 2, 1}),
			other:    golist.NewList([]float32{4, 5}),
			expected: golist.NewList([]float32{2, 1, 2, 1, 4, 5}),
		},
		{
			Obj:      golist.NewList([]float64{2, 3, 3}),
			other:    golist.NewList([]float64{1, 1}),
			expected: golist.NewList([]float64{2, 3, 3, 1, 1}),
		},
		{
			Obj:      golist.NewList([]string{"Hello"}),
			other:    golist.NewList([]string{"world"}),
			expected: golist.NewList([]string{"Hello", "world"}),
		},
	}
	for _, tC := range testCases {
		got, err := tC.Obj.Add(tC.other)
		if err != nil {
			t.Errorf("[Error TestAdd]: %v", err)
		}

		if !got.IsEqual(tC.expected) {
			t.Errorf("[Error TestAdd] : Got: %v, Expected: %v.\n", got, tC.expected)
		}
	}
}

func TestListSumNo(t *testing.T) {
	var vint int = 5
	var vint32 int32 = 5
	var vint64 int64 = 5
	var vfloat32 float32 = 5
	var vfloat64 float64 = 5

	testCases := []struct {
		Obj      golist.List
		expected golist.List
		no       interface{}
	}{
		{
			Obj:      *golist.NewList([]int{2, 3, 4}),
			expected: *golist.NewList([]int{7, 8, 9}),
			no:       vint,
		},
		{
			Obj:      *golist.NewList([]int32{2, 3, 4}),
			expected: *golist.NewList([]int32{7, 8, 9}),
			no:       vint32,
		},
		{
			Obj:      *golist.NewList([]int64{2, 3, 4}),
			expected: *golist.NewList([]int64{7, 8, 9}),
			no:       vint64,
		},
		{
			Obj:      *golist.NewList([]float32{2, 3, 4}),
			expected: *golist.NewList([]float32{7, 8, 9}),
			no:       vfloat32,
		},
		{
			Obj:      *golist.NewList([]float64{2, 3, 4}),
			expected: *golist.NewList([]float64{7, 8, 9}),
			no:       vfloat64,
		},
	}
	for _, tC := range testCases {
		got, err := tC.Obj.ListSumNo(tC.no)
		if err != nil {
			t.Errorf("[Error SumListNo] : %v", err)
		}
		if !got.IsEqual(&tC.expected) {
			t.Errorf("[Error SumListNo] : Got: %v, Expected: %v.\n", got, &tC.expected)
		}

	}
}

func TestListSubtract(t *testing.T) {

	testCases := []struct {
		Obj      *golist.List
		other    *golist.List
		expected *golist.List
	}{
		{
			Obj:      golist.NewList([]int{2, 3}),
			other:    golist.NewList([]int{4, 5}),
			expected: golist.NewList([]int{-2, -2}),
		},
		{
			Obj:      golist.NewList([]int32{2, 5}),
			other:    golist.NewList([]int32{4, 5}),
			expected: golist.NewList([]int32{-2, 0}),
		},
		{
			Obj:      golist.NewList([]int64{2, 3}),
			other:    golist.NewList([]int64{4, 5}),
			expected: golist.NewList([]int64{-2, -2}),
		},
		{
			Obj:      golist.NewList([]float32{2, 3}),
			other:    golist.NewList([]float32{4, 5}),
			expected: golist.NewList([]float32{-2, -2}),
		},
		{
			Obj:      golist.NewList([]float64{2, 3}),
			other:    golist.NewList([]float64{1, 1}),
			expected: golist.NewList([]float64{1, 2}),
		},
	}
	for _, tC := range testCases {
		got, err := tC.Obj.ListSubtract(tC.other)
		if err != nil {
			t.Errorf("[Error TestListSubtract]: %v", err)
		}

		if !got.IsEqual(tC.expected) {
			t.Errorf("[Error TestListSubtract] : Got: %v, Expected: %v. type: %v\n", got, tC.expected, tC.expected.Type())
		}
	}
}

func TestListSubtractNo(t *testing.T) {
	var vint int = -5
	var vint32 int32 = -5
	var vint64 int64 = -5
	var vfloat32 float32 = -5
	var vfloat64 float64 = -5

	testCases := []struct {
		Obj      golist.List
		expected golist.List
		no       interface{}
	}{
		{
			Obj:      *golist.NewList([]int{2, 3, 4}),
			expected: *golist.NewList([]int{7, 8, 9}),
			no:       vint,
		},
		{
			Obj:      *golist.NewList([]int32{2, 3, 4}),
			expected: *golist.NewList([]int32{7, 8, 9}),
			no:       vint32,
		},
		{
			Obj:      *golist.NewList([]int64{2, 3, 4}),
			expected: *golist.NewList([]int64{7, 8, 9}),
			no:       vint64,
		},
		{
			Obj:      *golist.NewList([]float32{2, 3, 4}),
			expected: *golist.NewList([]float32{7, 8, 9}),
			no:       vfloat32,
		},
		{
			Obj:      *golist.NewList([]float64{2, 3, 4}),
			expected: *golist.NewList([]float64{7, 8, 9}),
			no:       vfloat64,
		},
	}
	for _, tC := range testCases {
		got, err := tC.Obj.ListSubtractNo(tC.no)
		if err != nil {
			t.Errorf("[Error SubtractListNo] : %v", err)
		}
		if !got.IsEqual(&tC.expected) {
			t.Errorf("[Error SubtractListNo] : Got: %v, Expected: %v.\n", got, &tC.expected)
		}

	}
}

func TestConvertToSlice(t *testing.T) {
	testCases := []struct {
		Obj      *golist.List
		expected *golist.List
		desc     string
	}{
		{
			Obj:      golist.NewList([]int{10, 5, 25, 200}),
			expected: golist.NewList([]float32{10, 5, 25, 200}),
			desc:     "convert to float32[]",
		},
		{
			Obj:      golist.NewList([]int32{10, 5, 25, 200}),
			expected: golist.NewList([]float64{10, 5, 25, 200}),
			desc:     "convert to Float64[]",
		},
		{
			Obj:      golist.NewList([]int64{10, 5, 25, 200}),
			expected: golist.NewList([]int{10, 5, 25, 200}),
			desc:     "convert to int[]",
		},
		{
			Obj:      golist.NewList([]float32{10, 5, 25, 200}),
			expected: golist.NewList([]int32{10, 5, 25, 200}),
			desc:     "convert to int32[]",
		},
		{
			Obj:      golist.NewList([]float64{10, 5, 25, 200}),
			expected: golist.NewList([]int64{10, 5, 25, 200}),
			desc:     "convert to int64[]",
		},
		{
			Obj:      golist.NewList([]float64{10, 5, 25, 200}),
			expected: golist.NewList([]string{"10", "5", "25", "200"}),
			desc:     "convert to string[]",
		},
		{
			Obj:      golist.NewList([]string{"10", "05", "25", "200"}),
			expected: golist.NewList([]float64{10, 5, 25, 200}),
			desc:     "convert to float64[]",
		},
		{
			Obj:      golist.NewList([]string{"Hello", "world"}),
			expected: golist.NewList([]string{"Hello", "world"}),
			desc:     "convert string[] to string[]",
		},
	}

	t.Run(testCases[0].desc, func(t *testing.T) {
		got, err := testCases[0].Obj.ConvertToSliceFloat32()
		if err != nil {
			t.Errorf("[Error ConvertToSlice](0) : %v", err)
		}
		gotList := golist.NewList(got)
		if !gotList.IsEqual(testCases[0].expected) {
			t.Errorf("Error ConvertToSlice](0) : Got: %v, Expected: %v.\n", got, testCases[1].expected)
		}
	})

	t.Run(testCases[1].desc, func(t *testing.T) {
		got, err := testCases[1].Obj.ConvertToSliceFloat64()
		if err != nil {
			t.Errorf("[Error ConvertToSlice](1) : %v", err)
		}
		gotList := golist.NewList(got)
		if !gotList.IsEqual(testCases[1].expected) {
			t.Errorf("Error ConvertToSlice](1) : Got: %v, Expected: %v.\n", got, testCases[1].expected)
		}
	})

	t.Run(testCases[2].desc, func(t *testing.T) {
		got, err := testCases[2].Obj.ConvertToSliceInt()
		if err != nil {
			t.Errorf("[Error ConvertToSlice](2) : %v", err)
		}
		gotList := golist.NewList(got)
		if !gotList.IsEqual(testCases[2].expected) {
			t.Errorf("Error ConvertToSlice](2) : Got: %v, Expected: %v.\n", got, testCases[2].expected)
		}
	})
}

func TestListDifference(t *testing.T) {

	testCases := []struct {
		Obj      *golist.List
		other    *golist.List
		expected *golist.List
	}{
		{
			Obj:      golist.NewList([]int{2, 3, 4}),
			other:    golist.NewList([]int{2, 1, 4}),
			expected: golist.NewList([]int{3}),
		},
		{
			Obj:      golist.NewList([]int32{2, 3, 4}),
			other:    golist.NewList([]int32{2, 3, 4}),
			expected: golist.NewList([]int32{}),
		},
		{
			Obj:      golist.NewList([]int64{2, 3, 4}),
			other:    golist.NewList([]int64{2, 3, 4}),
			expected: golist.NewList([]int64{}),
		},
		{
			Obj:      golist.NewList([]float32{2, 3, 4}),
			other:    golist.NewList([]float32{2, 2, 4}),
			expected: golist.NewList([]float32{3}),
		},
		{
			Obj:      golist.NewList([]float64{2, 3, 4}),
			other:    golist.NewList([]float64{1}),
			expected: golist.NewList([]float64{2, 3, 4}),
		},
	}
	for _, tC := range testCases {
		got, err := tC.Obj.Difference(tC.other)
		if err != nil {
			t.Errorf("[Error Difference] : %v", err)
		}
		if !got.IsEqual(tC.expected) {
			t.Errorf("[Error Difference] : Got: %v, Expected: %v.\n", got, tC.expected)
		}

	}
}
