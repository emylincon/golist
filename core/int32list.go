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

func SortInt32(list *[]int32, reverse bool) *[]int32 {
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

func InsertInt32(list *[]int32, element int32, index int) *[]int32 {
	postpend := append([]int32{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

func RemoveInt32(list *[]int32, element int32) (*[]int32, error) {
	index := IndexOfInt32(list, element)
	if index == -1 {
		return nil, errors.New("element not in list")
	}
	newList, _ := PopInt32(list, index)
	return &newList, nil
}

func ReverseInt32(list *[]int32) *[]int32 {
	newList := []int32{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

func SumInt32(list *[]int32) (sum int32) {
	for _, value := range *list {
		sum += value
	}
	return
}

func MaxInt32(list *[]int32) (max int32) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

func MinInt32(list *[]int32) (min int32) {
	for i, value := range *list {
		if i == 0 || value < min {
			min = value
		}
	}
	return
}

func _gcfInt32(a, b int32) int32 {
	var h, l int32
	if a > b {
		h, l = a, b
	} else {
		h, l = b, a
	}

	for {
		r := h % l
		if r == 0 {
			return l
		} else {
			h = l
			l = r
		}
	}

}

func GCFInt32(list *[]int32) (gcf int32, err error) {
	var count int
	if len(*list) < 2 {
		return (*list)[0], nil
	}
	count = len(*list)
	gcf = (*list)[0]
	for i := 1; i < count; i++ {
		b := (*list)[i]
		a := gcf
		if a < 0 || b < 0 {
			return 0, ErrNotZeroOrPositive
		} else if a == 0 || b == 0 {
			gcf = a + b
		} else {
			gcf = _gcfInt32(a, b)
		}
	}
	return gcf, nil
}
