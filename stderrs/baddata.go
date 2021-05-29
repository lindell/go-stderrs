package stderrs

import "errors"

// BadData is used as an error when data provided was incorrectly formated or does not meet the specification
type BadData struct {
	Message string
	Wrapped error
}

// NewBadData returns a new error BadData error
func NewBadData(message string) error {
	return BadData{
		Message: message,
	}
}

// NewWrappedBadData returns a new error BadData error that is wrapping another error
func NewWrappedBadData(err error, message string) error {
	return BadData{
		Message: message,
		Wrapped: err,
	}
}

type iBadData interface {
	BadData() bool
}

// IsBadData returns checks if a func BadData() bool exist and if that it returns true
func IsBadData(err error) bool {
	if err == nil {
		return false
	}

	if errBadData, ok := err.(iBadData); ok && errBadData.BadData() {
		return true
	}

	return IsBadData(errors.Unwrap(err))
}

// BadData indicates that this error is used when data provided was incorrectly formated or does not meet the specification
func (e BadData) BadData() bool {
	return true
}

// Error returns an error message
func (e BadData) Error() string {
	if e.Message == "" {
		return "data provided was incorrectly formated or does not meet the specification"
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e BadData) Unwrap() error {
	return e.Wrapped
}

// Temporary indicated that this error is not temporary
func (e BadData) Temporary() bool {
	return false
}

// HTTPStatusCode returns that this error should be represted as the 400 status code
func (e BadData) HTTPStatusCode() int {
	return 400
}
