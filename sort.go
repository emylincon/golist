package golist

import "github.com/emylincon/golist/core"

func (arr *List) Sort(reverse bool) interface{} {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		sorted := core.SortInt(&list, reverse)
		arr.list = *sorted
		return *sorted

	case []int32:
		list := arr.list.([]int32)
		sorted := core.SortInt32(&list, reverse)
		arr.list = *sorted
		return *sorted

	case []int64:
		list := arr.list.([]int64)
		sorted := core.SortInt64(&list, reverse)
		arr.list = *sorted
		return *sorted

	case []float32:
		list := arr.list.([]float32)
		sorted := core.SortFloat32(&list, reverse)
		arr.list = *sorted
		return *sorted

	case []float64:
		list := arr.list.([]float64)
		sorted := core.SortFloat64(&list, reverse)
		arr.list = *sorted
		return *sorted

	case []string:
		list := arr.list.([]string)
		sorted := core.SortString(&list, reverse)
		arr.list = *sorted
		return *sorted

	default:
		return nil
	}

}
