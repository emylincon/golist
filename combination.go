package golist

import (
	"github.com/emylincon/golist/core"
)

// adapted from https://github.com/mxschmitt/golang-combinations
// Combinations returns combinations of n number of elements for a given string array.
// e.g if n=2 it will return only 2 combined elements :=
// futhermore NewList([]string{"a", "b", "c"}).Combinations(2) = ["ab", "ac", "bc"]
// For n < 1, it equals to All and returns all combinations.
// for n > len(list); n = len(list)
func (arr *List) Combinations(n int) (*List, error) {

	switch arr.list.(type) {

	case []string:
		list := arr.list.([]string)
		set := core.SetString(list)
		list = core.CombinationsString(set, n)
		return NewList(list), nil

	default:
		return nil, ErrTypeNotsupported
	}

}
