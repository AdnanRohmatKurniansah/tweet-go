package utils

import "errors"

var (
	ErrNotFound = errors.New("Resource not found")
	ErrAlreadyExists = errors.New("Resource already exists")
	ErrForbidden = errors.New("Forbidden")
	ErrUnauthorized = errors.New("Unauthorized")
	ErrBadRequest = errors.New("Bad request")
)