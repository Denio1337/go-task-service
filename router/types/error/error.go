package error

import (
	"app/router/validator"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Constructs a new fiber error with given code and message
func New(code int, message string) *fiber.Error {
	return &fiber.Error{
		Code:    code,
		Message: message,
	}
}

// Constructs invalid validation error
func ValidationError(errs []validator.ValidationError) *fiber.Error {
	errMsgs := make([]string, 0)

	// Formatting list of validation errors
	for _, err := range errs {
		errMsgs = append(errMsgs, validationMessage(err.FailedField, err.Value, err.Tag))
	}

	return &fiber.Error{
		Code:    fiber.ErrBadRequest.Code,
		Message: strings.Join(errMsgs, ValidationErrorSeparator),
	}
}

// Constructs validation error message
func validationMessage(field string, value any, tag string) string {
	return fmt.Sprintf(
		"[%s]: '%v' | Needs to implement '%s'",
		field,
		value,
		tag,
	)
}
