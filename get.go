package golist

// Get:
// returns element from list using index. if index dont exist returns nil
func (arr *List) Get(index int) interface{} {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		if index >= arr.Len() {
			return nil
		}
		return list[index]

	case []int32:
		list := arr.list.([]int32)
		if index >= arr.Len() {
			return nil
		}
		return list[index]

	case []int64:
		list := arr.list.([]int64)
		if index >= arr.Len() {
			return nil
		}
		return list[index]

	case []float32:
		list := arr.list.([]float32)
		if index >= arr.Len() {
			return nil
		}
		return list[index]

	case []float64:
		list := arr.list.([]float64)
		if index >= arr.Len() {
			return nil
		}
		return list[index]

	case []string:
		list := arr.list.([]string)
		if index >= arr.Len() {
			return nil
		}
		return list[index]

	default:
		return nil
	}

}
