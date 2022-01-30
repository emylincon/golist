package core

import (
	"sort"
)

// AppendInt32 :
// implements append
func AppendInt32(array *[]int32, element int32) *[]int32 {
	newArray := append(*array, element)
	return &newArray
}

// IndexOfInt32 :
// returns index of element. returns -1 if element dont exist
func IndexOfInt32(array *[]int32, element int32) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

// LastInt32 :
// return last element in slice
func LastInt32(list []int32) (int32, error) {
	if len(list) == 0 {
		return 0, ErrListEmpty
	}
	return list[len(list)-1], nil
}

// SortInt32 :
// sorts slice
func SortInt32(list *[]int32, reverse bool) *[]int32 {
	if reverse {
		sort.SliceStable(*list, func(i, j int) bool {
			return (*list)[i] > (*list)[j]
		})
	} else {
		sort.SliceStable(*list, func(i, j int) bool {
			return (*list)[i] < (*list)[j]
		})
	}
	return list
}

// PopInt32 :
// removes element from slice
func PopInt32(list *[]int32, index int) ([]int32, int32) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

// CountInt32 :
// returns the number of times a given element appears in a slice
func CountInt32(list *[]int32, element int32) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

// ExtendInt32 :
// adds two slice together
func ExtendInt32(list *[]int32, another []int32) []int32 {
	return append((*list), another...)
}

// InsertInt32 :
// insert an element in a given position
func InsertInt32(list *[]int32, element int32, index int) *[]int32 {
	postpend := append([]int32{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

// RemoveInt32 :
// removes given element from slice
func RemoveInt32(list *[]int32, element int32) (*[]int32, error) {
	index := IndexOfInt32(list, element)
	if index == -1 {
		return nil, ErrNotInList
	}
	newList, _ := PopInt32(list, index)
	return &newList, nil
}

// ReverseInt32 :
// reverse the positions of elements in slice
func ReverseInt32(list *[]int32) *[]int32 {
	newList := []int32{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

// SumInt32 :
// sum of elements in slice
func SumInt32(list *[]int32) (sum int32) {
	for _, value := range *list {
		sum += value
	}
	return
}

// MaxInt32 :
// returns max element in slice
func MaxInt32(list *[]int32) (max int32) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

// MinInt32 :
// returns min element in slice
func MinInt32(list *[]int32) (min int32) {
	for i, value := range *list {
		if i == 0 || value < min {
			min = value
		}
	}
	return
}

// helper function for hcf and lcm
func lcmHcfInt32(list *[]int32, get func(a, b int32) int32) (result int32, err error) {
	var count int
	if len(*list) < 2 {
		return (*list)[0], nil
	}
	count = len(*list)
	result = (*list)[0]
	for i := 1; i < count; i++ {
		b := (*list)[i]
		a := result
		if a < 0 || b < 0 {
			return 0, ErrNotZeroOrPositive
		} else if a == 0 || b == 0 {
			result = a + b
		} else {
			result = get(a, b)
		}
	}
	return result, nil
}

// helper function for hcf
func _gcfInt32(a, b int32) int32 {
	var h, l int32
	if a > b {
		h, l = a, b
	} else {
		h, l = b, a
	}

	for {
		r := h % l
		if r == 0 {
			return l
		} else {
			h = l
			l = r
		}
	}

}

// GCFInt32 :
// returns greatest common factor of slice
func GCFInt32(list *[]int32) (gcf int32, err error) {
	return lcmHcfInt32(list, _gcfInt32)
}

// helper function for lcm
func _lcmInt32(a, b int32) int32 {
	if a == 0 && b == 0 {
		return 0
	}
	return (a * b) / _gcfInt32(a, b)
}

// LCMInt32 :
// returns lowest common multiple of slice
func LCMInt32(list *[]int32) (lcm int32, err error) {
	return lcmHcfInt32(list, _lcmInt32)
}

// SetInt32 :
// returns a set of slice i.e removes duplicates
func SetInt32(list []int32) (set []int32) {
	keys := map[int32]bool{}
	for _, key := range list {
		if !keys[key] {
			keys[key] = true
			set = append(set, key)
		}
	}
	return
}

// ListSumInt sums contents of two lists
func ListSumInt32(list []int32, other []int32) (sum []int32) {
	for i, v := range list {
		sum = append(sum, v+other[i])
	}
	return
}

// ListSumNoInt sums contents of two lists
func ListSumNoInt32(list []int32, no int32) (sum []int32) {
	for _, v := range list {
		sum = append(sum, v+no)
	}
	return
}

// ConvertToInt32 converts to slice to int32
func ConvertToInt32(array interface{}) (Intify []int32, err error) {

	switch array := array.(type) {
	case []int:
		for _, v := range array {
			Intify = append(Intify, int32(v))
		}

	case []int32:
		return array, nil

	case []int64:
		for _, v := range array {
			Intify = append(Intify, int32(v))
		}

	case []float32:
		for _, v := range array {
			Intify = append(Intify, int32(v))
		}

	case []float64:
		for _, v := range array {
			Intify = append(Intify, int32(v))
		}

	case []string:
		return nil, ErrTypeNotSupported

	default:
		return nil, ErrTypeNotSupported

	}
	return
}
