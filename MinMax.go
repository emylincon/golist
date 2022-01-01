package golist

import (
	"errors"

	"github.com/emylincon/golist/core"
)

func (arr *List) Max() (interface{}, error) {
	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		if len(list) == 0 {
			return nil, errors.New("list is empty")
		}
		return core.MaxInt(&list), nil

	case []int32:
		list := arr.list.([]int32)
		if len(list) == 0 {
			return nil, errors.New("list is empty")
		}
		return core.MaxInt32(&list), nil

	case []int64:
		list := arr.list.([]int64)
		if len(list) == 0 {
			return nil, errors.New("list is empty")
		}
		return core.MaxInt64(&list), nil

	case []float32:
		list := arr.list.([]float32)
		if len(list) == 0 {
			return nil, errors.New("list is empty")
		}
		return core.MaxFloat32(&list), nil

	case []float64:
		list := arr.list.([]float64)
		if len(list) == 0 {
			return nil, errors.New("list is empty")
		}
		return core.MaxFloat64(&list), nil

	case []string:
		list := arr.list.([]string)
		if len(list) == 0 {
			return nil, errors.New("list is empty")
		}
		return core.MaxString(&list), nil

	default:
		return nil, errors.New("unsupported type: interface type should be []float32, []float64, []int32, []int, []int64, or []string")

	}

}

func (arr *List) Min() (interface{}, error) {
	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		if len(list) == 0 {
			return nil, errors.New("list is empty")
		}
		return core.MinInt(&list), nil

	case []int32:
		list := arr.list.([]int32)
		if len(list) == 0 {
			return nil, errors.New("list is empty")
		}
		return core.MinInt32(&list), nil

	case []int64:
		list := arr.list.([]int64)
		if len(list) == 0 {
			return nil, errors.New("list is empty")
		}
		return core.MinInt64(&list), nil

	case []float32:
		list := arr.list.([]float32)
		if len(list) == 0 {
			return nil, errors.New("list is empty")
		}
		return core.MinFloat32(&list), nil

	case []float64:
		list := arr.list.([]float64)
		if len(list) == 0 {
			return nil, errors.New("list is empty")
		}
		return core.MinFloat64(&list), nil

	case []string:
		list := arr.list.([]string)
		if len(list) == 0 {
			return nil, errors.New("list is empty")
		}
		return core.MinString(&list), nil

	default:
		return nil, errors.New("unsupported type: interface type should be []float32, []float64, []int32, []int, []int64, or []string")

	}

}
