package core

import (
	"sort"
)

// AppendFloat32:
// implements append
func AppendFloat32(array *[]float32, element float32) *[]float32 {
	newArray := append(*array, element)
	return &newArray
}

// IndexOfFloat32:
// returns index of element. returns -1 if element dont exist
func IndexOfFloat32(array *[]float32, element float32) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

// LastFloat32:
// return last element in slice
func LastFloat32(list []float32) (float32, error) {
	if len(list) == 0 {
		return 0, ErrListEmpty
	}
	return list[len(list)-1], nil
}

// SortFloat32:
// sorts slice
func SortFloat32(list *[]float32, reverse bool) *[]float32 {
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

// PopFloat32:
// removes element from slice
func PopFloat32(list *[]float32, index int) ([]float32, float32) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

// CountFloat32:
// returns the number of times a given element appears in a slice
func CountFloat32(list *[]float32, element float32) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

// ExtendFloat32:
// adds two slice together
func ExtendFloat32(list *[]float32, another []float32) []float32 {
	return append((*list), another...)
}

// InsertFloat32:
// insert an element in a given position
func InsertFloat32(list *[]float32, element float32, index int) *[]float32 {
	postpend := append([]float32{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

// RemoveFloat32:
// removes given element from slice
func RemoveFloat32(list *[]float32, element float32) (*[]float32, error) {
	index := IndexOfFloat32(list, element)
	if index == -1 {
		return nil, ErrNotInList
	}
	newList, _ := PopFloat32(list, index)
	return &newList, nil
}

// ReverseFloat32:
// reverse the positions of elements in slice
func ReverseFloat32(list *[]float32) *[]float32 {
	newList := []float32{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

// SumFloat32:
// sum of elements in slice
func SumFloat32(list *[]float32) (sum float32) {
	for _, value := range *list {
		sum += value
	}
	return
}

// MaxFloat32:
// returns max element in slice
func MaxFloat32(list *[]float32) (max float32) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

// MinFloat32:
// returns min element in slice
func MinFloat32(list *[]float32) (min float32) {
	for i, value := range *list {
		if i == 0 || value < min {
			min = value
		}
	}
	return
}

// helper function for hcf and lcm
func lcmHcfFloat32(list *[]float32, get func(a, b float64) float64) (result float32, err error) {
	var count int
	if len(*list) < 2 {
		return (*list)[0], nil
	}
	count = len(*list)
	result = (*list)[0]
	for i := 1; i < count; i++ {
		b := float64((*list)[i])
		a := float64(result)
		if a < 0 || b < 0 {
			return 0, ErrNotZeroOrPositive
		} else if a == 0 || b == 0 {
			result = float32(a + b)
		} else {
			result = float32(get(a, b))
		}
	}
	return result, nil
}

// GCFFloat32:
// returns greatest common factor of slice
func GCFFloat32(list *[]float32) (gcf float32, err error) {
	return lcmHcfFloat32(list, _gcfFloat64)
}

// LCMFloat32:
// returns lowest common multiple of slice
func LCMFloat32(list *[]float32) (lcm float32, err error) {
	return lcmHcfFloat32(list, _lcmFloat64)
}

// SetFloat32:
// returns a set of slice i.e removes duplicates
func SetFloat32(list []float32) (set []float32) {
	keys := map[float32]bool{}
	for _, key := range list {
		if !keys[key] {
			keys[key] = true
			set = append(set, key)
		}
	}
	return
}

// ListSumInt sums contents of two lists
func ListSumFloat32(list []float32, other []float32) (sum []float32) {
	for i, v := range list {
		sum = append(sum, v+other[i])
	}
	return
}

// ConvertToFloat32 converts to slice to float32
func ConvertToFloat32(array interface{}) (Intify []float32, err error) {

	switch array := array.(type) {
	case []int:
		for _, v := range array {
			Intify = append(Intify, float32(v))
		}

	case []int32:
		for _, v := range array {
			Intify = append(Intify, float32(v))
		}

	case []int64:
		for _, v := range array {
			Intify = append(Intify, float32(v))
		}

	case []float32:
		return array, nil

	case []float64:
		for _, v := range array {
			Intify = append(Intify, float32(v))
		}

	case []string:
		return nil, ErrTypeNotSupported

	default:
		return nil, ErrTypeNotSupported

	}
	return
}
