package golist

import "github.com/emylincon/golist/core"

// adds two list together. panics if both lists are not of the same type
func (arr *List) Extend(slice interface{}) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := slice.([]int)
		arr.list = core.ExtendInt(&list, item)

	case []int32:
		list := arr.list.([]int32)
		item := slice.([]int32)
		arr.list = core.ExtendInt32(&list, item)

	case []int64:
		list := arr.list.([]int64)
		item := slice.([]int64)
		arr.list = core.ExtendInt64(&list, item)

	case []float32:
		list := arr.list.([]float32)
		item := slice.([]float32)
		arr.list = core.ExtendFloat32(&list, item)

	case []float64:
		list := arr.list.([]float64)
		item := slice.([]float64)
		arr.list = core.ExtendFloat64(&list, item)

	case []string:
		list := arr.list.([]string)
		item := slice.([]string)
		arr.list = core.ExtendString(&list, item)

	default:
		return
	}

}
