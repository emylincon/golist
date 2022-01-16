package golist

// List types
const (
	TypeListInt     = "golist.List[]int"
	TypeListInt32   = "golist.List[]int32"
	TypeListInt64   = "golist.List[]int64"
	TypeListFloat32 = "golist.List[]float32"
	TypeListFloat64 = "golist.List[]float64"
	TypeListString  = "golist.List[]string"
	TypeListUnknown = "golist.List[]unknown"
)

// returns string representation of the list.
func (arr *List) Type() (ltype string) {
	switch arr.list.(type) {
	case []int:
		return TypeListInt

	case []int32:
		return TypeListInt32

	case []int64:
		return TypeListInt64

	case []float32:
		return TypeListFloat32

	case []float64:
		return TypeListFloat64

	case []string:
		return TypeListString

	default:
		return TypeListUnknown
	}
}
