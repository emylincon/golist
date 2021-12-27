package golist

func (arr *List) Clear() {
	switch arr.list.(type) {
	case []int:
		arr.list = []int{}

	case []int32:
		arr.list = []int32{}

	case []int64:
		arr.list = []int64{}

	case []float32:
		arr.list = []float32{}

	case []float64:
		arr.list = []float64{}

	case []string:
		arr.list = []string{}

	default:
		return
	}
}
