package golist

import (
	"fmt"
	"strings"
)

func (arr *List) String() string {
	return strings.Replace(fmt.Sprintf("%v", arr.list), " ", ", ", -1)
}
