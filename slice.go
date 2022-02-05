package golist

func validateSlicingArgs(listLen, start, stop int) (int, error) {
	if start > listLen || stop > listLen {
		return stop, ErrIndexOutOfRange
	} else if start < 0 || stop < -1 {
		return stop, ErrIndexOutOfRange
	} else if stop == -1 {
		stop = listLen
		return stop, nil
	} else {
		return stop, nil
	}
}

// Slice a list. similar to []int{1,2,3,4}[start:stop]
func (arr *List) Slice(start int, stop int) (*List, error) {

	switch list := arr.list.(type) {
	case []int:
		stop, err := validateSlicingArgs(len(list), start, stop)
		if err != nil {
			return nil, err
		}
		return NewList(list[start:stop]), nil

	case []int32:
		stop, err := validateSlicingArgs(len(list), start, stop)
		if err != nil {
			return nil, err
		}
		return NewList(list[start:stop]), nil

	case []int64:
		stop, err := validateSlicingArgs(len(list), start, stop)
		if err != nil {
			return nil, err
		}
		return NewList(list[start:stop]), nil

	case []float32:
		stop, err := validateSlicingArgs(len(list), start, stop)
		if err != nil {
			return nil, err
		}
		return NewList(list[start:stop]), nil

	case []float64:
		stop, err := validateSlicingArgs(len(list), start, stop)
		if err != nil {
			return nil, err
		}
		return NewList(list[start:stop]), nil

	case []string:
		stop, err := validateSlicingArgs(len(list), start, stop)
		if err != nil {
			return nil, err
		}
		return NewList(list[start:stop]), nil

	default:
		return nil, ErrTypeNotsupported
	}

}
