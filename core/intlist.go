package core

import (
	"errors"
	"sort"
)

func AppendInt(array *[]int, element int) *[]int {
	newArray := append(*array, element)
	return &newArray
}

func IndexOfInt(array *[]int, element int) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

func LastInt(array []int) (int, error) {
	if len(array) == 0 {
		return 0, errors.New("list is empty")
	}
	return array[len(array)-1], nil
}

func SortInt(list *[]int) *[]int {
	ref_list := *list
	sort.Ints(ref_list)
	return &ref_list
}

func PopInt(list *[]int, index int) ([]int, int) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

func CountInt(list *[]int, element int) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

func ExtendInt(list *[]int, another []int) []int {
	return append((*list), another...)
}
