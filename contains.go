package golist

// ## list.Contains(interface{}) bool
// returns true if element exists, returns false otherwise
// ```golang
// list := golist.NewList([]string{"Hello", "World"})
// fmt.Println(list.Contains("okay"))  // false
// ```
func (arr *List) Contains(element interface{}) bool {
	if arr.Index(element) == -1 {
		return false
	} else {
		return true
	}

}
