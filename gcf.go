package golist

import (
	"errors"

	"github.com/emylincon/golist/core"
)

func (arr *List) GCF() (gcf interface{}, err error) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		return core.GCFInt(&list)

	case []int32:
		list := arr.list.([]int32)
		return core.GCFInt32(&list)

	case []int64:
		list := arr.list.([]int64)
		return core.GCFInt64(&list)

	case []float32:
		list := arr.list.([]float32)
		return core.GCFfloat32(&list)

	case []float64:
		list := arr.list.([]float64)
		return core.GCFfloat64(&list)

	case []string:
		return nil, errors.New("strings are not supported for this operation")

	default:
		return nil, ErrTypeNotsupported
	}

}

func (arr *List) HCF() (gcf interface{}, err error) {
	return arr.GCF()
}
