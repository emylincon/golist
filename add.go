package golist

import "github.com/emylincon/golist/core"

// Add :
// adds two list together and returns new list and possible error
func (arr *List) Add(other *List) (newList *List, err error) {
	if arr.Type() != other.Type() {
		return nil, ErrListsNotOfSameType
	} else if arr.Type() == TypeListUnknown {
		return nil, ErrTypeNotsupported
	}
	err = nil
	switch list := arr.list.(type) {
	case []int:
		otherList, ok := other.list.([]int)
		if !ok {
			return
		}
		newList = NewList(core.ExtendInt(&list, otherList))

	case []int32:
		otherList, ok := other.list.([]int32)
		if !ok {
			return
		}
		newList = NewList(core.ExtendInt32(&list, otherList))

	case []int64:
		otherList, ok := other.list.([]int64)
		if !ok {
			return
		}
		newList = NewList(core.ExtendInt64(&list, otherList))

	case []float32:
		otherList, ok := other.list.([]float32)
		if !ok {
			return
		}
		newList = NewList(core.ExtendFloat32(&list, otherList))

	case []float64:
		otherList, ok := other.list.([]float64)
		if !ok {
			return
		}
		newList = NewList(core.ExtendFloat64(&list, otherList))

	case []string:
		otherList, ok := other.list.([]string)
		if !ok {
			return
		}
		newList = NewList(core.ExtendString(&list, otherList))

	default:
		return nil, ErrTypeNotsupported
	}
	return
}
