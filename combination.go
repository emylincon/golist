package golist

import (
	"github.com/emylincon/golist/core"
)

// adapted from https://github.com/mxschmitt/golang-combinations
// Combinations returns Combinations of n number of elements for a given string array.
// e.g if n=2 it will return only 2 combined elements :=
// furthermore NewList([]string{"a", "b", "c"}).Combinations(2, "") = ["ab", "ac", "bc"]
// For n < 1, it equals to All and returns all Combinations.
// for n > len(list); n = len(list)
func (arr *List) Combinations(n int, joiner string) (*List, error) {

	switch arr.list.(type) {

	case []string:
		list := arr.list.([]string)
		set := core.SetString(list)
		list = core.CombinationsString(set, n, joiner)
		return NewList(list), nil

	default:
		return nil, ErrTypeNotsupported
	}

}

// adapted from https://github.com/mxschmitt/golang-combinations
// CombinationsMax returns combinations of n number of elements for a given string array.
// e.g if n=2 it will return combinations <= 2
// furthermore NewList([]string{"a", "b", "c"}).CombinationsMax(2, "") = ["a", "b", "c", "ab", "ac", "bc"]
// For n < 1, it equals to All and returns all combinations.
// for n > len(list); n = len(list)
func (arr *List) CombinationsMax(n int, joiner string) (*List, error) {

	switch arr.list.(type) {

	case []string:
		list := arr.list.([]string)
		set := core.SetString(list)
		list = core.CombinationsStringMax(set, n, joiner)
		return NewList(list), nil

	default:
		return nil, ErrTypeNotsupported
	}

}
