package core

import "errors"

func AppendInt64(list *[]int64, element int64) *[]int64 {
	newArray := append(*list, element)
	return &newArray
}

func IndexOfInt64(list *[]int64, element int64) int {
	for index, value := range *list {
		if value == element {
			return index
		}
	}
	return -1
}

func LastInt64(list []int64) (int64, error) {
	if len(list) == 0 {
		return 0, errors.New("list is empty")
	}
	return list[len(list)-1], nil
}

// TODO
func SortInt64(list *[]int64) *[]int64 {
	//ref_list := *list
	//sort.Ints(ref_list)
	return list
}

func PopInt64(list *[]int64, index int) ([]int64, int64) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

func CountInt64(list *[]int64, element int64) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

func ExtendInt64(list *[]int64, another []int64) []int64 {
	return append((*list), another...)
}
