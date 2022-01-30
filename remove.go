package golist

import (
	"github.com/emylincon/golist/core"
)

// Remove  :
// removes a given element from list. returns error if element do not exist
func (arr *List) Remove(element interface{}) error {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := element.(int)

		newlist, err := core.RemoveInt(&list, item)
		arr.list = *newlist
		return err

	case []int32:
		list := arr.list.([]int32)
		item := element.(int32)

		newlist, err := core.RemoveInt32(&list, item)
		arr.list = *newlist
		return err

	case []int64:
		list := arr.list.([]int64)
		item := element.(int64)

		newlist, err := core.RemoveInt64(&list, item)
		arr.list = *newlist
		return err

	case []float32:
		list := arr.list.([]float32)
		item := element.(float32)

		newlist, err := core.RemoveFloat32(&list, item)
		arr.list = *newlist
		return err

	case []float64:
		list := arr.list.([]float64)
		item := element.(float64)

		newlist, err := core.RemoveFloat64(&list, item)
		arr.list = *newlist
		return err

	case []string:
		list := arr.list.([]string)
		item := element.(string)

		newlist, err := core.RemoveString(&list, item)
		arr.list = *newlist
		return err

	default:
		return ErrTypeNotsupported
	}

}
