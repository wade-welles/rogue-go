package rogue

import (
	"errors"
)

var ErrVersion = errors.New("thread ID needs linux >= 3.4")
var ErrNotStack = errors.New("mapping is not a stack")
