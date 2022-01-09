![GitHub Workflow Status](https://img.shields.io/github/workflow/status/emylincon/golist/Go?style=for-the-badge) ![Version: 1.1.4](https://img.shields.io/badge/Version-1.1.4-informational?style=flat-square)  ![status](https://github.com/emylincon/golist/workflows/Go/badge.svg)
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

## list.Get(i) interface{}
Get an item in the list by index. `i` represents the index. Returns `nil` if index don't exist.
```golang
list := golist.NewList([]int{1,2,3})
item := list.Get(0)
fmt.Println(item)  // 1
```
## list.String()
Returns a string representation of the object
```golang
list := NewList([]int{3,2,1})
fmt.Println(list.String())  // [3, 2, 1]
```

## list.Append(x)
Add an item to the end of the list. Items must be of the same type.
```golang
list := golist.NewList([]int{1,2,3})
list.Append(7)
fmt.Println(list)  // [1, 2, 3, 7]
```

## list.Extend(slice)
Extend the list by appending all the items from a slice or array.
```golang
list := golist.NewList([]int{1,2,3})
list.Extend([]int{4,5})
fmt.Println(list)  // [1, 2, 3, 4, 5]
```

## list.Insert(x, i) error
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

## list.Remove(x) error
Remove the first item from the list whose value is equal to x. It raises a ValueError if there is no such item.
```golang
list := golist.NewList([]int{1, 2, 3})
err := list.Remove(2)
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(list) // [1, 3]
```

## list.Pop(i) interface{}
Remove the item at the given position in the list, and return it. i is the index of the element to be popped.
```golang
list := golist.NewList([]int{1, 2, 3})
popped := list.Pop(0)
fmt.Println(popped) // 1
fmt.Println(list)   // [2, 3]
```

## list.Clear()
Remove all items from the list.
```golang
list := golist.NewList([]int{1, 2, 3})
list.Clear()
fmt.Println(list) // []
```

## list.Slice(start, end) (*golist.List, error)
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

## list.Count(x) interface{}
Return the number of times x appears in the list.
```golang
list := golist.NewList([]int{1, 2, 3, 2})
count := list.Count(2)
fmt.Println(count) // 2
```

## list.Sort(reverse) interface{}
Sort the items of the list in place (the argument can be used for sort customization. `reverse` is `bool` so can be `true` or `false`.
```golang
list := golist.NewList([]int{3, 2, 1})
reverse := false
list.Sort(reverse)
fmt.Println(list) // [1, 2, 3]
```

## list.Sorted(reverse) *golist.List
Returns a list of Sorted items (the argument can be used for sort customization. `reverse` is `bool` so can be `true` or `false`).
```golang
list := golist.NewList([]int{3, 2, 1})
reverse := false
Newlist := list.Sorted(reverse)
fmt.Println(Newlist) // [1, 2, 3]
```

## list.Reverse() *golist.List
Returns new list with the elements of the list reversed.
```golang
list := golist.NewList([]int{5, 2, 7, 1})
newList := list.Reverse()
fmt.Println(newList) // [1, 7, 2, 5]
```

## list.Copy() (*golist.List, error)
Return a shallow copy of the list.
```golang
list := golist.NewList([]int{3, 2, 1})
copy, err := list.Copy()
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(copy) // [3, 2, 1]
```

## list.Sum() interface{}
Returns sum of the elements in the list. If it is list of string, it joins the strings with a space and returns it
```golang
list := golist.NewList([]int{3, 2, 1})
fmt.Println(list.Sum())  // 6
```

```golang
list := golist.NewList([]string{"Hello", "World"})
fmt.Println(list.Sum())  // "Hello World"
```

## list.List() interface{}
This is a getter that returns underlying slice interface.
```golang
list := golist.NewList([]int{3, 2, 1})
fmt.Println(list.List())  // [3 2 1]
```

## list.Join(joiner) string
This only works with string data types, panics otherwise. `joiner` is a string used to join the list.
```golang
list := golist.NewList([]string{"Hello", "World"})
fmt.Println(list.Join("-"))  // "Hello-World"
```

## list.Replace(x, i) error
Replaces an element at index i with element x. returns error if index does not exist. index of `-1` is equivalent to last item. This method is equivalent to working with slice (`a`) `a[1] = 10`
```golang
list := golist.NewList([]string{"Hello", "World"})
err := list.Replace("golang", -1)
if err != nil {
    fmt.Println(err)  // handle error
}
fmt.Println(list)  // ["Hello", "golang"]
```

## list.Max() (interface{}, error)
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

## list.Min() (interface{}, error)
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

## list.GCF() (interface{}, error)
Returns the Greatest Common Factor (GCF) or Highest Common Factor (HCF) of the numbers in the list. Only works with numbers. Returns error if called on list of strings. Uses [Euclidean algorithm](https://en.wikipedia.org/wiki/Euclidean_algorithm)
```golang
list := golist.NewList([]string{10, 15, 5})
gcf, err := list.GCF()
if err != nil {
    fmt.Println(err)  // handle error
}
fmt.Println(gcf)  // 5
```