package golist

// ## list.Exists(interface{}) bool
// returns true if element exists, returns false otherwise
// ```golang
// list := golist.NewList([]string{"Hello", "World"})
// fmt.Println(list.Exists("okay"))  // false
// ```
func (arr *List) Exists(element interface{}) bool {
	if arr.Index(element) == -1 {
		return false
	} else {
		return true
	}

}
