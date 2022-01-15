package golist

import "github.com/emylincon/golist/core"

// returns a count of how many times a given element appears in the list
// panics if given element type does not match the type of list
func (arr *List) Count(element interface{}) int {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := element.(int)
		return core.CountInt(&list, item)

	case []int32:
		list := arr.list.([]int32)
		item := element.(int32)
		return core.CountInt32(&list, item)

	case []int64:
		list := arr.list.([]int64)
		item := element.(int64)
		return core.CountInt64(&list, item)

	case []float32:
		list := arr.list.([]float32)
		item := element.(float32)
		return core.CountFloat32(&list, item)

	case []float64:
		list := arr.list.([]float64)
		item := element.(float64)
		return core.CountFloat64(&list, item)

	case []string:
		list := arr.list.([]string)
		item := element.(string)
		return core.CountString(&list, item)

	default:
		return -1
	}

}
