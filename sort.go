package golist

import "github.com/emylincon/golist/core"

func (arr *List) Sort() interface{} {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		sorted := core.SortInt(&list)
		return sorted

	case []int32:
		list := arr.list.([]int32)
		sorted := core.SortInt32(&list)
		return sorted

	case []int64:
		list := arr.list.([]int64)
		sorted := core.SortInt64(&list)
		return sorted

	case []float32:
		list := arr.list.([]float32)
		sorted := core.SortFloat32(&list)
		return sorted

	case []float64:
		list := arr.list.([]float64)
		sorted := core.SortFloat64(&list)
		return sorted

	case []string:
		list := arr.list.([]string)
		sorted := core.SortString(&list)
		return sorted

	default:
		return nil
	}

}
