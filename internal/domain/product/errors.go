package domain

import "errors"

var ErrNotFound = errors.New("PRODUCT IS NOT FOUND")
var ErrSlugEmpty = errors.New("PRODUCT SLUG CANNOT BE EMPTY")
var ErrNameEmpty = errors.New("PRODUCT NAME CANNOT BE EMPTY")
var ErrInvalidPrice = errors.New("PRODUCT PRICE CANNOT BE NEGATIVE")
