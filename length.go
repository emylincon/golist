package golist

// returns length of list
func (arr *List) Len() int {
	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		return len(list)

	case []int32:
		list := arr.list.([]int32)
		return len(list)

	case []int64:
		list := arr.list.([]int64)
		return len(list)

	case []float32:
		list := arr.list.([]float32)
		return len(list)

	case []float64:
		list := arr.list.([]float64)
		return len(list)

	case []string:
		list := arr.list.([]string)
		return len(list)

	default:
		return -1
	}
}
