package golist

// Copy  :
// returns a Copy of the list
func (arr *List) Copy() (*List, error) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		return NewList(list), nil

	case []int32:
		list := arr.list.([]int32)
		return NewList(list), nil

	case []int64:
		list := arr.list.([]int64)
		return NewList(list), nil

	case []float32:
		list := arr.list.([]float32)
		return NewList(list), nil

	case []float64:
		list := arr.list.([]float64)
		return NewList(list), nil

	case []string:
		list := arr.list.([]string)
		return NewList(list), nil

	default:
		return nil, ErrTypeNotsupported
	}

}
