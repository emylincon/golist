package core

import (
	"fmt"
	"math/bits"
	"sort"
	"strings"
)

// AppendString:
// implements append
func AppendString(array *[]string, element string) *[]string {
	newArray := append(*array, element)
	return &newArray
}

// IndexOfString:
// returns index of element. returns -1 if element dont exist
func IndexOfString(array *[]string, element string) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

// LastString:
// return last element in slice
func LastString(array []string) (string, error) {
	if len(array) == 0 {
		return "", ErrListEmpty
	}
	return array[len(array)-1], nil
}

// SortString:
// sorts slice
func SortString(list *[]string, reverse bool) *[]string {
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

// PopString:
// removes element from slice
func PopString(list *[]string, index int) ([]string, string) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

// CountString:
// returns the number of times a given element appears in a slice
func CountString(list *[]string, element string) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

// ExtendString:
// adds two slice together
func ExtendString(list *[]string, another []string) []string {
	return append((*list), another...)
}

// InsertString:
// insert an element in a given position
func InsertString(list *[]string, element string, index int) *[]string {
	postpend := append([]string{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

// RemoveString:
// removes given element from slice
func RemoveString(list *[]string, element string) (*[]string, error) {
	index := IndexOfString(list, element)
	if index == -1 {
		return nil, ErrNotInList
	}
	newList, _ := PopString(list, index)
	return &newList, nil
}

// ReverseString:
// reverse the positions of elements in slice
func ReverseString(list *[]string) *[]string {
	newList := []string{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

// SumString:
// joins elements in slice
func SumString(list *[]string) (sum string) {
	sum = strings.Join(*list, " ")
	return
}

// MaxString:
// returns max element in slice
func MaxString(list *[]string) (max string) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

// MinString:
// returns min element in slice
func MinString(list *[]string) (min string) {
	for i, value := range *list {
		if i == 0 || value < min {
			min = value
		}
	}
	return
}

// CombinationsString:
// adapted from https://github.com/mxschmitt/golang-combinations
func CombinationsString(set []string, n int, joiner string) (subsets []string) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset += set[object] + joiner
			}
		}
		// remove unwanted joiner and add subset to subsets
		if len(joiner) != 0 {
			subset = subset[:len(subset)-1]
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

// CombinationsStringMax: should return combinations of max length.
// For example: a list of [a, b, c] with max combination 2, should return 1 and 2 combinations.
// i.e combinations length <= n
// adapted from https://github.com/mxschmitt/golang-combinations
func CombinationsStringMax(set []string, n int, joiner string) (subsets []string) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) > n {
			continue
		}

		var subset string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset += set[object] + joiner
			}
		}
		// remove unwanted joiner and add subset to subsets
		if len(joiner) != 0 {
			subset = subset[:len(subset)-1]
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

// SetString
// returns a set of slice i.e removes duplicates
func SetString(list []string) (set []string) {
	keys := map[string]bool{}
	for _, key := range list {
		if !keys[key] {
			keys[key] = true
			set = append(set, key)
		}
	}
	return
}

func ConvertToString(array interface{}) (stringfy []string, err error) {

	switch array := array.(type) {
	case []int:
		for _, v := range array {
			stringfy = append(stringfy, fmt.Sprintf("%v", v))
		}

	case []int32:
		for _, v := range array {
			stringfy = append(stringfy, fmt.Sprintf("%v", v))
		}

	case []int64:
		for _, v := range array {
			stringfy = append(stringfy, fmt.Sprintf("%v", v))
		}

	case []float32:
		for _, v := range array {
			stringfy = append(stringfy, fmt.Sprintf("%v", v))
		}

	case []float64:
		for _, v := range array {
			stringfy = append(stringfy, fmt.Sprintf("%v", v))
		}

	case []string:
		return array, nil

	default:
		return nil, ErrTypeNotSupported

	}
	return
}
