package golist

func (arr *List) Replace(element interface{}, index int) error {
	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := element.(int)
		if index >= len(list) || index < -1 {
			return ErrIndexOutOfRange
		} else if index == -1 {
			index = len(list) - 1
		}
		list[index] = item
		arr.list = list
		return nil
	case []int32:
		list := arr.list.([]int32)
		item := element.(int32)
		if index >= len(list) || index < -1 {
			return ErrIndexOutOfRange
		} else if index == -1 {
			index = len(list) - 1
		}
		list[index] = item
		arr.list = list
		return nil
	case []int64:
		list := arr.list.([]int64)
		item := element.(int64)
		if index >= len(list) || index < -1 {
			return ErrIndexOutOfRange
		} else if index == -1 {
			index = len(list) - 1
		}
		list[index] = item
		arr.list = list
		return nil
	case []float32:
		list := arr.list.([]float32)
		item := element.(float32)
		if index >= len(list) || index < -1 {
			return ErrIndexOutOfRange
		} else if index == -1 {
			index = len(list) - 1
		}
		list[index] = item
		arr.list = list
		return nil
	case []float64:
		list := arr.list.([]float64)
		item := element.(float64)
		if index >= len(list) || index < -1 {
			return ErrIndexOutOfRange
		} else if index == -1 {
			index = len(list) - 1
		}
		list[index] = item
		arr.list = list
		return nil
	case []string:
		list := arr.list.([]string)
		item := element.(string)
		if index >= len(list) || index < -1 {
			return ErrIndexOutOfRange
		} else if index == -1 {
			index = len(list) - 1
		}
		list[index] = item
		arr.list = list
		return nil
	default:
		return ErrTypeNotsupported
	}
}
