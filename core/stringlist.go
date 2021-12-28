package core

import (
	"errors"
	"sort"
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
		return "", errors.New("list is empty")
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
		return nil, errors.New("element not in list")
	}
	newList, _ := PopString(list, index)
	return &newList, nil
}
