package golist

import "github.com/emylincon/golist/core"

// Sum  :
// sums all element in the list
func (arr *List) Sum() interface{} {
	switch list := arr.list.(type) {
	case []int:
		return core.SumInt(&list)

	case []int32:
		return core.SumInt32(&list)

	case []int64:
		return core.SumInt64(&list)

	case []float32:
		return core.SumFloat32(&list)

	case []float64:
		return core.SumFloat64(&list)

	case []string:
		return core.SumString(&list)

	default:
		return nil
	}
}
