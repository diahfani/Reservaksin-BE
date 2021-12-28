package businesses

import "errors"

var (
	ErrDuplicateData         = errors.New("data already exists")
	ErrEmailPasswordNotFound = errors.New("data not found")
	ErrInternalServer        = errors.New("something wrong")
)
