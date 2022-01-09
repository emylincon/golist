package golist

import "errors"

var (
	ErrTypeNotsupported    = errors.New("interface type should be []float32, []float64, []int32, []int, []int64, or []string")
	ErrStringsNotsupported = errors.New("strings are not supported for this operation")
	ErrIndexOutOfRange     = errors.New("index error: list index out of range")
)
