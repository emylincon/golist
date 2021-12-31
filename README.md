![status](https://github.com/emylincon/golist/workflows/Go/badge.svg)

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/emylincon/golist/Go?style=for-the-badge)
# golist
A customized go list with index, sort, append, pop, count, clear and last item methods. 
It supports all of the following data structure although only `int` is used in the example:
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
item := list.Get(0)
fmt.Println(item)  // 1
```

## list.Append(x)
Add an item to the end of the list.
```golang
list := NewList([]int{1,2,3})
list.Append(7)
fmt.Println(list)  // &{[1 2 3 7]}
```

## list.Extend(slice)
Extend the list by appending all the items from a slice or array.
```golang
list := NewList([]int{1,2,3})
list.Extend([]int{4,5})
fmt.Println(list)  // &{[1 2 3 4 5]}
```

## list.Insert(x, i) error
Insert an item at a given position. The first argument is the index of the element before which to insert, so a.insert(0, x) inserts at the front of the list, and a.Insert(x, len(a)) is equivalent to a.Append(x). Returns error is any
```golang
list := NewList([]int{1,2,3})
err := list.Insert(4,3)
if err {
    fmt.Println(err) // handle error
}
fmt.Println(list)  // &{[1 2 3 4]}
```
* The above inserts item 4 to position 3 which is the end of the list

## list.Remove(x) (error)
Remove the first item from the list whose value is equal to x. It raises a ValueError if there is no such item.
```golang
list := NewList([]int{1,2,3})
err := list.Remove(2)
fmt.Println(list)  // &{[2 3]}
```

## list.Pop(i)
Remove the item at the given position in the list, and return it. i is the index of the element to be popped.
```golang
list := NewList([]int{1,2,3})
popped := list.Pop(0)
fmt.Println(popped)  // 1
fmt.Println(list)  // &{[2 3]}
```


## list.Clear()
Remove all items from the list.
```golang
list := NewList([]int{1,2,3})
list.Clear()
fmt.Println(list)  // &{[]}
```

## list.Slice(start, end)
The optional arguments start and end are interpreted as in the slice notation and are used to limit the search to a particular subsequence of the list. The returned index is computed relative to the beginning of the full sequence rather than the start argument.
```golang
list := NewList([]int{1,2,3,2})
start, stop := 0,2
NewList := list.Slice(start,stop)
fmt.Println(NewList)  // &{[1 2]}
```

## list.Count(x)
Return the number of times x appears in the list.
```golang
list := NewList([]int{1,2,3,2})
count := list.Count(2)
fmt.Println(count)  // 2
```

## list.Sort(reverse)
Sort the items of the list in place (the arguments can be used for sort customization. `reverse` is `bool` so can be `true` or `false`.
```golang
list := NewList([]int{3,2,1})
list.Sort(reverse=false)
fmt.Println(list)  // &{[1 2 3]}
```

## list.Reverse() 
Reverse the elements of the list in place.
```golang
list := NewList([]int{5,2,7,1})
list.Sort(reverse=false)
fmt.Println(list)  // &{[1 7 2 5}}
```

## list.Copy()
Return a shallow copy of the list.
```golang
list := NewList([]int{3,2,1})
copy := list.Copy()
fmt.Println(copy)  // &{[3 2 1]}
```

## list.String()
Returns a string representation of the object
```golang
list := NewList([]int{3,2,1})
fmt.Println(list.String())  // [3, 2, 1]
```

## list.Sum() interface{}
Returns sum of the elements in the list. If it is list of string, it joins the strings with a space and returns it
```golang
list := NewList([]int{3,2,1})
fmt.Println(list.Sum())  // 6
```

```golang
list := NewList([]string{"Hello", "World"})
fmt.Println(list.Sum())  // "Hello World"
```

## list.List() interface{}
This is a getter that returns underlying slice interface.
```golang
list := NewList([]int{3,2,1})
fmt.Println(list.List())  // [3 2 1]
```