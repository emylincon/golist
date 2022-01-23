package golist

import "github.com/emylincon/golist/core"

// ListSum returns the sum of contents of two lists
func (arr *List) ListSum(other *List) (*List, error) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		otherArr, ok := other.list.([]int)
		if !ok {
			return nil, ErrListsNotOfSameType
		}
		sum := core.ListSumInt(list, otherArr)
		return NewList(sum), nil

	case []int32:
		list := arr.list.([]int32)
		otherArr, ok := other.list.([]int32)
		if !ok {
			return nil, ErrListsNotOfSameType
		}
		sum := core.ListSumInt32(list, otherArr)
		return NewList(sum), nil

	case []int64:
		list := arr.list.([]int64)
		otherArr, ok := other.list.([]int64)
		if !ok {
			return nil, ErrListsNotOfSameType
		}
		sum := core.ListSumInt64(list, otherArr)
		return NewList(sum), nil

	case []float32:
		list := arr.list.([]float32)
		otherArr, ok := other.list.([]float32)
		if !ok {
			return nil, ErrListsNotOfSameType
		}
		sum := core.ListSumFloat32(list, otherArr)
		return NewList(sum), nil

	case []float64:
		list := arr.list.([]float64)
		otherArr, ok := other.list.([]float64)
		if !ok {
			return nil, ErrListsNotOfSameType
		}
		sum := core.ListSumFloat64(list, otherArr)
		return NewList(sum), nil

	case []string:
		return nil, ErrStringsNotsupported

	default:
		return nil, ErrTypeNotsupported
	}

}
