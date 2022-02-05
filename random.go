package golist

import (
	"math/rand"
	"time"
)

// Rand  :
// ## list.Rand() interface{}
// Returns a random element from list.
//
// ```golang
// list := golist.NewList([]string{"Hello", "World"})
// fmt.Println(list.Rand())  // World
// ```
func (arr *List) Rand() interface{} {
	rand.Seed(time.Now().Unix())
	switch list := arr.list.(type) {
	case []int:
		return list[rand.Intn(arr.Len())]

	case []int32:
		return list[rand.Intn(arr.Len())]

	case []int64:
		return list[rand.Intn(arr.Len())]

	case []float32:
		return list[rand.Intn(arr.Len())]

	case []float64:
		return list[rand.Intn(arr.Len())]

	case []string:
		return list[rand.Intn(arr.Len())]

	default:
		return ErrTypeNotsupported
	}

}
