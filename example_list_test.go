// Copyright 2021 Emeka Ugwuanyi. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package golist_test

import (
	"fmt"

	"github.com/emylincon/golist"
)

func Example() {
	list := golist.NewList([]int{1, 2, 3})
	// Get an item in the list by index. `i` represents the index. Returns `nil` if index don't exist.
	item := list.Get(0)
	fmt.Println(item)

	// To return an item's index in the list, use list.index. works in reverse of `list.Get(i)`. `x` represents the item. Returns `-1` if item don't exist.
	index := list.Index(2)
	fmt.Println(index)

	// Append an item to the end of the list. Items must be of the same type.
	list.Append(7)
	fmt.Println(list)

	// Extend the list by appending all the items from a slice or array.
	list.Extend([]int{4, 5})
	fmt.Println(list)

	// Reverse elements in list.
	fmt.Println(list.Reverse())

	// sum elements in list
	fmt.Println(list.Sum())

	// Output:
	// 1
	// 1
	// [1, 2, 3, 7]
	// [1, 2, 3, 7, 4, 5]
	// [5, 4, 7, 3, 2, 1]
	// 22
}
