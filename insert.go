package golist

import (
	"github.com/emylincon/golist/core"
)

// Insert  :
// inserts an element at a given location. returns error if index is out of range
func (arr *List) Insert(element interface{}, index int) error {
	if index > arr.Len() {
		return ErrIndexOutOfRange
	}
	switch list := arr.list.(type) {
	case []int:
		item, ok := element.(int)
		if !ok {
			return ErrTypeNotSame
		}
		arr.list = *core.InsertInt(&list, item, index)
		return nil

	case []int32:
		item, ok := element.(int32)
		if !ok {
			return ErrTypeNotSame
		}
		arr.list = *core.InsertInt32(&list, item, index)
		return nil

	case []int64:
		item, ok := element.(int64)
		if !ok {
			return ErrTypeNotSame
		}
		arr.list = *core.InsertInt64(&list, item, index)
		return nil

	case []float32:
		item, ok := element.(float32)
		if !ok {
			return ErrTypeNotSame
		}
		arr.list = *core.InsertFloat32(&list, item, index)
		return nil

	case []float64:
		item, ok := element.(float64)
		if !ok {
			return ErrTypeNotSame
		}
		arr.list = *core.InsertFloat64(&list, item, index)
		return nil

	case []string:
		item, ok := element.(string)
		if !ok {
			return ErrTypeNotSame
		}
		arr.list = *core.InsertString(&list, item, index)
		return nil

	default:
		return ErrTypeNotsupported
	}

}
