package golist

import (
	"github.com/emylincon/golist/core"
)

// ListDivide divides list with other list
func (arr *List) ListDivide(other *List) (*List, error) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		otherArr, ok := other.list.([]int)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListDivideInt(list, otherArr)
		return NewList(sum), nil

	case []int32:
		list := arr.list.([]int32)
		otherArr, ok := other.list.([]int32)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListDivideInt32(list, otherArr)
		return NewList(sum), nil

	case []int64:
		list := arr.list.([]int64)
		otherArr, ok := other.list.([]int64)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListDivideInt64(list, otherArr)
		return NewList(sum), nil

	case []float32:
		list := arr.list.([]float32)
		otherArr, ok := other.list.([]float32)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListDivideFloat32(list, otherArr)
		return NewList(sum), nil

	case []float64:
		list := arr.list.([]float64)
		otherArr, ok := other.list.([]float64)
		err := validateListOp(ok, arr.Len(), other.Len())
		if err != nil {
			return &List{}, err
		}
		sum := core.ListDivideFloat64(list, otherArr)
		return NewList(sum), nil

	case []string:
		return &List{}, ErrStringsNotsupported

	default:
		return &List{}, ErrTypeNotsupported
	}

}

// ListDivideNo Divide all elements in list with a number
func (arr *List) ListDivideNo(no interface{}) (*List, error) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		o, ok := no.(int)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListDivideNoInt(list, o)
		return NewList(sum), nil

	case []int32:
		list := arr.list.([]int32)
		o, ok := no.(int32)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListDivideNoInt32(list, o)
		return NewList(sum), nil

	case []int64:
		list := arr.list.([]int64)
		o, ok := no.(int64)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListDivideNoInt64(list, o)
		return NewList(sum), nil

	case []float32:
		list := arr.list.([]float32)
		o, ok := no.(float32)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListDivideNoFloat32(list, o)
		return NewList(sum), nil

	case []float64:
		list := arr.list.([]float64)
		o, ok := no.(float64)
		if !ok {
			return &List{}, ErrTypeNotSame
		}
		sum := core.ListDivideNoFloat64(list, o)
		return NewList(sum), nil

	case []string:
		return &List{}, ErrStringsNotsupported

	default:
		return &List{}, ErrTypeNotsupported
	}

}
