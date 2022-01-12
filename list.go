package golist

import (
	"github.com/emylincon/golist/core"
)

type Lists interface {
	List() interface{}
}

type List struct {
	list interface{}
}

func NewList(list interface{}) *List {
	newlist := &List{
		list: list,
	}
	return newlist
}

func (arr *List) List() interface{} {
	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		return list

	case []int32:
		list := arr.list.([]int32)
		return list

	case []int64:
		list := arr.list.([]int64)
		return list

	case []float32:
		list := arr.list.([]float32)
		return list

	case []float64:
		list := arr.list.([]float64)
		return list

	case []string:
		list := arr.list.([]string)
		return list

	default:
		return nil
	}
}

func (arr *List) Append(element interface{}) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := element.(int)
		arr.list = *core.AppendInt(&list, item)

	case []int32:
		list := arr.list.([]int32)
		item := element.(int32)
		arr.list = *core.AppendInt32(&list, item)

	case []int64:
		list := arr.list.([]int64)
		item := element.(int64)
		arr.list = *core.AppendInt64(&list, item)

	case []float32:
		list := arr.list.([]float32)
		item := element.(float32)
		arr.list = *core.AppendFloat32(&list, item)

	case []float64:
		list := arr.list.([]float64)
		item := element.(float64)
		arr.list = *core.AppendFloat64(&list, item)

	case []string:
		list := arr.list.([]string)
		item := element.(string)
		arr.list = *core.AppendString(&list, item)

	default:
		return
	}

}

func (arr *List) Index(element interface{}) int {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := element.(int)
		return core.IndexOfInt(&list, item)

	case []int32:
		list := arr.list.([]int32)
		item := element.(int32)
		return core.IndexOfInt32(&list, item)

	case []int64:
		list := arr.list.([]int64)
		item := element.(int64)
		return core.IndexOfInt64(&list, item)

	case []float32:
		list := arr.list.([]float32)
		item := element.(float32)
		return core.IndexOfFloat32(&list, item)

	case []float64:
		list := arr.list.([]float64)
		item := element.(float64)
		return core.IndexOfFloat64(&list, item)

	case []string:
		list := arr.list.([]string)
		item := element.(string)
		return core.IndexOfString(&list, item)

	default:
		return -1
	}

}

func (arr *List) Last() (interface{}, error) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		return core.LastInt(list)

	case []int32:
		list := arr.list.([]int32)

		return core.LastInt32(list)

	case []int64:
		list := arr.list.([]int64)
		return core.LastInt64(list)

	case []float32:
		list := arr.list.([]float32)
		return core.LastFloat32(list)

	case []float64:
		list := arr.list.([]float64)
		return core.LastFloat64(list)

	case []string:
		list := arr.list.([]string)
		return core.LastString(list)

	default:
		return nil, ErrTypeNotsupported
	}

}
