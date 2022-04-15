package core

import (
	"sort"
	"strconv"
)

// AppendInt64 :
// implements append
func AppendInt64(list []int64, element int64) []int64 {
	newArray := append(list, element)
	return newArray
}

// IndexOfInt64 :
// returns index of element. returns -1 if element dont exist
func IndexOfInt64(list *[]int64, element int64) int {
	for index, value := range *list {
		if value == element {
			return index
		}
	}
	return -1
}

// LastInt64 :
// return last element in slice
func LastInt64(list []int64) (int64, error) {
	if len(list) == 0 {
		return 0, ErrListEmpty
	}
	return list[len(list)-1], nil
}

// SortInt64 :
// sorts slice
func SortInt64(list *[]int64, reverse bool) *[]int64 {
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

// PopInt64 :
// removes element from slice
func PopInt64(list *[]int64, index int) ([]int64, int64) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

// CountInt64 :
// returns the number of times a given element appears in a slice
func CountInt64(list *[]int64, element int64) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

// ExtendInt64 :
// adds two slice together
func ExtendInt64(list *[]int64, another []int64) []int64 {
	return append((*list), another...)
}

// InsertInt64 :
// insert an element in a given position
func InsertInt64(list *[]int64, element int64, index int) *[]int64 {
	postpend := append([]int64{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

// RemoveInt64 :
// removes given element from slice
func RemoveInt64(list *[]int64, element int64) (*[]int64, error) {
	index := IndexOfInt64(list, element)
	if index == -1 {
		return nil, ErrNotInList
	}
	newList, _ := PopInt64(list, index)
	return &newList, nil
}

// ReverseInt64 :
// reverse the positions of elements in slice
func ReverseInt64(list *[]int64) *[]int64 {
	newList := []int64{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

// SumInt64 :
// sum of elements in slice
func SumInt64(list *[]int64) (sum int64) {
	for _, value := range *list {
		sum += value
	}
	return
}

// MaxInt64 :
// returns max element in slice
func MaxInt64(list *[]int64) (max int64) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

// MinInt64 :
// returns min element in slice
func MinInt64(list *[]int64) (min int64) {
	for i, value := range *list {
		if i == 0 || value < min {
			min = value
		}
	}
	return
}

// helper function for hcf and lcm
func lcmHcfInt64(list *[]int64, get func(a, b int64) int64) (result int64, err error) {
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
func _gcfInt64(a, b int64) int64 {
	var h, l int64
	if a > b {
		h, l = a, b
	} else {
		h, l = b, a
	}

	for {
		r := h % l
		if r == 0 {
			return l
		}
		h = l
		l = r

	}

}

// GCFInt64 :
// returns greatest common factor of slice
func GCFInt64(list *[]int64) (gcf int64, err error) {
	return lcmHcfInt64(list, _gcfInt64)
}

// helper function for lcm
func _lcmInt64(a, b int64) int64 {
	if a == 0 && b == 0 {
		return 0
	}
	return (a * b) / _gcfInt64(a, b)
}

// LCMInt64 :
// returns lowest common multiple of slice
func LCMInt64(list *[]int64) (lcm int64, err error) {
	return lcmHcfInt64(list, _lcmInt64)
}

// SetInt64 :
// returns a set of slice i.e removes duplicates
func SetInt64(list []int64) (set []int64) {
	keys := map[int64]bool{}
	for _, key := range list {
		if !keys[key] {
			keys[key] = true
			set = append(set, key)
		}
	}
	return
}

// ListSumInt64 sums contents of two lists
func ListSumInt64(list []int64, other []int64) (sum []int64) {
	for i, v := range list {
		sum = append(sum, v+other[i])
	}
	return
}

// ListSumNoInt64 sums contents of two lists
func ListSumNoInt64(list []int64, no int64) (sum []int64) {
	for _, v := range list {
		sum = append(sum, v+no)
	}
	return
}

// ListSubtractInt64 subtracts the contents of two lists
func ListSubtractInt64(list []int64, other []int64) (sum []int64) {
	for i, v := range list {
		sum = append(sum, v-other[i])
	}
	return
}

// ListSubtractNoInt64 subtracts an element from the contents a list
func ListSubtractNoInt64(list []int64, no int64) (sum []int64) {
	for _, v := range list {
		sum = append(sum, v-no)
	}
	return
}

// ListMultiplyInt64 multiply contents of two lists
func ListMultiplyInt64(list []int64, other []int64) (product []int64) {
	for i, v := range list {
		product = append(product, v*other[i])
	}
	return
}

// ListMultiplyNoInt64 multiply a given number with all elements in list
func ListMultiplyNoInt64(list []int64, no int64) (product []int64) {
	for _, v := range list {
		product = append(product, v*no)
	}
	return
}

// ListDivideInt64 divide the contents of two lists
func ListDivideInt64(list []int64, other []int64) (product []int64) {
	for i, v := range list {
		product = append(product, v/other[i])
	}
	return
}

// ListDivideNoInt64 divide a given number with all elements in list
func ListDivideNoInt64(list []int64, no int64) (product []int64) {
	for _, v := range list {
		product = append(product, v/no)
	}
	return
}

// ConvertToInt64 converts to slice to int64
func ConvertToInt64(array interface{}) (Intify []int64, err error) {

	switch array := array.(type) {
	case []int:
		for _, v := range array {
			Intify = append(Intify, int64(v))
		}

	case []int32:
		for _, v := range array {
			Intify = append(Intify, int64(v))
		}

	case []int64:
		return array, nil

	case []float32:
		for _, v := range array {
			Intify = append(Intify, int64(v))
		}

	case []float64:
		for _, v := range array {
			Intify = append(Intify, int64(v))
		}

	case []string:
		for _, v := range array {
			intVar, err := strconv.ParseInt(v, 0, 64)
			if err != nil {
				return nil, err
			}
			Intify = append(Intify, int64(intVar))
		}

	default:
		return nil, ErrTypeNotSupported

	}
	return
}
