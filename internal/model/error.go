package model

import (
	"fmt"
	"slices"
)

type ServiceError struct {
	Errors  []error
	Message string
}

func NewServiceError(err error, message string) *ServiceError {
	return &ServiceError{
		Errors:  []error{err},
		Message: message,
	}
}

func FromError(source, newError error) *ServiceError {
	var message string
	var errors []error

	serviceErrorFromSource, ok := source.(*ServiceError)
	if !ok {
		message = "api.internal"
		errors = []error{source, newError}
	} else {
		message = serviceErrorFromSource.Message
		errors = append(serviceErrorFromSource.Errors, newError)
	}

	return &ServiceError{
		Errors:  errors,
		Message: message,
	}
}

func BuildSubErrorWithOperation(op, message string) error {
	return fmt.Errorf("%s:%s", op, message)
}

func BuildErrorWithOperation(op, message string, err error) error {
	return fmt.Errorf("%s:%s:%s", op, message, err.Error())
}

func (s *ServiceError) Error() string {
	var result string

	errors := make([]error, len(s.Errors))
	copy(errors, s.Errors)

	slices.Reverse(errors)

	for _, err := range errors {
		result += err.Error() + ";\n"
	}

	return result
}

func (s *ServiceError) ErrorStringSlice() []string {
	var result []string

	errors := make([]error, len(s.Errors))
	copy(errors, s.Errors)

	slices.Reverse(errors)

	for _, err := range errors {
		result = append(result, err.Error()+";")
	}

	return result
}
