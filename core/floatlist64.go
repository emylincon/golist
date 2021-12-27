package core

import (
	"errors"
	"sort"
)

func AppendFloat64(array *[]float64, element float64) *[]float64 {
	newArray := append(*array, element)
	return &newArray
}

func IndexOfFloat64(array *[]float64, element float64) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

func LastFloat64(list []float64) (float64, error) {
	if len(list) == 0 {
		return 0, errors.New("list is empty")
	}
	return list[len(list)-1], nil
}

func SortFloat64(list *[]float64) *[]float64 {
	ref_list := *list
	sort.Float64s(ref_list)
	return &ref_list
}

func PopFloat64(list *[]float64, index int) ([]float64, float64) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

func CountFloat64(list *[]float64, element float64) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

func ExtendFloat64(list *[]float64, another []float64) []float64 {
	return append((*list), another...)
}
