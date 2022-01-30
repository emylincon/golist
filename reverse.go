package golist

import "github.com/emylincon/golist/core"

// Reverse  :
// reverse elements in the list
func (arr *List) Reverse() *List {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		newList := *core.ReverseInt(&list)
		return NewList(newList)

	case []int32:
		list := arr.list.([]int32)
		newList := *core.ReverseInt32(&list)
		return NewList(newList)

	case []int64:
		list := arr.list.([]int64)
		newList := *core.ReverseInt64(&list)
		return NewList(newList)

	case []float32:
		list := arr.list.([]float32)
		newList := *core.ReverseFloat32(&list)
		return NewList(newList)

	case []float64:
		list := arr.list.([]float64)
		newList := *core.ReverseFloat64(&list)
		return NewList(newList)

	case []string:
		list := arr.list.([]string)
		newList := *core.ReverseString(&list)
		return NewList(newList)

	default:
		return NewList([]int{})
	}

}
