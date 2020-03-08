package gcm

import (
	"fmt"
)

type sentinelError string

func (s sentinelError) Error() string {
	return string(s)
}

const (
	ProviderNotDefinedError sentinelError = "no provider defined for the specified key"
	ValueNotProvidedError   sentinelError = "no value was provided for the specified key"
)

type ValidationError struct {
	Err error
	V   interface{}
}

func NewValidationError(err error, v interface{}) *ValidationError {
	return &ValidationError{err, v}
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("failed validation: %s", v.Err.Error())
}

func (v *ValidationError) Unwrap() error {
	return v.Err
}

type InvalidTypeError struct {
	ExpectedType string
	ActualType   string
}

func NewInvalidTypeError(expected, actual string) *InvalidTypeError {
	return &InvalidTypeError{
		ExpectedType: expected,
		ActualType:   actual,
	}

}

func (i *InvalidTypeError) Error() string {
	return fmt.Sprintf("invalid type: expected '%s', received '%s'", i.ExpectedType, i.ActualType)
}
