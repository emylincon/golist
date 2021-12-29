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

func PopFloat64(list *[]float64, index int) ([]float64, float64) {
	listD := *list
	popped := listD[index]
	newList := append(listD[:index], listD[index+1:]...)
	return newList, popped
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

func InsertFloat64(list *[]float64, element float64, index int) *[]float64 {
	postpend := append([]float64{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

func RemoveFloat64(list *[]float64, element float64) (*[]float64, error) {
	index := IndexOfFloat64(list, element)
	if index == -1 {
		return nil, errors.New("element not in list")
	}
	newList, _ := PopFloat64(list, index)
	return &newList, nil
}

func ReverseFloat64(list *[]float64) *[]float64 {
	newList := []float64{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}
