package golist

import "errors"

func (arr *List) Get(index int) (interface{}, error) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		if index >= arr.Len() {
			return nil, errors.New("index error: list index out of range")
		}
		return list[index], nil

	case []int32:
		list := arr.list.([]int32)
		if index >= arr.Len() {
			return nil, errors.New("index error: list index out of range")
		}
		return list[index], nil

	case []int64:
		list := arr.list.([]int64)
		if index >= arr.Len() {
			return nil, errors.New("index error: list index out of range")
		}
		return list[index], nil

	case []float32:
		list := arr.list.([]float32)
		if index >= arr.Len() {
			return nil, errors.New("index error: list index out of range")
		}
		return list[index], nil

	case []float64:
		list := arr.list.([]float64)
		if index >= arr.Len() {
			return nil, errors.New("index error: list index out of range")
		}
		return list[index], nil

	case []string:
		list := arr.list.([]string)
		if index >= arr.Len() {
			return nil, errors.New("index error: list index out of range")
		}
		return list[index], nil

	default:
		return nil, errors.New("interface type should be []float32, []float64, []int32, []int, []int64, or []string")
	}

}
