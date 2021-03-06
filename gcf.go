package golist

import (
	"github.com/emylincon/golist/core"
)

// GCF  :
// returns greatest common factor of the list. returns error if type of list is string or any of the unsupported types
func (arr *List) GCF() (gcf interface{}, err error) {

	switch list := arr.list.(type) {
	case []int:
		return core.GCFInt(&list)

	case []int32:
		return core.GCFInt32(&list)

	case []int64:
		return core.GCFInt64(&list)

	case []float32:
		return core.GCFFloat32(&list)

	case []float64:
		return core.GCFFloat64(&list)

	case []string:
		return nil, ErrStringsNotsupported

	default:
		return nil, ErrTypeNotsupported
	}

}

// HCF  :
// returns greatest common factor of the list. returns error if type of list is string or any of the unsupported types
func (arr *List) HCF() (gcf interface{}, err error) {
	return arr.GCF()
}
