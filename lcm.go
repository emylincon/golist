package golist

import (
	"github.com/emylincon/golist/core"
)

// LCM  :
// returns Lowest common multiple of a list. returns error if list type is string
func (arr *List) LCM() (lcm interface{}, err error) {

	switch list := arr.list.(type) {
	case []int:
		return core.LCMInt(&list)

	case []int32:
		return core.LCMInt32(&list)

	case []int64:
		return core.LCMInt64(&list)

	case []float32:
		return core.LCMFloat32(&list)

	case []float64:
		return core.LCMFloat64(&list)

	case []string:
		return nil, ErrStringsNotsupported

	default:
		return nil, ErrTypeNotsupported
	}

}
