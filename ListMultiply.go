package golist

import (
	"github.com/emylincon/golist/core"
)

// ListMultiply returns the product of contents of two lists
func (arr *List) ListMultiply(other *List) (*List, error) {

	switch list := arr.list.(type) {
	case []int:
		otherArr, ok := other.list.([]int)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListMultiplyInt(list, otherArr)
		return NewList(sum), nil

	case []int32:
		otherArr, ok := other.list.([]int32)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListMultiplyInt32(list, otherArr)
		return NewList(sum), nil

	case []int64:
		otherArr, ok := other.list.([]int64)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListMultiplyInt64(list, otherArr)
		return NewList(sum), nil

	case []float32:
		otherArr, ok := other.list.([]float32)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListMultiplyFloat32(list, otherArr)
		return NewList(sum), nil

	case []float64:
		otherArr, ok := other.list.([]float64)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListMultiplyFloat64(list, otherArr)
		return NewList(sum), nil

	case []string:
		return &List{}, ErrStringsNotsupported

	default:
		return &List{}, ErrTypeNotsupported
	}

}

// ListMultiplyNo multiply a given number with all elements in list
func (arr *List) ListMultiplyNo(no interface{}) (*List, error) {

	switch list := arr.list.(type) {
	case []int:
		o, ok := no.(int)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListMultiplyNoInt(list, o)
		return NewList(sum), nil

	case []int32:
		o, ok := no.(int32)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListMultiplyNoInt32(list, o)
		return NewList(sum), nil

	case []int64:
		o, ok := no.(int64)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListMultiplyNoInt64(list, o)
		return NewList(sum), nil

	case []float32:
		o, ok := no.(float32)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListMultiplyNoFloat32(list, o)
		return NewList(sum), nil

	case []float64:
		o, ok := no.(float64)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListMultiplyNoFloat64(list, o)
		return NewList(sum), nil

	case []string:
		return &List{}, ErrStringsNotsupported

	default:
		return &List{}, ErrTypeNotsupported
	}

}
