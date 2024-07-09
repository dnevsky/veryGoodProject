package errors

import "errors"

var (
	ErrNotFound      = errors.New("row not found")
	ErrAlreadyExists = errors.New("already exists")
)
