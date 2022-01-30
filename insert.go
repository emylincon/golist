package golist

import (
	"github.com/emylincon/golist/core"
)

// Insert :
// inserts an element at a given location. returns error if index is out of range
func (arr *List) Insert(element interface{}, index int) error {
	if index > arr.Len() {
		return ErrIndexOutOfRange
	}
	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := element.(int)
		arr.list = *core.InsertInt(&list, item, index)
		return nil

	case []int32:
		list := arr.list.([]int32)
		item := element.(int32)
		arr.list = *core.InsertInt32(&list, item, index)
		return nil

	case []int64:
		list := arr.list.([]int64)
		item := element.(int64)
		arr.list = *core.InsertInt64(&list, item, index)
		return nil

	case []float32:
		list := arr.list.([]float32)
		item := element.(float32)
		arr.list = *core.InsertFloat32(&list, item, index)
		return nil

	case []float64:
		list := arr.list.([]float64)
		item := element.(float64)
		arr.list = *core.InsertFloat64(&list, item, index)
		return nil

	case []string:
		list := arr.list.([]string)
		item := element.(string)
		arr.list = *core.InsertString(&list, item, index)
		return nil

	default:
		return ErrTypeNotsupported
	}

}
