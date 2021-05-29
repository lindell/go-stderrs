package stderrs

import "errors"

// ServiceUnavailable is used as an error when a service was unavailable
type ServiceUnavailable struct {
	Message string
	Wrapped error
}

// NewServiceUnavailable returns a new error ServiceUnavailable error
func NewServiceUnavailable(message string) error {
	return ServiceUnavailable{
		Message: message,
	}
}

// NewWrappedServiceUnavailable returns a new error ServiceUnavailable error that is wrapping another error
func NewWrappedServiceUnavailable(err error, message string) error {
	return ServiceUnavailable{
		Message: message,
		Wrapped: err,
	}
}

type iServiceUnavailable interface {
	ServiceUnavailable() bool
}

// IsServiceUnavailable returns checks if a func ServiceUnavailable() bool exist and if that it returns true
func IsServiceUnavailable(err error) bool {
	if err == nil {
		return false
	}

	if errServiceUnavailable, ok := err.(iServiceUnavailable); ok && errServiceUnavailable.ServiceUnavailable() {
		return true
	}

	return IsServiceUnavailable(errors.Unwrap(err))
}

// ServiceUnavailable indicates that this error is used when a service was unavailable
func (e ServiceUnavailable) ServiceUnavailable() bool {
	return true
}

// Error returns an error message
func (e ServiceUnavailable) Error() string {
	if e.Message == "" {
		return "a service was unavailable"
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e ServiceUnavailable) Unwrap() error {
	return e.Wrapped
}

// Temporary indicated that this error is temporary
func (e ServiceUnavailable) Temporary() bool {
	return true
}

// HTTPStatusCode returns that this error should be represted as the 503 status code
func (e ServiceUnavailable) HTTPStatusCode() int {
	return 503
}
