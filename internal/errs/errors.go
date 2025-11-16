package errs

import "errors"

var (
	ErrNotFound   = errors.New("not found")
	ErrConflict   = errors.New("conflict")
	ErrInvalid    = errors.New("invalid data")
	ErrForbidden  = errors.New("forbidden")
	ErrConflictFK = errors.New("conflict foreign key")
)
