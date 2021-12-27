package golist

import "github.com/emylincon/golist/core"

func (arr *List) Extend(element interface{}) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := element.([]int)
		arr.list = core.ExtendInt(&list, item)

	case []int32:
		list := arr.list.([]int32)
		item := element.([]int32)
		arr.list = core.ExtendInt32(&list, item)

	case []int64:
		list := arr.list.([]int64)
		item := element.([]int64)
		arr.list = core.ExtendInt64(&list, item)

	case []float32:
		list := arr.list.([]float32)
		item := element.([]float32)
		arr.list = core.ExtendFloat32(&list, item)

	case []float64:
		list := arr.list.([]float64)
		item := element.([]float64)
		arr.list = core.ExtendFloat64(&list, item)

	case []string:
		list := arr.list.([]string)
		item := element.([]string)
		arr.list = core.ExtendString(&list, item)

	default:
		return
	}

}
