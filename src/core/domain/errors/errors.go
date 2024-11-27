package errors

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

const (
	unexpectedErrorMessage = "An unexpected error occurred. Please, contact the support."
	GENERIC                = iota
	INTERNAL
	VALIDATION
	NOT_FOUND
)

type Error interface {
	error
	String() string
	Messages() []string
	CausedInternally() bool
	CausedByValidation() bool
	CausedByNotFound() bool
	Metadata() map[string]interface{}
	ValidationMessagesByMetadataFields(field []string) []string
}

type errorImpl struct {
	err      []string
	origin   int
	metadata map[string]interface{}
}

func new(errs []string, errType int, metadata map[string]interface{}) Error {
	return &errorImpl{errs, errType, metadata}
}

func New(err error) Error {
	return new([]string{err.Error()}, GENERIC, nil)
}

func NewWithMetadata(err error, metadata map[string]interface{}) Error {
	return new([]string{err.Error()}, GENERIC, metadata)
}

func NewNotFound(message string) Error {
	return new([]string{message}, NOT_FOUND, nil)
}

func NewFromString(message string) Error {
	return new([]string{message}, GENERIC, nil)
}

func NewInternal(err error) Error {
	return new([]string{err.Error()}, INTERNAL, nil)
}

func NewValidation(messages []string) Error {
	return new(messages, VALIDATION, nil)
}

func NewValidationWithMetadata(messages []string, metadata map[string]interface{}) Error {
	return new(messages, VALIDATION, metadata)
}

func NewValidationFromString(message string) Error {
	return new([]string{message}, VALIDATION, nil)
}

func NewUnexpected() Error {
	return NewInternal(errors.New(unexpectedErrorMessage))
}

func (e *errorImpl) String() string {
	return strings.Join(e.err, " & ")
}

func (e *errorImpl) Messages() []string {
	return e.err
}

func (e *errorImpl) CausedInternally() bool {
	return e.origin == INTERNAL
}

func (e *errorImpl) CausedByValidation() bool {
	return e.origin == VALIDATION
}

func (e *errorImpl) CausedByNotFound() bool {
	return e.origin == NOT_FOUND
}

func (e *errorImpl) Metadata() map[string]interface{} {
	return e.metadata
}

func (e *errorImpl) Error() string {
	return fmt.Sprint(e.err)
}

func (e *errorImpl) ValidationMessagesByMetadataFields(fields []string) []string {
	if e.metadata == nil || e.metadata["fields"] == nil {
		return nil
	}
	var metadataFields = e.metadata["fields"].([]string)
	var messages = []string{}
	for index, f := range metadataFields {
		if slices.Contains(fields, f) {
			messages = append(messages, e.Messages()[index])
		}
	}
	return messages
}
