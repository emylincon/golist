package golist

import (
	"github.com/emylincon/golist/core"
)

// validateListSum validates
func validateListSum(ok bool, listLen, otherLen int) error {
	if !ok {
		return ErrListsNotOfSameType
	}
	if listLen != otherLen {
		return ErrListsNotOfSameLen
	}
	return nil
}

// ListSum returns the sum of contents of two lists
func (arr *List) ListSum(other *List) (*List, error) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		otherArr, ok := other.list.([]int)
		err := validateListSum(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSumInt(list, otherArr)
		return NewList(sum), nil

	case []int32:
		list := arr.list.([]int32)
		otherArr, ok := other.list.([]int32)
		err := validateListSum(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSumInt32(list, otherArr)
		return NewList(sum), nil

	case []int64:
		list := arr.list.([]int64)
		otherArr, ok := other.list.([]int64)
		err := validateListSum(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSumInt64(list, otherArr)
		return NewList(sum), nil

	case []float32:
		list := arr.list.([]float32)
		otherArr, ok := other.list.([]float32)
		err := validateListSum(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSumFloat32(list, otherArr)
		return NewList(sum), nil

	case []float64:
		list := arr.list.([]float64)
		otherArr, ok := other.list.([]float64)
		err := validateListSum(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSumFloat64(list, otherArr)
		return NewList(sum), nil

	case []string:
		return &List{}, ErrStringsNotsupported

	default:
		return &List{}, ErrTypeNotsupported
	}

}
