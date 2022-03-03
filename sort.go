package golist

import "github.com/emylincon/golist/core"

// Sort  :
// sorts elements in the list in ascending order or descending order in place.
func (arr *List) Sort(reverse bool) {

	switch list := arr.list.(type) {
	case []int:
		sorted := core.SortInt(&list, reverse)
		arr.list = *sorted

	case []int32:
		sorted := core.SortInt32(&list, reverse)
		arr.list = *sorted

	case []int64:
		sorted := core.SortInt64(&list, reverse)
		arr.list = *sorted

	case []float32:
		sorted := core.SortFloat32(&list, reverse)
		arr.list = *sorted

	case []float64:
		sorted := core.SortFloat64(&list, reverse)
		arr.list = *sorted

	case []string:
		sorted := core.SortString(&list, reverse)
		arr.list = *sorted

	}

}

// Sorted  :
// sorts elements in the list in ascending order or descending order and returns new list.
func (arr *List) Sorted(reverse bool) *List {

	switch list := arr.list.(type) {
	case []int:
		sorted := core.SortInt(&list, reverse)
		return NewList(*sorted)

	case []int32:
		sorted := core.SortInt32(&list, reverse)
		return NewList(*sorted)

	case []int64:
		sorted := core.SortInt64(&list, reverse)
		return NewList(*sorted)

	case []float32:
		sorted := core.SortFloat32(&list, reverse)
		return NewList(*sorted)

	case []float64:
		sorted := core.SortFloat64(&list, reverse)
		return NewList(*sorted)

	case []string:
		sorted := core.SortString(&list, reverse)
		return NewList(*sorted)

	default:
		return nil
	}

}
