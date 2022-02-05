package golist

// Len  :
// returns length of list
func (arr *List) Len() int {
	switch list := arr.list.(type) {
	case []int:
		return len(list)

	case []int32:
		return len(list)

	case []int64:
		return len(list)

	case []float32:
		return len(list)

	case []float64:
		return len(list)

	case []string:
		return len(list)

	default:
		return -1
	}
}
