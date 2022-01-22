package golist

func validateIndex(index, listLen int) (int, error) {
	if index >= listLen || index < -1 {
		return index, ErrIndexOutOfRange
	} else if index == -1 {
		index = listLen - 1
		return index, nil
	} else {
		return index, nil
	}
}

// replaces an element in the lsit using the elements index. returns error if index do not exist.
func (arr *List) Replace(element interface{}, index int) error {
	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := element.(int)
		index, err := validateIndex(index, len(list))
		if err != nil {
			return err
		}
		list[index] = item
		arr.list = list
		return nil
	case []int32:
		list := arr.list.([]int32)
		item := element.(int32)
		index, err := validateIndex(index, len(list))
		if err != nil {
			return err
		}
		list[index] = item
		arr.list = list
		return nil
	case []int64:
		list := arr.list.([]int64)
		item := element.(int64)
		index, err := validateIndex(index, len(list))
		if err != nil {
			return err
		}
		list[index] = item
		arr.list = list
		return nil
	case []float32:
		list := arr.list.([]float32)
		item := element.(float32)
		index, err := validateIndex(index, len(list))
		if err != nil {
			return err
		}
		list[index] = item
		arr.list = list
		return nil
	case []float64:
		list := arr.list.([]float64)
		item := element.(float64)
		index, err := validateIndex(index, len(list))
		if err != nil {
			return err
		}
		list[index] = item
		arr.list = list
		return nil
	case []string:
		list := arr.list.([]string)
		item := element.(string)
		index, err := validateIndex(index, len(list))
		if err != nil {
			return err
		}
		list[index] = item
		arr.list = list
		return nil
	default:
		return ErrTypeNotsupported
	}
}
