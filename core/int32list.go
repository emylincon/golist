package core

import (
	"errors"
	"sort"
)

func AppendInt32(array *[]int32, element int32) *[]int32 {
	newArray := append(*array, element)
	return &newArray
}

func IndexOfInt32(array *[]int32, element int32) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

func LastInt32(list []int32) (int32, error) {
	if len(list) == 0 {
		return 0, errors.New("list is empty")
	}
	return list[len(list)-1], nil
}

func SortInt32(list *[]int32) *[]int32 {
	sort.SliceStable(*list, func(i, j int) bool {
		return (*list)[i] < (*list)[j]
	})
	return list
}

func PopInt32(list *[]int32, index int) ([]int32, int32) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

func CountInt32(list *[]int32, element int32) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

func ExtendInt32(list *[]int32, another []int32) []int32 {
	return append((*list), another...)
}
