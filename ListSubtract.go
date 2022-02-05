package golist

import (
	"github.com/emylincon/golist/core"
)

// ListSubtract subtracts the contents of two lists
func (arr *List) ListSubtract(other *List) (*List, error) {

	switch list := arr.list.(type) {
	case []int:
		otherArr, ok := other.list.([]int)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSubtractInt(list, otherArr)
		return NewList(sum), nil

	case []int32:
		otherArr, ok := other.list.([]int32)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSubtractInt32(list, otherArr)
		return NewList(sum), nil

	case []int64:
		otherArr, ok := other.list.([]int64)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSubtractInt64(list, otherArr)
		return NewList(sum), nil

	case []float32:
		otherArr, ok := other.list.([]float32)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSubtractFloat32(list, otherArr)
		return NewList(sum), nil

	case []float64:
		otherArr, ok := other.list.([]float64)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListSubtractFloat64(list, otherArr)
		return NewList(sum), nil

	case []string:
		return &List{}, ErrStringsNotsupported

	default:
		return &List{}, ErrTypeNotsupported
	}

}

// ListSubtractNo subtracts a given number from all elements in list
func (arr *List) ListSubtractNo(no interface{}) (*List, error) {

	switch list := arr.list.(type) {
	case []int:
		o, ok := no.(int)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListSubtractNoInt(list, o)
		return NewList(sum), nil

	case []int32:
		o, ok := no.(int32)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListSubtractNoInt32(list, o)
		return NewList(sum), nil

	case []int64:
		o, ok := no.(int64)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListSubtractNoInt64(list, o)
		return NewList(sum), nil

	case []float32:
		o, ok := no.(float32)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListSubtractNoFloat32(list, o)
		return NewList(sum), nil

	case []float64:
		o, ok := no.(float64)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListSubtractNoFloat64(list, o)
		return NewList(sum), nil

	case []string:
		return &List{}, ErrStringsNotsupported

	default:
		return &List{}, ErrTypeNotsupported
	}

}
