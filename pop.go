package golist

import "github.com/emylincon/golist/core"

func (arr *List) Pop(index int) interface{} {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		var popped int
		arr.list, popped = core.PopInt(&list, index)
		return popped

	case []int32:
		list := arr.list.([]int32)
		var popped int32
		arr.list, popped = core.PopInt32(&list, index)
		return popped

	case []int64:
		list := arr.list.([]int64)
		var popped int64
		arr.list, popped = core.PopInt64(&list, index)
		return popped

	case []float32:
		list := arr.list.([]float32)
		var popped float32
		arr.list, popped = core.PopFloat32(&list, index)
		return popped

	case []float64:
		list := arr.list.([]float64)
		var popped float64
		arr.list, popped = core.PopFloat64(&list, index)
		return popped

	case []string:
		list := arr.list.([]string)
		var popped string
		arr.list, popped = core.PopString(&list, index)
		return popped

	default:
		return nil
	}

}
