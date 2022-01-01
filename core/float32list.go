package core

import (
	"errors"
	"sort"
)

func AppendFloat32(array *[]float32, element float32) *[]float32 {
	newArray := append(*array, element)
	return &newArray
}

func IndexOfFloat32(array *[]float32, element float32) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

func LastFloat32(list []float32) (float32, error) {
	if len(list) == 0 {
		return 0, errors.New("list is empty")
	}
	return list[len(list)-1], nil
}

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

func PopFloat32(list *[]float32, index int) ([]float32, float32) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

func CountFloat32(list *[]float32, element float32) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

func ExtendFloat32(list *[]float32, another []float32) []float32 {
	return append((*list), another...)
}

func InsertFloat32(list *[]float32, element float32, index int) *[]float32 {
	postpend := append([]float32{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

func RemoveFloat32(list *[]float32, element float32) (*[]float32, error) {
	index := IndexOfFloat32(list, element)
	if index == -1 {
		return nil, errors.New("element not in list")
	}
	newList, _ := PopFloat32(list, index)
	return &newList, nil
}

func ReverseFloat32(list *[]float32) *[]float32 {
	newList := []float32{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

func SumFloat32(list *[]float32) (sum float32) {
	for _, value := range *list {
		sum += value
	}
	return
}

func MaxFloat32(list *[]float32) (max float32) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}
