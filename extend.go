package golist

import "github.com/emylincon/golist/core"

// extends current list with input slice. errors if both underlying lists are not of the same type
func (arr *List) Extend(other interface{}) error {
	if arr.Type() != NewList(other).Type() {
		return ErrListsNotOfSameType
	} else if arr.Type() == TypeListUnknown {
		return ErrTypeNotsupported
	}
	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := other.([]int)
		arr.list = core.ExtendInt(&list, item)

	case []int32:
		list := arr.list.([]int32)
		item := other.([]int32)
		arr.list = core.ExtendInt32(&list, item)

	case []int64:
		list := arr.list.([]int64)
		item := other.([]int64)
		arr.list = core.ExtendInt64(&list, item)

	case []float32:
		list := arr.list.([]float32)
		item := other.([]float32)
		arr.list = core.ExtendFloat32(&list, item)

	case []float64:
		list := arr.list.([]float64)
		item := other.([]float64)
		arr.list = core.ExtendFloat64(&list, item)

	case []string:
		list := arr.list.([]string)
		item := other.([]string)
		arr.list = core.ExtendString(&list, item)

	default:
		return ErrTypeNotsupported
	}
	return nil

}
