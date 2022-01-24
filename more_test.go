package golist

import "testing"

func TestRemove(t *testing.T) {
	var vint int = 5
	var vint32 int32 = 5
	var vint64 int64 = 5
	var vfloat32 float32 = 5
	var vfloat64 float64 = 5
	var vstring string = "5"

	testCases := []struct {
		Obj      List
		expected List
		remove   interface{}
	}{
		{
			Obj:      *NewList([]int{2, vint, 3, 4}),
			expected: *NewList([]int{2, 3, 4}),
			remove:   vint,
		},
		{
			Obj:      *NewList([]int32{2, 3, vint32, 4}),
			expected: *NewList([]int32{2, 3, 4}),
			remove:   vint32,
		},
		{
			Obj:      *NewList([]int64{2, 3, 4, vint64}),
			expected: *NewList([]int64{2, 3, 4}),
			remove:   vint64,
		},
		{
			Obj:      *NewList([]float32{vfloat32, 2, 3, 4}),
			expected: *NewList([]float32{2, 3, 4}),
			remove:   vfloat32,
		},
		{
			Obj:      *NewList([]float64{vfloat64, 2, 3, 4}),
			expected: *NewList([]float64{2, 3, 4}),
			remove:   vfloat64,
		},
		{
			Obj:      *NewList([]string{vstring, "2", "3", "4"}),
			expected: *NewList([]string{"2", "3", "4"}),
			remove:   vstring,
		},
	}
	for _, tC := range testCases {
		err := tC.Obj.Remove(tC.remove)
		if err != nil {
			t.Errorf("[Error Removing] : %v", err)
		}
		got := tC.Obj
		if !got.IsEqual(&tC.expected) {
			t.Errorf("[Error Inserting | not equal] : Got: %v, Expected: %v.\n", &got, &tC.expected)
		}

	}
}

func TestCopy(t *testing.T) {
	testCases := []struct {
		Obj      List
		expected List
	}{
		{
			Obj:      *NewList([]int{2, 3, 4}),
			expected: *NewList([]int{2, 3, 4}),
		},
		{
			Obj:      *NewList([]int32{2, 3, 4}),
			expected: *NewList([]int32{2, 3, 4}),
		},
		{
			Obj:      *NewList([]int64{2, 3, 4}),
			expected: *NewList([]int64{2, 3, 4}),
		},
		{
			Obj:      *NewList([]float32{2, 3, 4}),
			expected: *NewList([]float32{2, 3, 4}),
		},
		{
			Obj:      *NewList([]float64{2, 3, 4}),
			expected: *NewList([]float64{2, 3, 4}),
		},
		{
			Obj:      *NewList([]string{"2", "3", "4"}),
			expected: *NewList([]string{"2", "3", "4"}),
		},
	}
	for _, tC := range testCases {
		got, err := tC.Obj.Copy()
		if err != nil {
			t.Errorf("Error [TestCopy] %v.\n", err)
		}
		for i := 0; i < got.Len(); i++ {
			Gotitem := got.Get(i)
			Expecteditem := tC.expected.Get(i)
			if Gotitem != Expecteditem {
				t.Errorf("Error [TestCopy] Got: %v Expected: %v \n.", got, tC.expected)
			}
		}

	}
}

func TestSlice(t *testing.T) {
	testCases := []struct {
		Obj      List
		expected List
		start    int
		stop     int
	}{
		{
			Obj:      *NewList([]int{2, 3, 4}),
			expected: *NewList([]int{2, 3}),
			start:    0,
			stop:     2,
		},
		{
			Obj:      *NewList([]int32{2, 3, 4}),
			expected: *NewList([]int32{2}),
			start:    0,
			stop:     1,
		},
		{
			Obj:      *NewList([]int64{2, 3, 4}),
			expected: *NewList([]int64{2, 3, 4}),
			start:    0,
			stop:     -1,
		},
		{
			Obj:      *NewList([]float32{2, 3, 4}),
			expected: *NewList([]float32{}),
			start:    0,
			stop:     0,
		},
		{
			Obj:      *NewList([]float64{2, 3, 4}),
			expected: *NewList([]float64{2, 3, 4}),
			start:    0,
			stop:     -1,
		},
		{
			Obj:      *NewList([]string{"2", "3", "4"}),
			expected: *NewList([]string{"2", "3", "4"}),
			start:    0,
			stop:     -1,
		},
	}
	for _, tC := range testCases {
		got, err := tC.Obj.Slice(tC.start, tC.stop)
		if err != nil {
			t.Errorf("Error [TestSlice] %v.\n", err)
		}
		for i := 0; i < got.Len(); i++ {
			Gotitem := got.Get(i)
			Expecteditem := tC.expected.Get(i)
			if Gotitem != Expecteditem {
				t.Errorf("Error [TestSlice] Got: %v Expected: %v \n.", got, tC.expected)
			}
		}

	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		Obj      List
		expected List
	}{
		{
			Obj:      *NewList([]int{2, 3, 4}),
			expected: *NewList([]int{4, 3, 2}),
		},
		{
			Obj:      *NewList([]int32{2, 3, 4}),
			expected: *NewList([]int32{4, 3, 2}),
		},
		{
			Obj:      *NewList([]int64{2, 3, 4}),
			expected: *NewList([]int64{4, 3, 2}),
		},
		{
			Obj:      *NewList([]float32{2, 3, 4}),
			expected: *NewList([]float32{4, 3, 2}),
		},
		{
			Obj:      *NewList([]float64{2, 3, 4}),
			expected: *NewList([]float64{4, 3, 2}),
		},
		{
			Obj:      *NewList([]string{"2", "3", "4"}),
			expected: *NewList([]string{"4", "3", "2"}),
		},
	}
	for _, tC := range testCases {
		got := tC.Obj.Reverse()

		if !got.IsEqual(&tC.expected) {
			t.Errorf("Error [TestReverse] Got: %v Expected: %v \n.", got, &tC.expected)
		}

	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		Obj      List
		expected string
	}{
		{
			Obj:      *NewList([]int{2, 3, 4}),
			expected: "[2, 3, 4]",
		},
		{
			Obj:      *NewList([]int32{2, 3, 4}),
			expected: "[2, 3, 4]",
		},
		{
			Obj:      *NewList([]int64{2, 3, 4}),
			expected: "[2, 3, 4]",
		},
		{
			Obj:      *NewList([]float32{2, 3, 4}),
			expected: "[2, 3, 4]",
		},
		{
			Obj:      *NewList([]float64{2, 3, 4}),
			expected: "[2, 3, 4]",
		},
		{
			Obj:      *NewList([]string{"2", "3", "4"}),
			expected: "[\"2\", \"3\", \"4\"]",
		},
	}
	for _, tC := range testCases {
		got := tC.Obj.String()

		if got != tC.expected {
			t.Errorf("Error [TestString] Got: %v Expected: %v \n.", got, tC.expected)
		}

	}
}

func TestSum(t *testing.T) {
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
			Obj:      *NewList([]int{2, 3}),
			expected: vint,
		},
		{
			Obj:      *NewList([]int32{2, 3}),
			expected: vint32,
		},
		{
			Obj:      *NewList([]int64{2, 3}),
			expected: vint64,
		},
		{
			Obj:      *NewList([]float32{2, 3}),
			expected: vfloat32,
		},
		{
			Obj:      *NewList([]float64{2, 3}),
			expected: vfloat64,
		},
		{
			Obj:      *NewList([]string{"Hello", "world"}),
			expected: "Hello world",
		},
	}
	for _, tC := range testCases {
		got := tC.Obj.Sum()

		if got != tC.expected {
			t.Errorf("[Error TestSum | not equal] : Got: %v, Expected: %v.\n", got, tC.expected)
		}

	}
}

func TestList(t *testing.T) {

	testCases := []struct {
		Obj      List
		expected interface{}
	}{
		{
			Obj:      *NewList([]int{2, 3}),
			expected: []int{2, 3},
		},
		{
			Obj:      *NewList([]int32{2, 3}),
			expected: []int32{2, 3},
		},
		{
			Obj:      *NewList([]int64{2, 3}),
			expected: []int64{2, 3},
		},
		{
			Obj:      *NewList([]float32{2, 3}),
			expected: []float32{2, 3},
		},
		{
			Obj:      *NewList([]float64{2, 3}),
			expected: []float64{2, 3},
		},
		{
			Obj:      *NewList([]string{"Hello", "world"}),
			expected: []string{"Hello", "world"},
		},
	}
	for _, tC := range testCases {
		got := tC.Obj.List()
		gotObj := NewList(got)
		expected := NewList(tC.expected)
		for i := 0; i < expected.Len(); i++ {
			if gotObj.Get(i) != expected.Get(i) {
				t.Errorf("[Error TestList | not equal] : Got: %v, Expected: %v.\n", got, tC.expected)
			}
		}

	}
}

func TestJoin(t *testing.T) {
	testCases := []struct {
		Obj      *List
		expected string
		joiner   string
	}{
		{
			Obj:      NewList([]string{"Hello", "world"}),
			joiner:   ",",
			expected: "Hello,world",
		},
		{
			Obj:      NewList([]string{"Hello", "world"}),
			joiner:   " ",
			expected: "Hello world",
		},
		{
			Obj:      NewList([]string{"Hello", "world"}),
			joiner:   "-",
			expected: "Hello-world",
		},
		{
			Obj:      NewList([]string{"Hello", "world"}),
			joiner:   "",
			expected: "Helloworld",
		},
	}
	for _, tC := range testCases {
		got := tC.Obj.Join(tC.joiner)
		if got != tC.expected {
			t.Errorf("[Error TestJoin] : Got: %v, Expected: %v.\n", got, tC.expected)
		}
	}
}

func TestReplace(t *testing.T) {
	var vint int = 5
	var vint32 int32 = 5
	var vint64 int64 = 5
	var vfloat32 float32 = 5
	var vfloat64 float64 = 5

	testCases := []struct {
		Obj      *List
		expected *List
		replace  interface{}
		index    int
	}{
		{
			Obj:      NewList([]int{2, 3}),
			expected: NewList([]int{vint, 3}),
			replace:  vint,
			index:    0,
		},
		{
			Obj:      NewList([]int32{2, 3}),
			expected: NewList([]int32{5, 3}),
			replace:  vint32,
			index:    0,
		},
		{
			Obj:      NewList([]int64{2, 3}),
			expected: NewList([]int64{2, 5}),
			replace:  vint64,
			index:    1,
		},
		{
			Obj:      NewList([]float32{2, 3}),
			expected: NewList([]float32{2, 5}),
			replace:  vfloat32,
			index:    -1,
		},
		{
			Obj:      NewList([]float64{2, 3}),
			expected: NewList([]float64{2, 5}),
			replace:  vfloat64,
			index:    1,
		},
		{
			Obj:      NewList([]string{"Hello", "world"}),
			expected: NewList([]string{"Hello", "golang"}),
			replace:  "golang",
			index:    -1,
		},
	}
	for _, tC := range testCases {
		err := tC.Obj.Replace(tC.replace, tC.index)
		if err != nil {
			t.Errorf("[Error TestReplace] : %v\n", err)
		}
		got := tC.Obj

		for i := 0; i < got.Len(); i++ {
			if got.Get(i) != tC.expected.Get(i) {
				t.Errorf("[Error TestReplace] : Got: %v, Expected: %v.\n", got, tC.expected)
			}
		}

	}
}

func TestGCF(t *testing.T) {
	var vint int = 5
	var vint32 int32 = 5
	var vint64 int64 = 5
	var vfloat32 float32 = 5
	var vfloat64 float64 = 5
	var edge float64 = 0.3
	var edge2 float64 = 0

	testCases := []struct {
		Obj      *List
		expected interface{}
	}{
		{
			Obj:      NewList([]int{10, 5, 25, 200}),
			expected: vint,
		},
		// edge case
		{
			Obj:      NewList([]int{10, 5, 0, 0}),
			expected: vint,
		},
		{
			Obj:      NewList([]int32{10, 5, 25, 200}),
			expected: vint32,
		},
		{
			Obj:      NewList([]int64{10, 5, 25, 200}),
			expected: vint64,
		},
		{
			Obj:      NewList([]float32{10, 5, 25, 200}),
			expected: vfloat32,
		},
		{
			Obj:      NewList([]float64{10, 5, 25, 200}),
			expected: vfloat64,
		},
		// edge case
		{
			Obj:      NewList([]float64{6.3, 12}),
			expected: edge,
		},
		// edge case
		{
			Obj:      NewList([]float64{0, 0}),
			expected: edge2,
		},
		{
			Obj:      NewList([]string{"Hello", "world"}),
			expected: nil,
		},
	}
	for _, tC := range testCases {
		got, _ := tC.Obj.GCF()
		if got != tC.expected {
			t.Errorf("GCF Error : %v != %v", tC.expected, got)
		}

	}
}

func TestLCM(t *testing.T) {
	var vint int = 200
	var vint32 int32 = 200
	var vint64 int64 = 200
	var vfloat32 float32 = 200
	var vfloat64 float64 = 200
	var edge float64 = 252
	var edge1 int = 10
	var edge2 float64 = 0

	testCases := []struct {
		Obj      *List
		expected interface{}
	}{
		{
			Obj:      NewList([]int{10, 5, 25, 200}),
			expected: vint,
		},
		// edge case
		{
			Obj:      NewList([]int{10, 5, 0, 0}),
			expected: edge1,
		},
		{
			Obj:      NewList([]int32{10, 5, 25, 200}),
			expected: vint32,
		},
		{
			Obj:      NewList([]int64{10, 5, 25, 200}),
			expected: vint64,
		},
		{
			Obj:      NewList([]float32{10, 5, 25, 200}),
			expected: vfloat32,
		},
		{
			Obj:      NewList([]float64{10, 5, 25, 200}),
			expected: vfloat64,
		},
		// edge case
		{
			Obj:      NewList([]float64{6.3, 12}),
			expected: edge,
		},
		// edge case
		{
			Obj:      NewList([]float64{0, 0}),
			expected: edge2,
		},
		{
			Obj:      NewList([]string{"Hello", "world"}),
			expected: nil,
		},
	}
	for _, tC := range testCases {
		got, _ := tC.Obj.LCM()
		if got != tC.expected {
			t.Errorf("LCM Error : %v != %v", tC.expected, got)
		}

	}
}

func TestType(t *testing.T) {

	testCases := []struct {
		Obj      *List
		expected string
	}{
		{
			Obj:      NewList([]int{10, 5, 25, 200}),
			expected: "golist.List[]int",
		},
		{
			Obj:      NewList([]int32{10, 5, 25, 200}),
			expected: "golist.List[]int32",
		},
		{
			Obj:      NewList([]int64{10, 5, 25, 200}),
			expected: "golist.List[]int64",
		},
		{
			Obj:      NewList([]float32{10, 5, 25, 200}),
			expected: "golist.List[]float32",
		},
		{
			Obj:      NewList([]float64{10, 5, 25, 200}),
			expected: "golist.List[]float64",
		},
		{
			Obj:      NewList([]string{"Hello", "world"}),
			expected: "golist.List[]string",
		},
	}
	for _, tC := range testCases {
		got := tC.Obj.Type()
		if got != tC.expected {
			t.Errorf("Type Error : %v != %v", tC.expected, got)
		}

	}
}

func TestListSum(t *testing.T) {
	testCases := []struct {
		Obj      *List
		other    *List
		Error    error
		expected *List
	}{
		{
			Obj:      NewList([]int{2, 3, 4}),
			other:    NewList([]int{1, 1, 1}),
			Error:    nil,
			expected: NewList([]int{3, 4, 5}),
		},
		{
			Obj:      NewList([]int32{2, 3, 4}),
			other:    NewList([]int32{1, 1, 1}),
			Error:    nil,
			expected: NewList([]int32{3, 4, 5}),
		},
		{
			Obj:      NewList([]int64{2, 3, 4}),
			other:    NewList([]int64{1, 1, 1}),
			Error:    nil,
			expected: NewList([]int64{3, 4, 5}),
		},
		{
			Obj:      NewList([]float32{2, 3, 4}),
			other:    NewList([]float32{1, 1, 1}),
			Error:    nil,
			expected: NewList([]float32{3, 4, 5}),
		},
		{
			Obj:      NewList([]float64{2, 3, 4}),
			other:    NewList([]float64{1, 1, 1}),
			Error:    nil,
			expected: NewList([]float64{3, 4, 5}),
		},
		// edge cases
		{
			Obj:      NewList([]float64{2, 3, 4}),
			other:    NewList([]float64{1, 1, 1, 2}),
			Error:    ErrListsNotOfSameLen,
			expected: &List{},
		},
	}
	for _, tC := range testCases {
		got, err := tC.Obj.ListSum(tC.other)
		if err != tC.Error {
			t.Errorf("Errors Not same ListSum: Got: %v Expected: %v \n", err, tC.Error)
		}

		if !got.IsEqual(tC.expected) {
			t.Errorf("Error [TestListSum] Got: %v Expected: %v \n.", got, tC.expected)
		}

	}
}
