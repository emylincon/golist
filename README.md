![golist](img/golist.png)

<p align="center">
<a style="text-decoration:none" href="https://img.shields.io/github/workflow/status/emylincon/golist/Go?style=for-the-badge" target="_blank">
    <img src="https://img.shields.io/github/workflow/status/emylincon/golist/Go?style=for-the-badge" alt="Build Status" />
</a>
</p>

![Version: 1.3.0](https://img.shields.io/badge/Version-1.3.0-informational?style=flat-square)  ![status](https://github.com/emylincon/golist/workflows/Go/badge.svg) [![GoDoc](https://godoc.org/github.com/emylincon/golist?status.svg)](https://godoc.org/github.com/emylincon/golist) [![Go Report Card](https://goreportcard.com/badge/github.com/emylincon/golist)](https://goreportcard.com/report/github.com/emylincon/golist)


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
index := list.index(2)
fmt.Println(index)  // 1
```

## `list.String()`
Returns a string representation of the object
```golang
list := NewList([]int{3,2,1})
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

## `list.Sort(reverse bool) interface{}`
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
list := NewList([]string{"a", "b", "c"})
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
list3 := list1.SumList(list2)
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

## `list.SumListNo(no interface{}) (*golist.List, err)`
Add number to all elements in list. Example
```golang
list1 := NewList([]int{1,1})
no := 2
list3 := list1.SumListNo(no)
fmt.Println(list3) // [3,3]
```
