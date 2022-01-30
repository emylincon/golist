// Copyright 2021 Emeka Ugwuanyi. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package golist

import (
	"github.com/emylincon/golist/core"
)

// Lists :
// List interface and all methods
type Lists interface {
	List() interface{}                            // returns underlying slice
	Get(index int) interface{}                    // get item with index
	Append(element interface{})                   // append item
	Index(element interface{}) int                // get item's index
	Last() (interface{}, error)                   // get last item
	Max() (interface{}, error)                    // max item
	Min() (interface{}, error)                    // min item
	Type() string                                 // type of list
	Sum() interface{}                             // sum list items
	String() string                               // string repr
	LCM() (interface{}, error)                    // lcm of lest
	GCF() (interface{}, error)                    // gcf of list
	HCF() (interface{}, error)                    // same as GCF
	Clear()                                       // remove all elements in list
	Copy() (*List, error)                         // return a copy of list
	Count(element interface{}) int                // count how many times an item appears in list
	Extend(element interface{}) error             // extend list
	Pop(index int) interface{}                    // remove item
	Insert(element interface{}, index int) error  // insert item
	Rand() interface{}                            // return random item
	Remove(element interface{}) error             // remove item
	Replace(element interface{}, index int) error // replace item
	Reverse() *List                               // reverse list
	Set() (*List, error)                          // remove duplicates
	Slice(start int, stop int) (*List, error)     // return sub list
	Sort(reverse bool) interface{}                // sort list
	Add(other *List) *List                        // add 2 lists
	IsEqual(other *List) bool                     // checks 2 lists are equal
	Combinations(n int, joiner string) *List      // combinations of items in list
	CombinationsMax(n int, joiner string) *List   // combinations of items in list
	Contains(element interface{}) bool            // check if item in list
}

// List struct
type List struct {
	list interface{}
}

// NewList :
// list constructor
func NewList(list interface{}) *List {
	newlist := &List{
		list: list,
	}
	return newlist
}

// List :
// returns underlying slice or array
func (arr *List) List() interface{} {
	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		return list

	case []int32:
		list := arr.list.([]int32)
		return list

	case []int64:
		list := arr.list.([]int64)
		return list

	case []float32:
		list := arr.list.([]float32)
		return list

	case []float64:
		list := arr.list.([]float64)
		return list

	case []string:
		list := arr.list.([]string)
		return list

	default:
		return nil
	}
}

// Append :
// append element to list
func (arr *List) Append(element interface{}) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := element.(int)
		arr.list = *core.AppendInt(&list, item)

	case []int32:
		list := arr.list.([]int32)
		item := element.(int32)
		arr.list = *core.AppendInt32(&list, item)

	case []int64:
		list := arr.list.([]int64)
		item := element.(int64)
		arr.list = *core.AppendInt64(&list, item)

	case []float32:
		list := arr.list.([]float32)
		item := element.(float32)
		arr.list = *core.AppendFloat32(&list, item)

	case []float64:
		list := arr.list.([]float64)
		item := element.(float64)
		arr.list = *core.AppendFloat64(&list, item)

	case []string:
		list := arr.list.([]string)
		item := element.(string)
		arr.list = *core.AppendString(&list, item)

	default:
		return
	}

}

// Index :
// returns index of a given element. returns  -1 if element don't exist
func (arr *List) Index(element interface{}) int {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		item := element.(int)
		return core.IndexOfInt(&list, item)

	case []int32:
		list := arr.list.([]int32)
		item := element.(int32)
		return core.IndexOfInt32(&list, item)

	case []int64:
		list := arr.list.([]int64)
		item := element.(int64)
		return core.IndexOfInt64(&list, item)

	case []float32:
		list := arr.list.([]float32)
		item := element.(float32)
		return core.IndexOfFloat32(&list, item)

	case []float64:
		list := arr.list.([]float64)
		item := element.(float64)
		return core.IndexOfFloat64(&list, item)

	case []string:
		list := arr.list.([]string)
		item := element.(string)
		return core.IndexOfString(&list, item)

	default:
		return -1
	}

}

// Last :
// returns last element in the list
func (arr *List) Last() (interface{}, error) {

	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)
		return core.LastInt(list)

	case []int32:
		list := arr.list.([]int32)

		return core.LastInt32(list)

	case []int64:
		list := arr.list.([]int64)
		return core.LastInt64(list)

	case []float32:
		list := arr.list.([]float32)
		return core.LastFloat32(list)

	case []float64:
		list := arr.list.([]float64)
		return core.LastFloat64(list)

	case []string:
		list := arr.list.([]string)
		return core.LastString(list)

	default:
		return nil, ErrTypeNotsupported
	}

}
