![golist](img/golist.png)

<p align="center">
 <a style="text-decoration:none" href="https://img.shields.io/github/workflow/status/emylincon/golist/Go?style=for-the-badge" target="_blank">
     <img src="https://img.shields.io/github/workflow/status/emylincon/golist/Go?style=for-the-badge" alt="Build Status" />
 </a>
 <a style="text-decoration:none" href="https://img.shields.io/badge/Version-1.4.5-informational?style=flat-square" target="_blank">
     <img src="https://img.shields.io/badge/Version-1.4.5-informational?style=flat-square" alt="Version: 1.4.5" />
 </a>
 <a style="text-decoration:none" href="https://github.com/emylincon/golist/workflows/Go/badge.svg" target="_blank">
     <img src="https://github.com/emylincon/golist/workflows/Go/badge.svg" alt="Status" />
 </a>
 <a style="text-decoration:none" href="https://godoc.org/github.com/emylincon/golist" target="_blank">
     <img src="https://godoc.org/github.com/emylincon/golist?status.svg" alt="GoDoc" />
 </a>
 <a style="text-decoration:none" href="https://goreportcard.com/report/github.com/emylincon/golist" target="_blank">
     <img src="https://goreportcard.com/badge/github.com/emylincon/golist" alt="Go Report Card" />
 </a>
 </p>


# golist
A customized go list with index, sort, append, pop, count, clear and last item methods.
It supports all of the following data structures although the examples below are mostly `int`:
* `int`
* `int32`
* `int64`
* `float32`
* `float64`
* `string`

## To use module
* Import
```golang
import (
    "github.com/emylincon/golist"
)
```
* Download
```cmd
go get github.com/emylincon/golist
```

Here are all of the methods of the list objects:

# To Contribute
- [x] Download and install pre-commit: `pip install pre-commit`
- [x] Install precommit hook in repo: `pre-commit install`
- [x] before push run fmt: `gofmt -w .`

## `list.Get(i int) interface{}`
Get an item in the list by index. `i` represents the index. Returns `nil` if index don't exist.
```golang
list := golist.NewList([]int{1,2,3})
item := list.Get(0)
fmt.Println(item)  // 1
```

## `list.Index(x interface{}) int`
Get an item's index in the list. works in reverse of `list.Get(i)`. `x` represents the item. Returns `-1` if item don't exist.
```golang
list := golist.NewList([]int{1,2,3})
index := list.Index(2)
fmt.Println(index)  // 1
```

## `list.String()`
Returns a string representation of the object
```golang
list := golist.NewList([]int{3,2,1})
fmt.Println(list.String())  // [3, 2, 1]
```

## `list.Append(x interface{})`
Add an item to the end of the list. Items must be of the same type.
```golang
list := golist.NewList([]int{1,2,3})
list.Append(7)
fmt.Println(list)  // [1, 2, 3, 7]
```

## `list.Extend(slice interface{})`
Extend the list by appending all the items from a slice or array.
```golang
list := golist.NewList([]int{1,2,3})
list.Extend([]int{4,5})
fmt.Println(list)  // [1, 2, 3, 4, 5]
```

## `list.Insert(x interface{}, i int) error`
Insert an item at a given position. The first argument is the element while the second is the index to insert the element, so `list.insert(x, 0)` inserts at the front of the list, and `list.Insert(x, len(a))` is equivalent to `list.Append(x)`. Returns error is any
```golang
list := golist.NewList([]int{1, 2, 3})
err := list.Insert(4, 3)
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(list) // [1, 2, 3, 4]
```
* The above inserts item 4 to position 3 which is the end of the list

## `list.Remove(x interface{}) error`
Remove the first item from the list whose value is equal to x. It raises a ValueError if there is no such item.
```golang
list := golist.NewList([]int{1, 2, 3})
err := list.Remove(2)
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(list) // [1, 3]
```

## `list.Pop(i int) interface{}`
Remove the item at the given position in the list, and return it. `i` is the index of the element to be popped.
```golang
list := golist.NewList([]int{1, 2, 3})
popped := list.Pop(0)
fmt.Println(popped) // 1
fmt.Println(list)   // [2, 3]
```

## `list.Clear()`
Remove all items from the list.
```golang
list := golist.NewList([]int{1, 2, 3})
list.Clear()
fmt.Println(list) // []
```

## `list.Slice(start int, end int) (*golist.List, error)`
The arguments start and end are interpreted as in the slice notation and are used to return a particular subsequence of the list. The returned index is computed relative to the beginning of the full sequence rather than the start argument.
```golang
list := golist.NewList([]int{1, 2, 3, 2})
start, stop := 0, 2
NewList, err := list.Slice(start, stop)
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(NewList) // [1, 2]
```

## `list.Count(x interface{}) int`
Return the number of times x appears in the list. Returns -1 if element not found or the given element type does not match the type of list
```golang
list := golist.NewList([]int{1, 2, 3, 2})
count := list.Count(2)
fmt.Println(count) // 2
```

## `list.Sort(reverse bool)`
Sort the items of the list in place (the argument can be used for sort customization. `reverse` is `bool` so can be `true` or `false`.
```golang
list := golist.NewList([]int{3, 2, 1})
reverse := false
list.Sort(reverse)
fmt.Println(list) // [1, 2, 3]
```

## `list.Sorted(reverse bool) *golist.List`
Returns a list of Sorted items (the argument can be used for sort customization. `reverse` is `bool` so can be `true` or `false`).
```golang
list := golist.NewList([]int{3, 2, 1})
reverse := false
Newlist := list.Sorted(reverse)
fmt.Println(Newlist) // [1, 2, 3]
```

## `list.Reverse() *golist.List`
Returns new list with the elements of the list reversed.
```golang
list := golist.NewList([]int{5, 2, 7, 1})
newList := list.Reverse()
fmt.Println(newList) // [1, 7, 2, 5]
```

## `list.Copy() (*golist.List, error)`
Return a shallow copy of the list.
```golang
list := golist.NewList([]int{3, 2, 1})
copy, err := list.Copy()
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(copy) // [3, 2, 1]
```

## `list.Sum() interface{}`
Returns sum of the elements in the list. If it is list of string, it joins the strings with a space and returns it
```golang
list := golist.NewList([]int{3, 2, 1})
fmt.Println(list.Sum())  // 6
```

```golang
list := golist.NewList([]string{"Hello", "World"})
fmt.Println(list.Sum())  // "Hello World"
```

## `list.List() interface{}`
This is a getter that returns underlying slice interface.
```golang
list := golist.NewList([]int{3, 2, 1})
fmt.Println(list.List())  // [3 2 1]
```

## `list.Join(joiner string) string`
This only works with string data types, panics otherwise. `joiner` is a string used to join the list.
```golang
list := golist.NewList([]string{"Hello", "World"})
fmt.Println(list.Join("-"))  // "Hello-World"
```

## `list.Replace(x interface{}, i int) error`
Replaces an element at index i with element x. returns error if index does not exist. index of `-1` is equivalent to last item. This method is equivalent to working with slice (`a`) `a[1] = 10`
```golang
list := golist.NewList([]string{"Hello", "World"})
err := list.Replace("golang", -1)
if err != nil {
    fmt.Println(err)  // handle error
}
fmt.Println(list)  // ["Hello", "golang"]
```

## `list.Max() (interface{}, error)`
Returns max item in list. returns err if list is empty
```golang
list := golist.NewList([]string{"Hello", "World"})
max, err := list.Max()
if err != nil {
    fmt.Println(err)  // handles error
}
fmt.Println(max)  // "World"
```
```golang
list := golist.NewList([]int{3, 2, 1})
max, err := list.Max()
if err != nil {
    fmt.Println(err)  // handles error
}
fmt.Println(max)  // 3
```

## `list.Min() (interface{}, error)`
Returns min item in list. returns err if list is empty
```golang
list := golist.NewList([]string{"Hello", "World"})
min, err := list.Min()
if err != nil {
    fmt.Println(err)  // handles error
}
fmt.Println(min)  // "Hello"
```
```golang
list := golist.NewList([]int{3, 2, 1})
min, err := list.Min()
if err != nil {
    fmt.Println(err)  // handles error
}
fmt.Println(min)  // 1
```

## Loop through List
```golang
list := golist.NewList([]int{3, 2, 1})
for i := 0; i < list.Len(); i++ {
    fmt.Println(list.Get(i))
}
```
Output
```
3
2
1
```

## `list.GCF() (interface{}, error)`
Returns the Greatest Common Factor (GCF) or Highest Common Factor (HCF) of the numbers in the list. Only works with numbers. Returns error if called on list of strings. Uses [Euclidean algorithm](https://en.wikipedia.org/wiki/Euclidean_algorithm)
```golang
list := golist.NewList([]int{10, 15, 5})
gcf, err := list.GCF()
if err != nil {
    fmt.Println(err)  // handle error
}
fmt.Println(gcf)  // 5
```

## `list.LCM() (interface{}, error)`
Returns the [Least Common Multiple (LCM)](https://mathworld.wolfram.com/LeastCommonMultiple.html) of the numbers in the list. Only works with numbers. Returns error if called on list of strings. Uses [Euclidean algorithm](https://en.wikipedia.org/wiki/Euclidean_algorithm)
```golang
list := golist.NewList([]int{10, 15, 5})
lcm, err := list.LCM()
if err != nil {
    fmt.Println(err)  // handle error
}
fmt.Println(lcm)  // 30
```

## `list.Type() string`
returns the type of list
```golang
list := golist.NewList([]string{"Hello", "World"})
fmt.Println(list.Type())  // golist.List[]string
```

## `list.Rand() interface{}`
Returns a random element from list.
```golang
list := golist.NewList([]string{"Hello", "World"})
fmt.Println(list.Rand())  // World
```

## `list.Contains(element interface{}) bool`
returns true if element exists, returns false otherwise
```golang
list := golist.NewList([]string{"Hello", "World"})
fmt.Println(list.Contains("okay"))  // false
```

## `list.Combinations(n int, joiner string) (*golist.List, error)`
This is adapted from [Link](https://github.com/mxschmitt/golang-combinations). `joiner` is a string used to join the strings. Combinations returns combinations of n number of elements for a given string array.e.g if `n=2` it will return only 2 combined elements.
Furthermore `NewList([]string{"a", "b", "c"}).Combinations(2, "") = ["ab", "ac", "bc"]`.
* For `n < 1`, it equals to All and returns all combinations.
* For `n > len(list)` then `n = len(list)`
```golang
list := golist.NewList([]string{"a", "b", "c"})
combinedList := list.Combinations(2, " ")
fmt.Println(combinedList)  // ["a b", "a c", "b c"]
combinedList = list.Combinations(2, ",")
fmt.Println(combinedList)  // ["a,b", "a,c", "b,c"] <nil>
```

## `list.CombinationsMax(n int, joiner string) (*golist.List, error)`
Variation of `list.Combinations`. Difference is that for a given `n`, it returns combination lengths <= n, rather than only n.

```golang
list:= golist.NewList([]string{"a", "b", "c"})
fmt.Println(list.CombinationsMax(2, "")) // ["a", "b", "ab", "c", "ac", "bc"] <nil>
```

## `list.IsEqual(other *golist.List) bool`
returns true if both lists are equal, returns false otherwise
```golang
list := golist.NewList([]string{"Hello", "World"})
fmt.Println(list.IsEqual([]string{"a", "b"}))  // false
```

## `list.Set() (*golist.List, err)`
returns a new list with duplicates removed
```golang
list := golist.NewList([]int{1, 1, 1, 2, 3, 3, 4, 5, 6, 6})
fmt.Println(list.Set()) // [1, 2, 3, 4, 5, 6]
```

## `list.Add(other *golist.List) (golist.List, err)`
Adds two list together and returns a new list which is the result of the addition.
```golang
list := golist.NewList([]int{1, 0, 1})
other := golist.NewList([]int{0, 2, 0})
newList, err := list.Add(other)
if err != nil {
    log.Println(err) // handle error
}
fmt.Println(newList) // [1, 0, 1, 0, 2, 0]
```

## `list.ListSum(other *golist.List) (*golist.List, err)`
Add the content of two lists. The lists must be of the same type and have equal length. Example:
```golang
list1 := golist.NewList([]int{1,1})
list2 := golist.NewList([]int{2,2})
list3 := list1.ListSum(list2)
fmt.Println(list3) // [3,3]
```
## `list.ListSumNo(no interface{}) (*golist.List, err)`
Add number to all elements in list. Example
```golang
list1 := golist.NewList([]int{1,1})
no := 2
list3 := list1.ListSumNo(no)
fmt.Println(list3) // [3,3]
```

## `list.ConvertTo(t golist.ListType) (*golist.List, err)`
converts list from type `a` to type `b`. Example
```golang
list := golist.NewList([]int{1,1})
fmt.Println(list.Type()) // golist.List[]int
list.ConvertTo(golist.TypeListInt32)
fmt.Println(list.Type()) // golist.List[]int32
```

## `list.ListSubtract(other *golist.List) (*golist.List, err)`
Subtract the content of two lists. The lists must be of the same type and have equal length. Example:
```golang
list1 := golist.NewList([]int{1,1})
list2 := golist.NewList([]int{2,2})
list3, err := list1.ListSubtract(list2)
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(list3) // [-1, -1]
```

## `list.ListSubtractNo(no interface{}) (*golist.List, err)`
Subtract number from all elements in list. Example
```golang
list1 := golist.NewList([]int{1,1})
var no int = 2
list3, err := list1.ListSubtractNo(no)
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(list3) // [-1, -1]
```

## `list.ListMultiply(other *golist.List) (*golist.List, err)`
Multiply the content of two lists. The lists must be of the same type and have equal length. Example:
```golang
list1 := golist.NewList([]int{1,1})
list2 := golist.NewList([]int{2,2})
list3, err := list1.ListMultiply(list2)
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(list3) // [2, 2]
```

## `list.ListMultiplyNo(no interface{}) (*golist.List, err)`
Multiply a number with all elements in list. Example
```golang
list1 := golist.NewList([]int{1,1})
var no int = 2
list3, err := list1.ListMultiplyNo(no)
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(list3) // [2, 2]
```

## `list.ListDivide(other *golist.List) (*golist.List, err)`
Divide the content of two lists. The lists must be of the same type and have equal length. Example:
```golang
list1 := golist.NewList([]int{8,6})
list2 := golist.NewList([]int{2,2})
list3, err := list1.ListDivide(list2)
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(list3) // [4, 3]
```


## `list.ListDivideNo(no interface{}) (*golist.List, err)`
Divide all elements in list with no. Example
```golang
list1 := golist.NewList([]int{12,2})
var no int = 2
list3, err := list1.ListDivideNo(no)
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(list3) // [6, 1]
```

## `list.ConvertToSliceFloat32() ([]float32, error)`
Converts golist to []float32. Example
```golang
list := golist.NewList([]int{12,2})

slice, err := list.ConvertToSliceFloat32()
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Printf("%T", slice) // []float32
```

## `list.ConvertToSliceFloat64() ([]float64, error)`
Converts golist to []float64. Example
```golang
list := golist.NewList([]int{12,2})

slice, err := list.ConvertToSliceFloat64()
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Printf("%T", slice) // []float64
```

## `list.ConvertToSliceInt64() ([]int64, error)`
Converts golist to []int64. Example
```golang
list := golist.NewList([]int{12,2})

slice, err := list.ConvertToSliceInt64()
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Printf("%T", slice) // []int64
```

## `list.ConvertToSliceInt32() ([]int32, error)`
Converts golist to []int32. Example
```golang
list := golist.NewList([]int{12,2})

slice, err := list.ConvertToSliceInt32()
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Printf("%T", slice) // []int32
```

## `list.ConvertToSliceInt() ([]int, error)`
Converts golist to []int. Example
```golang
list := golist.NewList([]int{12,2})

slice, err := list.ConvertToSliceInt()
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Printf("%T", slice) // []int
```

## `list.ConvertToSliceString() ([]string, error)`
Converts golist to []string. Example
```golang
list := golist.NewList([]int{12,2})

slice, err := list.ConvertToSliceString()
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Printf("%T", slice) // []string
```

## `list.Difference(other *List) (*List, error)`
Difference returns the elements in `list` that aren't in `other`. Example
```golang
list := golist.NewList([]int{1,2,3,4})
other := golist.NewList([]int{3,4,5})
diff, err := list.Difference(other)
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(diff) // [1, 2]
```

## `list.DifferenceBoth(other *List) (*List, error)`
DifferenceBoth returns the elements that aren't in both lists. Example
```golang
list := golist.NewList([]int{1,2,3,4})
other := golist.NewList([]int{3,4,5})
diff, err := list.DifferenceBoth(other)
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(diff) // [1, 2, 5]
```
