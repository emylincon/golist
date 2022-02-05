package golist

// Copy  :
// returns a Copy of the list
func (arr *List) Copy() (*List, error) {

	switch list := arr.list.(type) {
	case []int:
		return NewList(list), nil

	case []int32:
		return NewList(list), nil

	case []int64:
		return NewList(list), nil

	case []float32:
		return NewList(list), nil

	case []float64:
		return NewList(list), nil

	case []string:
		return NewList(list), nil

	default:
		return nil, ErrTypeNotsupported
	}

}
