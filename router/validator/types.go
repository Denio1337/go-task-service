package validator

import "github.com/go-playground/validator/v10"

// Package instance of validator
var instance *validator.Validate

// Validation error structure
type ValidationError struct {
	FailedField string
	Tag         string
	Value       any
}
