package golist

import "strings"

func (arr *List) Join(joiner string) string {
	switch arr.list.(type) {
	case []string:
		list := arr.list.([]string)
		return strings.Join(list, joiner)
	default:
		panic("list.Join can only be used with strings")
	}
}
