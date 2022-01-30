package core

import (
	"sort"
)

// AppendInt:
// implements append
func AppendInt(array *[]int, element int) *[]int {
	newArray := append(*array, element)
	return &newArray
}

// IndexOfInt:
// returns index of element. returns -1 if element dont exist
func IndexOfInt(array *[]int, element int) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

// LastInt:
// return last element in slice
func LastInt(array []int) (int, error) {
	if len(array) == 0 {
		return 0, ErrListEmpty
	}
	return array[len(array)-1], nil
}

// SortInt:
// sorts slice
func SortInt(list *[]int, reverse bool) *[]int {
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

// PopInt:
// removes element from slice
func PopInt(list *[]int, index int) ([]int, int) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

// CountInt:
// returns the number of times a given element appears in a slice
func CountInt(list *[]int, element int) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

// ExtendInt:
// adds two slice together
func ExtendInt(list *[]int, another []int) []int {
	return append((*list), another...)
}

// InsertInt:
// insert an element in a given position
func InsertInt(list *[]int, element int, index int) *[]int {
	postpend := append([]int{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

// RemoveInt:
// removes given element from slice
func RemoveInt(list *[]int, element int) (*[]int, error) {
	index := IndexOfInt(list, element)
	if index == -1 {
		return nil, ErrNotInList
	}
	newList, _ := PopInt(list, index)
	return &newList, nil
}

// ReverseInt:
// reverse the positions of elements in slice
func ReverseInt(list *[]int) *[]int {
	newList := []int{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

// SumInt:
// sum of elements in slice
func SumInt(list *[]int) (sum int) {
	for _, value := range *list {
		sum += value
	}
	return
}

// MaxInt:
// returns max element in slice
func MaxInt(list *[]int) (max int) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

// MinInt:
// returns min element in slice
func MinInt(list *[]int) (min int) {
	for i, value := range *list {
		if i == 0 || value < min {
			min = value
		}
	}
	return
}

// helper function for hcf and lcm
func lcmHcfInt(list *[]int, get func(a, b int) int) (result int, err error) {
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

// helper function for gcf
func _gcfInt(a, b int) int {
	var h, l int
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

// GCFInt:
func GCFInt(list *[]int) (gcf int, err error) {
	return lcmHcfInt(list, _gcfInt)
}

// helper function for lcm
func _lcmInt(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}
	return (a * b) / _gcfInt(a, b)
}

// LCMInt:
// returns lowest common multiple of slice
func LCMInt(list *[]int) (lcm int, err error) {
	return lcmHcfInt(list, _lcmInt)
}

// SetInt:
// returns a set of slice i.e removes duplicates
func SetInt(list []int) (set []int) {
	keys := map[int]bool{}
	for _, key := range list {
		if !keys[key] {
			keys[key] = true
			set = append(set, key)
		}
	}
	return
}

// ListSumInt sums contents of two lists
func ListSumInt(list []int, other []int) (sum []int) {
	for i, v := range list {
		sum = append(sum, v+other[i])
	}
	return
}

// ListSumNoInt sums contents of two lists
func ListSumNoInt(list []int, no int) (sum []int) {
	for _, v := range list {
		sum = append(sum, v+no)
	}
	return
}

// ListSubtractInt sums contents of two lists
func ListSubtractInt(list []int, other []int) (sum []int) {
	for i, v := range list {
		sum = append(sum, v-other[i])
	}
	return
}

// ListSubtractNoInt sums contents of two lists
func ListSubtractNoInt(list []int, no int) (sum []int) {
	for _, v := range list {
		sum = append(sum, v-no)
	}
	return
}

// ConvertToInt converts to slice to int
func ConvertToInt(array interface{}) (Intify []int, err error) {

	switch array := array.(type) {
	case []int:
		return array, nil

	case []int32:
		for _, v := range array {
			Intify = append(Intify, int(v))
		}

	case []int64:
		for _, v := range array {
			Intify = append(Intify, int(v))
		}

	case []float32:
		for _, v := range array {
			Intify = append(Intify, int(v))
		}

	case []float64:
		for _, v := range array {
			Intify = append(Intify, int(v))
		}

	case []string:
		return nil, ErrTypeNotSupported

	default:
		return nil, ErrTypeNotSupported

	}
	return
}
