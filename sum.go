package golist

import "github.com/emylincon/golist/core"

// Sum:
// sums all element in the list
func (arr *List) Sum() interface{} {
	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		return core.SumInt(&list)

	case []int32:
		list := arr.list.([]int32)
		return core.SumInt32(&list)

	case []int64:
		list := arr.list.([]int64)
		return core.SumInt64(&list)

	case []float32:
		list := arr.list.([]float32)
		return core.SumFloat32(&list)

	case []float64:
		list := arr.list.([]float64)
		return core.SumFloat64(&list)

	case []string:
		list := arr.list.([]string)
		return core.SumString(&list)

	default:
		return nil
	}
}
