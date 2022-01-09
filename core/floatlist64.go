package core

import (
	"errors"
	"math"
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

func SumFloat64(list *[]float64) (sum float64) {
	for _, value := range *list {
		sum += value
	}
	return
}

func MaxFloat64(list *[]float64) (max float64) {
	for i, value := range *list {
		if i == 0 || value > max {
			max = value
		}
	}
	return
}

func MinFloat64(list *[]float64) (min float64) {
	for i, value := range *list {
		if i == 0 || value < min {
			min = value
		}
	}
	return
}

func rounder(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func RoundFloat64(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(rounder(num*output)) / output
}

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

func GCFfloat64(list *[]float64) (gcf float64, err error) {
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
			gcf = _gcfFloat64(a, b)
		}
	}
	return gcf, nil
}
