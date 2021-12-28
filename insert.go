package golist

import (
	"errors"

	"github.com/emylincon/golist/core"
)

func (arr *List) Insert(element interface{}, index int) error {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := element.(int)
		if index > arr.Len() {
			return errors.New("index error: list index out of range")
		}
		arr.list = *core.InsertInt(&list, item, index)
		return nil

	case []int32:
		list := arr.list.([]int32)
		item := element.(int32)
		if index > arr.Len() {
			return errors.New("index error: list index out of range")
		}
		arr.list = *core.InsertInt32(&list, item, index)
		return nil

	case []int64:
		list := arr.list.([]int64)
		item := element.(int64)
		if index > arr.Len() {
			return errors.New("index error: list index out of range")
		}
		arr.list = *core.InsertInt64(&list, item, index)
		return nil

	case []float32:
		list := arr.list.([]float32)
		item := element.(float32)
		if index > arr.Len() {
			return errors.New("index error: list index out of range")
		}
		arr.list = *core.InsertFloat32(&list, item, index)
		return nil

	case []float64:
		list := arr.list.([]float64)
		item := element.(float64)
		if index > arr.Len() {
			return errors.New("index error: list index out of range")
		}
		arr.list = *core.InsertFloat64(&list, item, index)
		return nil

	case []string:
		list := arr.list.([]string)
		item := element.(string)
		if index > arr.Len() {
			return errors.New("index error: list index out of range")
		}
		arr.list = *core.InsertString(&list, item, index)
		return nil

	default:
		return errors.New("interface type should be []float32, []float64, []int32, []int, []int64, or []string")
	}

}
