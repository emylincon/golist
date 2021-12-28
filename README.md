![status](https://github.com/emylincon/golist/workflows/Go/badge.svg)

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/emylincon/golist/Go?style=for-the-badge)
# golist
A customized go list with index, sort, append, pop, count, clear and last item methods. The list data type has some more methods. 

It supports the following data structure:
* `int`
* `int32`
* `int64`
* `float32`
* `float64`
* `string`

Here are all of the methods of the list objects:

## list.Get(i) (interface{}, error)
Get an item in the list by index. `i` represents the index. 
```golang
list := NewList([]int{1,2,3})
item, err := list.Get(0)
if err != nil {
    fmt.Println(err)  // handle error
}
fmt.Println(item)  // 1
```

## list.Append(x)
Add an item to the end of the list.
```golang
list := NewList([]int{1,2,3})
list.Append(7)
fmt.Println(list)  // {1,2,3,7}
```

## list.Extend(slice)
Extend the list by appending all the items from a slice or array.
```golang
list := NewList([]int{1,2,3})
list.Extend([]int{4,5})
fmt.Println(list)  // {1,2,3,4,5}
```

## list.Insert(x, i) error
Insert an item at a given position. The first argument is the index of the element before which to insert, so a.insert(0, x) inserts at the front of the list, and a.Insert(x, len(a)) is equivalent to a.Append(x). Returns error is any
```golang
list := NewList([]int{1,2,3})
err := list.Insert(4,3)
if err {
    fmt.Println(err) // handle error
}
fmt.Println(list)  // {1,2,3,4}
```
* The above inserts item 4 to position 3 which is the end of the list

## list.Remove(x) #TODO
Remove the first item from the list whose value is equal to x. It raises a ValueError if there is no such item.

## list.Pop(i)
Remove the item at the given position in the list, and return it. i is the index of the element to be popped.
```golang
list := NewList([]int{1,2,3})
popped := list.Pop(0)
fmt.Println(popped)  // 1
fmt.Println(list)  // {2,3}
```

## list.Clear()
Remove all items from the list.
```golang
list := NewList([]int{1,2,3})
list.Clear()
fmt.Println(list)  // {}
```

## list.Slice(start, step, end) #TODO
Return zero-based index in the list of the first item whose value is equal to x. Raises a ValueError if there is no such item.

The optional arguments start and end are interpreted as in the slice notation and are used to limit the search to a particular subsequence of the list. The returned index is computed relative to the beginning of the full sequence rather than the start argument.

## list.Count(x)
Return the number of times x appears in the list.
```golang
list := NewList([]int{1,2,3,2})
count := list.Count(2)
fmt.Println(count)  // 2
```

## list.sort(reverse)
Sort the items of the list in place (the arguments can be used for sort customization. `reverse` is `bool` so can be `true` or `false`.
```golang
list := NewList([]int{3,2,1})
list.Sort(reverse=false)
fmt.Println(list)  // {1,2,3}
```

## list.reverse() #TODO
Reverse the elements of the list in place.

## list.copy()   #TODO
Return a shallow copy of the list. Equivalent to a[:].