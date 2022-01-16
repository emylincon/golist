package golist

import "errors"

// Error variables used within the module
var (
	ErrTypeNotsupported    = errors.New("type error: interface type should be []float32, []float64, []int32, []int, []int64, or []string")
	ErrStringsNotsupported = errors.New("operation error: strings are not supported for this operation")
	ErrIndexOutOfRange     = errors.New("index error: list index out of range")
	ErrListsNotOfSameType  = errors.New("type error: lists not of same type")
)
