package core

import (
	"errors"
)

var (
	ErrNotZeroOrPositive = errors.New("numbers must be zero or positive")
	ErrNotInList         = errors.New("element not in list")
	ErrListEmpty         = errors.New("list is empty")
)
