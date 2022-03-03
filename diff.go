package golist

// Difference returns the elements in `list` that aren't in `other`.
func (arr *List) Difference(other *List) (*List, error) {
	if arr.Type() != other.Type() {
		return nil, ErrListsNotOfSameType
	} else if arr.Type() == TypeListUnknown {
		return nil, ErrTypeNotsupported
	}
	mb := make(map[interface{}]struct{}, other.Len())
	for i := 0; i < other.Len(); i++ {
		mb[other.Get(i)] = struct{}{}
	}

	diff, _ := arr.Copy()
	diff.Clear()
	for i := 0; i < arr.Len(); i++ {
		if _, found := mb[arr.Get(i)]; !found {
			diff.Append(arr.Get(i))
		}
	}
	return diff, nil
}
