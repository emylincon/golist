package golist

import "github.com/emylincon/golist/core"

// ConvertToSliceFloat32 converts golist to []float32
func (arr *List) ConvertToSliceFloat32() ([]float32, error) {
	return core.ConvertToFloat32(arr.list)
}

// ConvertToSliceFloat64 converts golist to []float64
func (arr *List) ConvertToSliceFloat64() ([]float64, error) {
	return core.ConvertToFloat64(arr.list)
}

// ConvertToSliceInt64 converts golist to []int64
func (arr *List) ConvertToSliceInt64() ([]int64, error) {
	return core.ConvertToInt64(arr.list)
}

// ConvertToSliceInt32 converts golist to []int32
func (arr *List) ConvertToSliceInt32() ([]int32, error) {
	return core.ConvertToInt32(arr.list)
}

// ConvertToSliceInt converts golist to []int
func (arr *List) ConvertToSliceInt() ([]int, error) {
	return core.ConvertToInt(arr.list)
}

// ConvertToSliceString converts golist to []string
func (arr *List) ConvertToSliceString() ([]string, error) {
	return core.ConvertToString(arr.list)
}
