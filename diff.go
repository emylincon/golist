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

// DifferenceBoth returns the elements that aren't in both lists.
func (arr *List) DifferenceBoth(other *List) (*List, error) {
	if arr.Type() != other.Type() {
		return nil, ErrListsNotOfSameType
	} else if arr.Type() == TypeListUnknown {
		return nil, ErrTypeNotsupported
	}
	diff, _ := arr.Copy()
	diff.Clear()
	m := map[interface{}]int{}

	for i := 0; i < arr.Len(); i++ {
		m[arr.Get(i)] = 1
	}
	for i := 0; i < other.Len(); i++ {
		m[other.Get(i)] = m[other.Get(i)] + 1
	}

	for mKey, mVal := range m {
		if mVal == 1 {
			diff.Append(mKey)
		}
	}

	return diff, nil

}
