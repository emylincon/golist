package core

import "errors"

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

// TODO
func SortFloat32(list *[]float32) *[]float32 {
	// ref_list := *list
	// sort.Float32s(ref_list)
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
