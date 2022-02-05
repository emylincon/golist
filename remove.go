package golist

import (
	"github.com/emylincon/golist/core"
)

// Remove  :
// removes a given element from list. returns error if element do not exist
func (arr *List) Remove(element interface{}) error {

	switch list := arr.list.(type) {
	case []int:
		item, ok := element.(int)
		if !ok {
			return ErrTypeNotSame
		}
		newlist, err := core.RemoveInt(&list, item)
		arr.list = *newlist
		return err

	case []int32:
		item, ok := element.(int32)
		if !ok {
			return ErrTypeNotSame
		}
		newlist, err := core.RemoveInt32(&list, item)
		arr.list = *newlist
		return err

	case []int64:
		item, ok := element.(int64)
		if !ok {
			return ErrTypeNotSame
		}
		newlist, err := core.RemoveInt64(&list, item)
		arr.list = *newlist
		return err

	case []float32:
		item, ok := element.(float32)
		if !ok {
			return ErrTypeNotSame
		}
		newlist, err := core.RemoveFloat32(&list, item)
		arr.list = *newlist
		return err

	case []float64:
		item, ok := element.(float64)
		if !ok {
			return ErrTypeNotSame
		}
		newlist, err := core.RemoveFloat64(&list, item)
		arr.list = *newlist
		return err

	case []string:
		item, ok := element.(string)
		if !ok {
			return ErrTypeNotSame
		}
		newlist, err := core.RemoveString(&list, item)
		arr.list = *newlist
		return err

	default:
		return ErrTypeNotsupported
	}

}
