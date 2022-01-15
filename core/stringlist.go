package core

import (
	"math/bits"
	"sort"
	"strings"
)

// implements append
func AppendString(array *[]string, element string) *[]string {
	newArray := append(*array, element)
	return &newArray
}

// returns index of element. returns -1 if element dont exist
func IndexOfString(array *[]string, element string) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

// return last element in slice
func LastString(array []string) (string, error) {
	if len(array) == 0 {
		return "", ErrListEmpty
	}
	return array[len(array)-1], nil
}

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

// removes element from slice
func PopString(list *[]string, index int) ([]string, string) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

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

// adds two slice together
func ExtendString(list *[]string, another []string) []string {
	return append((*list), another...)
}

// insert an element in a given position
func InsertString(list *[]string, element string, index int) *[]string {
	postpend := append([]string{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

// removes given element from slice
func RemoveString(list *[]string, element string) (*[]string, error) {
	index := IndexOfString(list, element)
	if index == -1 {
		return nil, ErrNotInList
	}
	newList, _ := PopString(list, index)
	return &newList, nil
}

// reverse the positions of elements in slice
func ReverseString(list *[]string) *[]string {
	newList := []string{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

// joins elements in slice
func SumString(list *[]string) (sum string) {
	sum = strings.Join(*list, " ")
	return
}

// returns max element in slice
func MaxString(list *[]string) (max string) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

// returns min element in slice
func MinString(list *[]string) (min string) {
	for i, value := range *list {
		if i == 0 || value < min {
			min = value
		}
	}
	return
}

// adapted from https://github.com/mxschmitt/golang-combinations
func CombinationsString(set []string, n int) (subsets []string) {
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
				subset += set[object]
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

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
