package validator

import "errors"

var ErrValidation = errors.New("validation error")

type ValidationError struct {
	Err     error
	ErrData map[string]string
}

func (e *ValidationError) Error() string {
	return e.Err.Error()
}
