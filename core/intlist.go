package core

import (
	"sort"
)

// implements append
func AppendInt(array *[]int, element int) *[]int {
	newArray := append(*array, element)
	return &newArray
}

// returns index of element. returns -1 if element dont exist
func IndexOfInt(array *[]int, element int) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

// return last element in slice
func LastInt(array []int) (int, error) {
	if len(array) == 0 {
		return 0, ErrListEmpty
	}
	return array[len(array)-1], nil
}

// sorts slice
func SortInt(list *[]int, reverse bool) *[]int {
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

// removes element from slice
func PopInt(list *[]int, index int) ([]int, int) {
	listD := *list
	popped := listD[index]
	newArray := append(listD[:index], listD[index+1:]...)
	return newArray, popped
}

// returns the number of times a given element appears in a slice
func CountInt(list *[]int, element int) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

// adds two slice together
func ExtendInt(list *[]int, another []int) []int {
	return append((*list), another...)
}

// insert an element in a given position
func InsertInt(list *[]int, element int, index int) *[]int {
	postpend := append([]int{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

// removes given element from slice
func RemoveInt(list *[]int, element int) (*[]int, error) {
	index := IndexOfInt(list, element)
	if index == -1 {
		return nil, ErrNotInList
	}
	newList, _ := PopInt(list, index)
	return &newList, nil
}

// reverse the positions of elements in slice
func ReverseInt(list *[]int) *[]int {
	newList := []int{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

// sum of elements in slice
func SumInt(list *[]int) (sum int) {
	for _, value := range *list {
		sum += value
	}
	return
}

// returns max element in slice
func MaxInt(list *[]int) (max int) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

// returns min element in slice
func MinInt(list *[]int) (min int) {
	for i, value := range *list {
		if i == 0 || value < min {
			min = value
		}
	}
	return
}

// helper function for hcf and lcm
func lcmHcfInt(list *[]int, get func(a, b int) int) (result int, err error) {
	var count int
	if len(*list) < 2 {
		return (*list)[0], nil
	}
	count = len(*list)
	result = (*list)[0]
	for i := 1; i < count; i++ {
		b := (*list)[i]
		a := result
		if a < 0 || b < 0 {
			return 0, ErrNotZeroOrPositive
		} else if a == 0 || b == 0 {
			result = a + b
		} else {
			result = get(a, b)
		}
	}
	return result, nil
}

// helper function for gcf
func _gcfInt(a, b int) int {
	var h, l int
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

func GCFInt(list *[]int) (gcf int, err error) {
	return lcmHcfInt(list, _gcfInt)
}

// helper function for lcm
func _lcmInt(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}
	return (a * b) / _gcfInt(a, b)
}

// returns lowest common multiple of slice
func LCMInt(list *[]int) (lcm int, err error) {
	return lcmHcfInt(list, _lcmInt)
}

// returns a set of slice i.e removes duplicates
func SetInt(list []int) (set []int) {
	keys := map[int]bool{}
	for _, key := range list {
		if !keys[key] {
			keys[key] = true
			set = append(set, key)
		}
	}
	return
}
