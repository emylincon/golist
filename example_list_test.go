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
	fmt.Println("Get(0) :", item)

	// To return an item's index in the list, use list.index. works in reverse of `list.Get(i)`. `x` represents the item. Returns `-1` if item don't exist.
	index := list.Index(2)
	fmt.Println("Index(2) :", index)

	// Append an item to the end of the list. Items must be of the same type.
	list.Append(7)
	fmt.Println("Append(7) :", list)

	// Extend the list by appending all the items from a slice or array.
	list.Extend([]int{4, 5})
	fmt.Println("Extend([]int{4, 5) :", list)

	// Reverse elements in list.
	fmt.Println("Reverse() :", list.Reverse())

	// sum elements in list
	fmt.Println("Sum() :", list.Sum())

	// remove a given element from list
	err := list.Remove(7)
	if err != nil {
		fmt.Println(err) // handle error
	}
	fmt.Println("Remove(7) :", list)

	// Subtract the content of two lists.
	list1 := golist.NewList([]int{1, 1})
	list2 := golist.NewList([]int{2, 2})
	list3, err := list1.ListSubtract(list2)
	if err != nil {
		fmt.Println(err) // handle error
	}
	fmt.Println("ListSubtract :", list3)

	// Output:
	// Get(0) : 1
	// Index(2) : 1
	// Append(7) : [1, 2, 3, 7]
	// Extend([]int{4, 5) : [1, 2, 3, 7, 4, 5]
	// Reverse() : [5, 4, 7, 3, 2, 1]
	// Sum() : 22
	// Remove(7) : [1, 2, 3, 4, 5]
	// ListSubtract : [-1, -1]
}
