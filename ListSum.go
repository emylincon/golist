package golist

import (
	"github.com/emylincon/golist/core"
)

// validateListOp validates
func validateListOp(ok bool, listLen, otherLen int) error {
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

	switch list := arr.list.(type) {
	case []int:
		otherArr, ok := other.list.([]int)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSumInt(list, otherArr)
		return NewList(sum), nil

	case []int32:
		otherArr, ok := other.list.([]int32)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSumInt32(list, otherArr)
		return NewList(sum), nil

	case []int64:
		otherArr, ok := other.list.([]int64)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSumInt64(list, otherArr)
		return NewList(sum), nil

	case []float32:
		otherArr, ok := other.list.([]float32)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSumFloat32(list, otherArr)
		return NewList(sum), nil

	case []float64:
		otherArr, ok := other.list.([]float64)
		err := validateListOp(ok, arr.Len(), other.Len())
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

// ListSumNo adds a given number to all elements in list
func (arr *List) ListSumNo(no interface{}) (*List, error) {

	switch list := arr.list.(type) {
	case []int:
		o, ok := no.(int)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListSumNoInt(list, o)
		return NewList(sum), nil

	case []int32:
		o, ok := no.(int32)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListSumNoInt32(list, o)
		return NewList(sum), nil

	case []int64:
		o, ok := no.(int64)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListSumNoInt64(list, o)
		return NewList(sum), nil

	case []float32:
		o, ok := no.(float32)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListSumNoFloat32(list, o)
		return NewList(sum), nil

	case []float64:
		o, ok := no.(float64)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListSumNoFloat64(list, o)
		return NewList(sum), nil

	case []string:
		return &List{}, ErrStringsNotsupported

	default:
		return &List{}, ErrTypeNotsupported
	}

}
