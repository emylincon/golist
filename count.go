package golist

import "github.com/emylincon/golist/core"

// Count returns a count of how many times a given element appears in the list
// Returns -1 if element not found or the given element type does not match the type of list
func (arr *List) Count(element interface{}) int {

	switch list := arr.list.(type) {
	case []int:
		item, ok := element.(int)
		if !ok {
			return -1
		}
		return core.CountInt(&list, item)

	case []int32:
		item, ok := element.(int32)
		if !ok {
			return -1
		}
		return core.CountInt32(&list, item)

	case []int64:
		item, ok := element.(int64)
		if !ok {
			return -1
		}
		return core.CountInt64(&list, item)

	case []float32:
		item, ok := element.(float32)
		if !ok {
			return -1
		}
		return core.CountFloat32(&list, item)

	case []float64:
		item, ok := element.(float64)
		if !ok {
			return -1
		}
		return core.CountFloat64(&list, item)

	case []string:
		item, ok := element.(string)
		if !ok {
			return -1
		}
		return core.CountString(&list, item)

	default:
		return -1
	}

}
