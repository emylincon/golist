package core

import (
	"errors"
)

// Error variables that are common within the module
var (
	ErrNotZeroOrPositive = errors.New("numbers must be zero or positive")
	ErrNotInList         = errors.New("element not in list")
	ErrListEmpty         = errors.New("list is empty")
)
