package golist

// IsEqual :
// checks if two lists are equal and returns a boolean
func (arr *List) IsEqual(other *List) bool {
	if arr.Len() != other.Len() {
		return false
	}
	for i := 0; i < arr.Len(); i++ {
		if arr.Get(i) != other.Get(i) {
			return false
		}
	}
	return true
}
