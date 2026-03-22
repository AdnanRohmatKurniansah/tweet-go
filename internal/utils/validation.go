package utils

import (
	"strings"
	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errors["error"] = err.Error()
		return errors
	}
	for _, e := range validationErrors {

		field := toSnakeCase(e.Field())

		switch e.Tag() {

		case "required":
			errors[field] = field + " is required"

		case "email":
			errors[field] = field + " must be a valid email address"

		case "min":
			errors[field] = field + " must be at least " + e.Param() + " characters"

		case "max":
			errors[field] = field + " must not exceed " + e.Param() + " characters"

		case "len":
			errors[field] = field + " must be exactly " + e.Param() + " characters"

		case "eq":
			errors[field] = field + " must be equal to " + e.Param()

		case "ne":
			errors[field] = field + " must not be equal to " + e.Param()

		case "eqfield":
			errors[field] = field + " must match " + toSnakeCase(e.Param())

		case "oneof":
			errors[field] = field + " must be one of: " + e.Param()

		case "numeric":
			errors[field] = field + " must be a number"

		case "alphanum":
			errors[field] = field + " must contain only letters and numbers"

		case "url":
			errors[field] = field + " must be a valid URL"

		case "uuid":
			errors[field] = field + " must be a valid UUID"

		case "gte":
			errors[field] = field + " must be greater than or equal to " + e.Param()

		case "lte":
			errors[field] = field + " must be less than or equal to " + e.Param()

		default:
			errors[field] = "Invalid value for " + field
		}
	}

	return errors
}

func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {

		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}

		result = append(result, r)
	}

	return strings.ToLower(string(result))
}