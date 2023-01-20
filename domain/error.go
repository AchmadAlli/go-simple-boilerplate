package domain

import "errors"

// General Error
var (
	ErrNotFound             = errors.New("requested item not found")
	ErrInvalidRequest       = errors.New("invalid request")
	ErrBadRequest           = errors.New("bad request")
	ErrInvalidID            = errors.New("invalid id value")
	ErrInternalServiceError = errors.New("internal service error")
	ErrTooManyRequest       = errors.New("too many request")
)
