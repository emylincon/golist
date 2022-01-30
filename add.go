package golist

import "github.com/emylincon/golist/core"

// Add:
// adds two list together and returns new list and possible error
func (arr *List) Add(other *List) (newList *List, err error) {
	if arr.Type() != other.Type() {
		return nil, ErrListsNotOfSameType
	} else if arr.Type() == TypeListUnknown {
		return nil, ErrTypeNotsupported
	}
	err = nil
	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		otherList := other.list.([]int)
		newList = NewList(core.ExtendInt(&list, otherList))

	case []int32:
		list := arr.list.([]int32)
		otherList := other.list.([]int32)
		newList = NewList(core.ExtendInt32(&list, otherList))

	case []int64:
		list := arr.list.([]int64)
		otherList := other.list.([]int64)
		newList = NewList(core.ExtendInt64(&list, otherList))

	case []float32:
		list := arr.list.([]float32)
		otherList := other.list.([]float32)
		newList = NewList(core.ExtendFloat32(&list, otherList))

	case []float64:
		list := arr.list.([]float64)
		otherList := other.list.([]float64)
		newList = NewList(core.ExtendFloat64(&list, otherList))

	case []string:
		list := arr.list.([]string)
		otherList := other.list.([]string)
		newList = NewList(core.ExtendString(&list, otherList))

	default:
		return nil, ErrTypeNotsupported
	}
	return
}
