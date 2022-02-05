package golist

import "github.com/emylincon/golist/core"

// Extend : This extends current list with input slice. errors if both underlying lists are not of the same type
func (arr *List) Extend(other interface{}) error {
	if arr.Type() != NewList(other).Type() {
		return ErrListsNotOfSameType
	} else if arr.Type() == TypeListUnknown {
		return ErrTypeNotsupported
	}
	switch list := arr.list.(type) {
	case []int:
		item, ok := other.([]int)
		if !ok {
			return ErrTypeNotSame
		}
		arr.list = core.ExtendInt(&list, item)

	case []int32:
		item, ok := other.([]int32)
		if !ok {
			return ErrTypeNotSame
		}
		arr.list = core.ExtendInt32(&list, item)

	case []int64:
		item, ok := other.([]int64)
		if !ok {
			return ErrTypeNotSame
		}
		arr.list = core.ExtendInt64(&list, item)

	case []float32:
		item, ok := other.([]float32)
		if !ok {
			return ErrTypeNotSame
		}
		arr.list = core.ExtendFloat32(&list, item)

	case []float64:
		item, ok := other.([]float64)
		if !ok {
			return ErrTypeNotSame
		}
		arr.list = core.ExtendFloat64(&list, item)

	case []string:
		item, ok := other.([]string)
		if !ok {
			return ErrTypeNotSame
		}
		arr.list = core.ExtendString(&list, item)

	default:
		return ErrTypeNotsupported
	}
	return nil

}
