package golist

func (arr *List) Type() (ltype string) {
	switch arr.list.(type) {
	case []int:
		ltype = "golist.List[]int"

	case []int32:
		ltype = "golist.List[]int32"

	case []int64:
		ltype = "golist.List[]int64"

	case []float32:
		ltype = "golist.List[]float32"

	case []float64:
		ltype = "golist.List[]float64"

	case []string:
		ltype = "golist.List[]string"

	default:
		ltype = "golist.List[]unknown"
	}
	return
}
