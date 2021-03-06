package golist

import "github.com/emylincon/golist/core"

// Set  :
// removes duplicates from list and returns new list
func (arr *List) Set() (setList *List, err error) {
	err = nil
	switch list := arr.list.(type) {
	case []int:
		setList = NewList(core.SetInt(list))
		return

	case []int32:
		setList = NewList(core.SetInt32(list))
		return

	case []int64:
		setList = NewList(core.SetInt64(list))
		return

	case []float32:
		setList = NewList(core.SetFloat32(list))
		return

	case []float64:
		setList = NewList(core.SetFloat64(list))
		return

	case []string:
		setList = NewList(core.SetString(list))
		return

	default:
		return nil, ErrTypeNotsupported
	}

}
