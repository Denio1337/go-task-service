package error

import "github.com/gofiber/fiber/v2"

// Separator for validation error messages
const ValidationErrorSeparator = " & "

// General router errors
var (
	// 400 Bad Request
	ErrBadRequest = fiber.NewError(fiber.StatusBadRequest, "invalid input data")

	// 401 Unauthorized
	ErrUnauthorized = fiber.NewError(fiber.StatusUnauthorized, "unauthorized access")

	// 403 Forbidden
	ErrForbidden = fiber.NewError(fiber.StatusForbidden, "access forbidden")

	// 404 Not Found
	ErrNotFound = fiber.NewError(fiber.StatusNotFound, "resource not found")

	// 409 Conflict
	ErrConflict = fiber.NewError(fiber.StatusConflict, "resource conflict")

	// 500 Internal Server Error
	ErrInternalServer = fiber.NewError(fiber.StatusInternalServerError, "internal server error")
)
