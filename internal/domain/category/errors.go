package domain

import "errors"

var ErrNotFound = errors.New("CATEGORY IS NOT FOUND")
var ErrNameEmpty = errors.New("CATEGORY NAME CANNOT BE EMPTY")
