package golist

import "strings"

// Join  :
// joins elements in a string list with a joiner
func (arr *List) Join(joiner string) string {
	switch list := arr.list.(type) {
	case []string:
		return strings.Join(list, joiner)
	default:
		panic("list.Join can only be used with strings")
	}
}
