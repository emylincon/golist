package core

import (
	"math"
	"sort"
)

// implements append
func AppendFloat64(array *[]float64, element float64) *[]float64 {
	newArray := append(*array, element)
	return &newArray
}

// returns index of element. returns -1 if element dont exist
func IndexOfFloat64(array *[]float64, element float64) int {
	for index, value := range *array {
		if value == element {
			return index
		}
	}
	return -1
}

// return last element in slice
func LastFloat64(list []float64) (float64, error) {
	if len(list) == 0 {
		return 0, ErrListEmpty
	}
	return list[len(list)-1], nil
}

// sorts slice
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

// removes element from slice
func PopFloat64(list *[]float64, index int) ([]float64, float64) {
	listD := *list
	popped := listD[index]
	newList := append(listD[:index], listD[index+1:]...)
	return newList, popped
}

// returns the number of times a given element appears in a slice
func CountFloat64(list *[]float64, element float64) int {
	count := 0
	for _, value := range *list {
		if value == element {
			count++
		}
	}
	return count
}

// adds two slice together
func ExtendFloat64(list *[]float64, another []float64) []float64 {
	return append((*list), another...)
}

// insert an element in a given position
func InsertFloat64(list *[]float64, element float64, index int) *[]float64 {
	postpend := append([]float64{element}, (*list)[index:]...)
	newList := append((*list)[:index], postpend...)
	return &newList
}

// removes given element from slice
func RemoveFloat64(list *[]float64, element float64) (*[]float64, error) {
	index := IndexOfFloat64(list, element)
	if index == -1 {
		return nil, ErrNotInList
	}
	newList, _ := PopFloat64(list, index)
	return &newList, nil
}

// reverse the positions of elements in slice
func ReverseFloat64(list *[]float64) *[]float64 {
	newList := []float64{}
	for i := len(*list) - 1; i >= 0; i-- {
		newList = append(newList, (*list)[i])
	}

	return &newList
}

// sum of elements in slice
func SumFloat64(list *[]float64) (sum float64) {
	for _, value := range *list {
		sum += value
	}
	return
}

// returns max element in slice
func MaxFloat64(list *[]float64) (max float64) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

// returns min element in slice
func MinFloat64(list *[]float64) (min float64) {
	for i, value := range *list {
		if i == 0 || value < min {
			min = value
		}
	}
	return
}

// helper function for RoundFloat64(num float64, precision int) float64
func rounder(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// rounds float with a precision
func RoundFloat64(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(rounder(num*output)) / output
}

// helper function for gcf
func _gcfFloat64(a, b float64) float64 {
	var h, l float64
	if a > b {
		h, l = a, b
	} else {
		h, l = b, a
	}

	for {
		r := RoundFloat64(math.Mod(h, l), 4)
		if r == 0 {
			return l
		} else {
			h = l
			l = r
		}
	}

}

// helper function for hcf and lcm
func lcmHcfFloat64(list *[]float64, get func(a, b float64) float64) (result float64, err error) {
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
func GCFFloat64(list *[]float64) (gcf float64, err error) {
	return lcmHcfFloat64(list, _gcfFloat64)
}

// helper function for lcm
func _lcmFloat64(a, b float64) float64 {
	if a == 0 && b == 0 {
		return 0
	}
	return (a * b) / _gcfFloat64(a, b)
}

// returns lowest common multiple of slice
func LCMFloat64(list *[]float64) (lcm float64, err error) {
	return lcmHcfFloat64(list, _lcmFloat64)
}

// returns a set of slice i.e removes duplicates
func SetFloat64(list []float64) (set []float64) {
	keys := map[float64]bool{}
	for _, key := range list {
		if !keys[key] {
			keys[key] = true
			set = append(set, key)
		}
	}
	return
}
