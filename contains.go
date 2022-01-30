package golist

// Contains :
// returns true if element exists, returns false otherwise
// ```golang
// list := golist.NewList([]string{"Hello", "World"})
// fmt.Println(list.Contains("okay"))  // false
// ```
func (arr *List) Contains(element interface{}) bool {
	return arr.Index(element) != -1

}
