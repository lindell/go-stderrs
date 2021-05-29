package stderrs

import "errors"

// NotFound is used as an error when a resource could not be found
type NotFound struct {
	Message string
	Wrapped error
}

// NewNotFound returns a new error NotFound error
func NewNotFound(message string) error {
	return NotFound{
		Message: message,
	}
}

// NewWrappedNotFound returns a new error NotFound error that is wrapping another error
func NewWrappedNotFound(err error, message string) error {
	return NotFound{
		Message: message,
		Wrapped: err,
	}
}

type iNotFound interface {
	NotFound() bool
}

// IsNotFound returns checks if a func NotFound() bool exist and if that it returns true
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	if errNotFound, ok := err.(iNotFound); ok && errNotFound.NotFound() {
		return true
	}

	return IsNotFound(errors.Unwrap(err))
}

// NotFound indicates that this error is used when a resource could not be found
func (e NotFound) NotFound() bool {
	return true
}

// Error returns an error message
func (e NotFound) Error() string {
	if e.Message == "" {
		return "a resource could not be found"
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e NotFound) Unwrap() error {
	return e.Wrapped
}

// Temporary indicated that this error is not temporary
func (e NotFound) Temporary() bool {
	return false
}

// HTTPStatusCode returns that this error should be represted as the 404 status code
func (e NotFound) HTTPStatusCode() int {
	return 404
}
