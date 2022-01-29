// Copyright 2021 Emeka Ugwuanyi. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package golist_test

import (
	"fmt"

	"github.com/emylincon/golist"
)

func ExampleNewList() {
	list := golist.NewList([]int{1, 2, 3})
	fmt.Println(list)
	// Output:
	// [1, 2, 3]
}
