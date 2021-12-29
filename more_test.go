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
		expected interface{}
		remove   interface{}
	}{
		{
			Obj:      *NewList([]int{2, vint, 3, 4}),
			expected: []int{2, 3, 4},
			remove:   vint,
		},
		{
			Obj:      *NewList([]int32{2, 3, vint32, 4}),
			expected: []int32{2, 3, 4},
			remove:   vint32,
		},
		{
			Obj:      *NewList([]int64{2, 3, 4, vint64}),
			expected: []int64{2, 3, 4},
			remove:   vint64,
		},
		{
			Obj:      *NewList([]float32{vfloat32, 2, 3, 4}),
			expected: []float32{2, 3, 4},
			remove:   vfloat32,
		},
		{
			Obj:      *NewList([]float64{vfloat64, 2, 3, 4}),
			expected: []float64{2, 3, 4},
			remove:   vfloat64,
		},
		{
			Obj:      *NewList([]string{vstring, "2", "3", "4"}),
			expected: []string{"2", "3", "4"},
			remove:   vstring,
		},
	}
	for _, tC := range testCases {
		err := tC.Obj.Remove(tC.remove)
		if err != nil {
			t.Errorf("[Error Removing] : %v", err)
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
			Gotitem, _ := got.Get(i)
			Expecteditem, _ := tC.expected.Get(i)
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
			Gotitem, _ := got.Get(i)
			Expecteditem, _ := tC.expected.Get(i)
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

		for i := 0; i < got.Len(); i++ {
			Gotitem, _ := got.Get(i)
			Expecteditem, _ := tC.expected.Get(i)
			if Gotitem != Expecteditem {
				t.Errorf("Error [TestReverse] Got: %v Expected: %v \n.", got, tC.expected)
			}
		}

	}
}
