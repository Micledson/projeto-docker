package response

import (
	"net/http"
	"projeto-docker/src/core/domain/errors"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/wallrony/go-validator/validator"
)

var validationErrorRegexCompiler = regexp.MustCompile(`^('.+?') (.*)`)

type ErrorMessage struct {
	error
	Code          int            `json:"status_code,omitempty"`
	Message       string         `json:"message"`
	InvalidFields []InvalidField `json:"invalid_fields,omitempty"`
	isInternal    bool
}

type InvalidField struct {
	FieldName   string `json:"field_name"`
	Description string `json:"description"`
}

type errorBuilder struct{}

var unprocessableEntityError = &echo.HTTPError{
	Code: http.StatusUnprocessableEntity,
}
var unsupportedMediaTypeError = &echo.HTTPError{
	Message: "Unsupported Media Type",
	Code:    http.StatusUnsupportedMediaType,
}

var notFoundError = &echo.HTTPError{
	Code: http.StatusNotFound,
}

var badRequestError = &echo.HTTPError{
	Code: http.StatusBadRequest,
}
var internalServerError = &echo.HTTPError{
	Code:    http.StatusInternalServerError,
	Message: "Ocorreu um erro inesperado. Por favor, contate o suporte.",
}
var unauthorizedError = &echo.HTTPError{
	Code: http.StatusUnauthorized,
}
var forbiddenError = &echo.HTTPError{
	Code: http.StatusForbidden,
}

func ErrorBuilder() *errorBuilder {
	return &errorBuilder{}
}

func (e *errorBuilder) NewFromDomain(err errors.Error) *echo.HTTPError {
	code := badRequestError.Code
	if err.CausedByValidation() {
		return e.unprocessableEntityErrorWithMessage(err.String())
	} else if err.CausedInternally() {
		return internalServerError
	} else if err.CausedByNotFound() {
		code = notFoundError.Code
	}
	return &echo.HTTPError{
		Code:    code,
		Message: err.String(),
	}
}

func (*errorBuilder) NewForbiddenError() *echo.HTTPError {
	return forbiddenError
}

func (*errorBuilder) NewNotFoundError() *echo.HTTPError {
	return notFoundError
}

func (*errorBuilder) NewUnauthorizedError() *echo.HTTPError {
	return unauthorizedError
}

func (*errorBuilder) NewUnsupportedMediaTypeError() *echo.HTTPError {
	return unsupportedMediaTypeError
}

func (*errorBuilder) NewBadRequestFromCoreError() *echo.HTTPError {
	return unsupportedMediaTypeError
}

func (*errorBuilder) NewInternalServerError() *echo.HTTPError {
	return internalServerError
}

func (*errorBuilder) badRequestErrorWithMessage(message string) *echo.HTTPError {
	err := badRequestError
	err.Message = message
	return err
}

func (*errorBuilder) internalErrorWithMessage(message string) *echo.HTTPError {
	err := internalServerError
	err.Message = message
	return err
}

func (*errorBuilder) unprocessableEntityErrorWithMessage(message string) *echo.HTTPError {
	err := unprocessableEntityError
	err.Message = message
	return err
}

func (e *errorBuilder) NewFromValidationError(valErr validator.ValidationError) *echo.HTTPError {
	err := errors.NewValidation(valErr.Messages())
	return e.unprocessableEntityErrorWithMessage(err.String())
}

func (e *ErrorMessage) Error() echo.HTTPError {
	return echo.HTTPError{
		Message: e.Message,
		Code:    e.Code,
	}
}

func (e *ErrorMessage) IsInternal() bool {
	return e.isInternal
}
