package golist

import "errors"

func (arr *List) Slice(start int, stop int) (*List, error) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		if start > len(list) || stop > len(list) {
			return nil, errors.New("index out of range")
		} else if start < 0 || stop < -1 {
			return nil, errors.New("index out of range")
		} else if stop == -1 {
			stop = len(list)
		}
		return NewList(list[start:stop]), nil

	case []int32:
		list := arr.list.([]int32)
		if start > len(list) || stop > len(list) {
			return nil, errors.New("index out of range")
		} else if start < 0 || stop < -1 {
			return nil, errors.New("index out of range")
		} else if stop == -1 {
			stop = len(list)
		}
		return NewList(list[start:stop]), nil

	case []int64:
		list := arr.list.([]int64)
		if start > len(list) || stop > len(list) {
			return nil, errors.New("index out of range")
		} else if start < 0 || stop < -1 {
			return nil, errors.New("index out of range")
		} else if stop == -1 {
			stop = len(list)
		}
		return NewList(list[start:stop]), nil

	case []float32:
		list := arr.list.([]float32)
		if start > len(list) || stop > len(list) {
			return nil, errors.New("index out of range")
		} else if start < 0 || stop < -1 {
			return nil, errors.New("index out of range")
		} else if stop == -1 {
			stop = len(list)
		}
		return NewList(list[start:stop]), nil

	case []float64:
		list := arr.list.([]float64)
		if start > len(list) || stop > len(list) {
			return nil, errors.New("index out of range")
		} else if start < 0 || stop < -1 {
			return nil, errors.New("index out of range")
		} else if stop == -1 {
			stop = len(list)
		}
		return NewList(list[start:stop]), nil

	case []string:
		list := arr.list.([]string)
		if start > len(list) || stop > len(list) {
			return nil, errors.New("index out of range")
		} else if start < 0 || stop < -1 {
			return nil, errors.New("index out of range")
		} else if stop == -1 {
			stop = len(list)
		}
		return NewList(list[start:stop]), nil

	default:
		return nil, errors.New("interface type should be []float32, []float64, []int32, []int, []int64, or []string")
	}

}
