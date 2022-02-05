package golist

// Get  :
// returns element from list using index. if index dont exist returns nil
func (arr *List) Get(index int) interface{} {

	switch list := arr.list.(type) {
	case []int:
		if index >= arr.Len() {
			return nil
		}
		return list[index]

	case []int32:
		if index >= arr.Len() {
			return nil
		}
		return list[index]

	case []int64:
		if index >= arr.Len() {
			return nil
		}
		return list[index]

	case []float32:
		if index >= arr.Len() {
			return nil
		}
		return list[index]

	case []float64:
		if index >= arr.Len() {
			return nil
		}
		return list[index]

	case []string:
		if index >= arr.Len() {
			return nil
		}
		return list[index]

	default:
		return nil
	}

}
