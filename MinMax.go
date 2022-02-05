package golist

import (
	"github.com/emylincon/golist/core"
)

// Max  :
// returns max element in list
func (arr *List) Max() (interface{}, error) {
	switch list := arr.list.(type) {
	case []int:
		if len(list) == 0 {
			return nil, core.ErrListEmpty
		}
		return core.MaxInt(&list), nil

	case []int32:
		if len(list) == 0 {
			return nil, core.ErrListEmpty
		}
		return core.MaxInt32(&list), nil

	case []int64:
		if len(list) == 0 {
			return nil, core.ErrListEmpty
		}
		return core.MaxInt64(&list), nil

	case []float32:
		if len(list) == 0 {
			return nil, core.ErrListEmpty
		}
		return core.MaxFloat32(&list), nil

	case []float64:
		if len(list) == 0 {
			return nil, core.ErrListEmpty
		}
		return core.MaxFloat64(&list), nil

	case []string:
		if len(list) == 0 {
			return nil, core.ErrListEmpty
		}
		return core.MaxString(&list), nil

	default:
		return nil, ErrTypeNotsupported
	}

}

// Min  :
// returns min element in list
func (arr *List) Min() (interface{}, error) {
	switch list := arr.list.(type) {
	case []int:
		if len(list) == 0 {
			return nil, core.ErrListEmpty
		}
		return core.MinInt(&list), nil

	case []int32:
		if len(list) == 0 {
			return nil, core.ErrListEmpty
		}
		return core.MinInt32(&list), nil

	case []int64:
		if len(list) == 0 {
			return nil, core.ErrListEmpty
		}
		return core.MinInt64(&list), nil

	case []float32:
		if len(list) == 0 {
			return nil, core.ErrListEmpty
		}
		return core.MinFloat32(&list), nil

	case []float64:
		if len(list) == 0 {
			return nil, core.ErrListEmpty
		}
		return core.MinFloat64(&list), nil

	case []string:
		if len(list) == 0 {
			return nil, core.ErrListEmpty
		}
		return core.MinString(&list), nil

	default:
		return nil, ErrTypeNotsupported
	}

}
