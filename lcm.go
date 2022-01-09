package golist

import (
	"github.com/emylincon/golist/core"
)

func (arr *List) LCM() (lcm interface{}, err error) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		return core.LCMInt(&list)

	case []int32:
		list := arr.list.([]int32)
		return core.LCMInt32(&list)

	case []int64:
		list := arr.list.([]int64)
		return core.LCMInt64(&list)

	case []float32:
		list := arr.list.([]float32)
		return core.LCMFloat32(&list)

	case []float64:
		list := arr.list.([]float64)
		return core.LCMFloat64(&list)

	case []string:
		return nil, ErrStringsNotsupported

	default:
		return nil, ErrTypeNotsupported
	}

}
