package golist

import "github.com/emylincon/golist/core"

// List types
const (
	TypeListInt     ListType = "golist.List[]int"
	TypeListInt32   ListType = "golist.List[]int32"
	TypeListInt64   ListType = "golist.List[]int64"
	TypeListFloat32 ListType = "golist.List[]float32"
	TypeListFloat64 ListType = "golist.List[]float64"
	TypeListString  ListType = "golist.List[]string"
	TypeListUnknown ListType = "golist.List[]unknown"
)

// ListType is a string representation of list type
type ListType string

// Type returns string representation of the list.
func (arr *List) Type() ListType {
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

// ConvertTo converts list to a new type
func (arr *List) ConvertTo(t ListType) (*List, error) {
	switch t {
	case TypeListInt:
		list, err := core.ConvertToInt(arr.list)
		if err != nil {
			return &List{}, err
		}
		return NewList(list), nil

	case TypeListInt32:
		list, err := core.ConvertToInt32(arr.list)
		if err != nil {
			return &List{}, err
		}
		return NewList(list), nil

	case TypeListInt64:
		list, err := core.ConvertToInt64(arr.list)
		if err != nil {
			return &List{}, err
		}
		return NewList(list), nil

	case TypeListFloat32:
		list, err := core.ConvertToFloat32(arr.list)
		if err != nil {
			return &List{}, err
		}
		return NewList(list), nil

	case TypeListFloat64:
		list, err := core.ConvertToFloat64(arr.list)
		if err != nil {
			return &List{}, err
		}
		return NewList(list), nil

	case TypeListString:
		list, err := core.ConvertToString(arr.list)
		if err != nil {
			return &List{}, err
		}
		return NewList(list), nil

	default:
		return &List{}, ErrTypeNotsupported
	}
}
