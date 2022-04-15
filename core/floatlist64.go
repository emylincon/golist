package core

import (
	"math"
	"sort"
	"strconv"
)

// AppendFloat64 :
// implements append
func AppendFloat64(array []float64, element float64) []float64 {
	newArray := append(array, element)
	return newArray
}

// IndexOfFloat64 :
// returns index of element. returns -1 if element dont exist
func IndexOfFloat64(array *[]float64, element float64) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

// LastFloat64 :
// return last element in slice
func LastFloat64(list []float64) (float64, error) {
	if len(list) == 0 {
		return 0, ErrListEmpty
	}
	return list[len(list)-1], nil
}

// SortFloat64 :
// sorts slice
func SortFloat64(list *[]float64, reverse bool) *[]float64 {
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

// PopFloat64 :
// removes element from slice
func PopFloat64(list *[]float64, index int) ([]float64, float64) {
	listD := *list
	popped := listD[index]
	newList := append(listD[:index], listD[index+1:]...)
	return newList, popped
}

// CountFloat64 :
// returns the number of times a given element appears in a slice
func CountFloat64(list *[]float64, element float64) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

// ExtendFloat64 :
// adds two slice together
func ExtendFloat64(list *[]float64, another []float64) []float64 {
	return append((*list), another...)
}

// InsertFloat64 :
// insert an element in a given position
func InsertFloat64(list *[]float64, element float64, index int) *[]float64 {
	postpend := append([]float64{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

// RemoveFloat64 :
// removes given element from slice
func RemoveFloat64(list *[]float64, element float64) (*[]float64, error) {
	index := IndexOfFloat64(list, element)
	if index == -1 {
		return nil, ErrNotInList
	}
	newList, _ := PopFloat64(list, index)
	return &newList, nil
}

// ReverseFloat64 :
// reverse the positions of elements in slice
func ReverseFloat64(list *[]float64) *[]float64 {
	newList := []float64{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

// SumFloat64 :
// sum of elements in slice
func SumFloat64(list *[]float64) (sum float64) {
	for _, value := range *list {
		sum += value
	}
	return
}

// MaxFloat64 :
// returns max element in slice
func MaxFloat64(list *[]float64) (max float64) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

// MinFloat64 :
// returns min element in slice
func MinFloat64(list *[]float64) (min float64) {
	for i, value := range *list {
		if i == 0 || value < min {
			min = value
		}
	}
	return
}

// helper function for RoundFloat64(num float64, precision int) float64
func rounder(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// RoundFloat64 :
// rounds float with a precision
func RoundFloat64(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(rounder(num*output)) / output
}

// helper function for gcf
func _gcfFloat64(a, b float64) float64 {
	var h, l float64
	if a > b {
		h, l = a, b
	} else {
		h, l = b, a
	}

	for {
		r := RoundFloat64(math.Mod(h, l), 4)
		if r == 0 {
			return l
		}
		h = l
		l = r
	}

}

// helper function for hcf and lcm
func lcmHcfFloat64(list *[]float64, get func(a, b float64) float64) (result float64, err error) {
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

// GCFFloat64 :
// helper function for gcf
func GCFFloat64(list *[]float64) (gcf float64, err error) {
	return lcmHcfFloat64(list, _gcfFloat64)
}

// helper function for lcm
func _lcmFloat64(a, b float64) float64 {
	if a == 0 && b == 0 {
		return 0
	}
	return (a * b) / _gcfFloat64(a, b)
}

// LCMFloat64 :
// returns lowest common multiple of slice
func LCMFloat64(list *[]float64) (lcm float64, err error) {
	return lcmHcfFloat64(list, _lcmFloat64)
}

// SetFloat64 :
// returns a set of slice i.e removes duplicates
func SetFloat64(list []float64) (set []float64) {
	keys := map[float64]bool{}
	for _, key := range list {
		if !keys[key] {
			keys[key] = true
			set = append(set, key)
		}
	}
	return
}

// ListSumFloat64 sums contents of two lists
func ListSumFloat64(list []float64, other []float64) (sum []float64) {
	for i, v := range list {
		sum = append(sum, v+other[i])
	}
	return
}

// ListSumNoFloat64 sums contents of two lists
func ListSumNoFloat64(list []float64, no float64) (sum []float64) {
	for _, v := range list {
		sum = append(sum, v+no)
	}
	return
}

// ListSubtractFloat64 subtracts the contents of two lists
func ListSubtractFloat64(list []float64, other []float64) (sum []float64) {
	for i, v := range list {
		sum = append(sum, v-other[i])
	}
	return
}

// ListSubtractNoFloat64 subtracts an element from the contents a list
func ListSubtractNoFloat64(list []float64, no float64) (sum []float64) {
	for _, v := range list {
		sum = append(sum, v-no)
	}
	return
}

// ListMultiplyFloat64 multiply contents of two lists
func ListMultiplyFloat64(list []float64, other []float64) (product []float64) {
	for i, v := range list {
		product = append(product, v*other[i])
	}
	return
}

// ListMultiplyNoFloat64 multiply a given number with all elements in list
func ListMultiplyNoFloat64(list []float64, no float64) (product []float64) {
	for _, v := range list {
		product = append(product, v*no)
	}
	return
}

// ListDivideFloat64 divide the contents of two lists
func ListDivideFloat64(list []float64, other []float64) (product []float64) {
	for i, v := range list {
		product = append(product, v/other[i])
	}
	return
}

// ListDivideNoFloat64 divide a given number with all elements in list
func ListDivideNoFloat64(list []float64, no float64) (product []float64) {
	for _, v := range list {
		product = append(product, v/no)
	}
	return
}

// ConvertToFloat64 converts to slice to float64
func ConvertToFloat64(array interface{}) (Intify []float64, err error) {

	switch array := array.(type) {
	case []int:
		for _, v := range array {
			Intify = append(Intify, float64(v))
		}

	case []int32:
		for _, v := range array {
			Intify = append(Intify, float64(v))
		}

	case []int64:
		for _, v := range array {
			Intify = append(Intify, float64(v))
		}

	case []float32:
		for _, v := range array {
			Intify = append(Intify, float64(v))
		}

	case []float64:
		return array, nil

	case []string:
		for _, v := range array {
			value, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, err
			}
			Intify = append(Intify, float64(value))
		}

	default:
		return nil, ErrTypeNotSupported

	}
	return
}
