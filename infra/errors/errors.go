package errors

import (
	"errors"
	"fmt"
)

// DError define
type DError struct {
	err  error
	code int
}

// New create new error
func New(text string) error {
	return errors.New(text)
}

// NewWithCode create new error
func NewWithCode(text string, code int) error {
	return DError{err: errors.New(text), code: code}
}

// Wrap create new error
func Wrap(err error, code int) error {
	return DError{
		err:  err,
		code: code,
	}
}

func (e DError) Error() string {
	return e.err.Error()
}

func (e DError) String() string {
	return fmt.Sprintf("%d: %s", e.code, e.err.Error())
}

// Code get error code
func (e DError) Code() int {
	return e.code
}
