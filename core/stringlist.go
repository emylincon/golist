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

func SortString(list *[]string) *[]string {
	ref_list := *list
	sort.Strings(ref_list)
	return &ref_list
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
