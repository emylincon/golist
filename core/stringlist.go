package core

import (
	"math/bits"
	"sort"
	"strings"
)

func AppendString(array *[]string, element string) *[]string {
	newArray := append(*array, element)
	return &newArray
}

func IndexOfString(array *[]string, element string) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

func LastString(array []string) (string, error) {
	if len(array) == 0 {
		return "", ErrListEmpty
	}
	return array[len(array)-1], nil
}

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

func PopString(list *[]string, index int) ([]string, string) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

func CountString(list *[]string, element string) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

func ExtendString(list *[]string, another []string) []string {
	return append((*list), another...)
}

func InsertString(list *[]string, element string, index int) *[]string {
	postpend := append([]string{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

func RemoveString(list *[]string, element string) (*[]string, error) {
	index := IndexOfString(list, element)
	if index == -1 {
		return nil, ErrNotInList
	}
	newList, _ := PopString(list, index)
	return &newList, nil
}

func ReverseString(list *[]string) *[]string {
	newList := []string{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

func SumString(list *[]string) (sum string) {
	sum = strings.Join(*list, " ")
	return
}

func MaxString(list *[]string) (max string) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

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
