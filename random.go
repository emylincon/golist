package golist

import (
	"math/rand"
	"time"
)

// Rand :
// ## list.Rand() interface{}
// Returns a random element from list.
//
// ```golang
// list := golist.NewList([]string{"Hello", "World"})
// fmt.Println(list.Rand())  // World
// ```
func (arr *List) Rand() interface{} {
	rand.Seed(time.Now().Unix())
	switch arr.list.(type) {
	case []int:
		list := arr.list.([]int)

		return list[rand.Intn(arr.Len())]

	case []int32:
		list := arr.list.([]int32)
		return list[rand.Intn(arr.Len())]

	case []int64:
		list := arr.list.([]int64)
		return list[rand.Intn(arr.Len())]

	case []float32:
		list := arr.list.([]float32)
		return list[rand.Intn(arr.Len())]

	case []float64:
		list := arr.list.([]float64)
		return list[rand.Intn(arr.Len())]

	case []string:
		list := arr.list.([]string)
		return list[rand.Intn(arr.Len())]

	default:
		return ErrTypeNotsupported
	}

}
