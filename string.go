package golist

import (
	"fmt"
	"strings"
)

// returns string representation of the list
func (arr *List) String() string {
	switch arr.list.(type) {
	case []string:
		list := arr.list.([]string)
		stringList := "["
		for i := 0; i < len(list); i++ {
			if i == len(list)-1 {
				stringList = fmt.Sprintf("%s\"%s\"]", stringList, list[i])
			} else {
				stringList = fmt.Sprintf("%s\"%s\", ", stringList, list[i])
			}

		}
		return stringList
	default:
		return strings.Replace(fmt.Sprintf("%v", arr.list), " ", ", ", -1)
	}

}
