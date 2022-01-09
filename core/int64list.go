package core

import (
	"errors"
	"sort"
)

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

func SortInt64(list *[]int64, reverse bool) *[]int64 {
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

func InsertInt64(list *[]int64, element int64, index int) *[]int64 {
	postpend := append([]int64{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

func RemoveInt64(list *[]int64, element int64) (*[]int64, error) {
	index := IndexOfInt64(list, element)
	if index == -1 {
		return nil, errors.New("element not in list")
	}
	newList, _ := PopInt64(list, index)
	return &newList, nil
}

func ReverseInt64(list *[]int64) *[]int64 {
	newList := []int64{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

func SumInt64(list *[]int64) (sum int64) {
	for _, value := range *list {
		sum += value
	}
	return
}

func MaxInt64(list *[]int64) (max int64) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

func MinInt64(list *[]int64) (min int64) {
	for i, value := range *list {
		if i == 0 || value < min {
			min = value
		}
	}
	return
}

func _gcfInt64(a, b int64) int64 {
	var h, l int64
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

func GCFInt64(list *[]int64) (gcf int64, err error) {
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
			gcf = _gcfInt64(a, b)
		}
	}
	return gcf, nil
}
