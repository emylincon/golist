package golist

func (arr *List) Exists(element interface{}) bool {
	if arr.Index(element) == -1 {
		return false
	} else {
		return true
	}

}
