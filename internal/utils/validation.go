package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func capitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(str[:1]) + str[1:]
}

func FormatValidationError(err error, obj interface{}) map[string]string {
	errors := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errors["error"] = err.Error()
		return errors
	}

	for _, e := range validationErrors {

		field := getFieldName(e, obj)
		label := capitalize(field)

		switch e.Tag() {

		case "required":
			errors[field] = label + " is required"

		case "email":
			errors[field] = label + " must be a valid email address"

		case "min":
			errors[field] = label + " must be at least " + e.Param() + " characters"

		case "max":
			errors[field] = label + " must not exceed " + e.Param() + " characters"

		case "len":
			errors[field] = label + " must be exactly " + e.Param() + " characters"

		case "eq":
			errors[field] = label + " must be equal to " + e.Param()

		case "ne":
			errors[field] = label + " must not be equal to " + e.Param()

		case "eqfield":
			errors[field] = label + " must match " + toSnakeCase(e.Param())

		case "oneof":
			errors[field] = label + " must be one of: " + e.Param()

		case "numeric":
			errors[field] = label + " must be a number"

		case "alphanum":
			errors[field] = label + " must contain only letters and numbers"

		case "url":
			errors[field] = label + " must be a valid URL"

		case "uuid":
			errors[field] = label + " must be a valid UUID"

		case "gte":
			errors[field] = label + " must be greater than or equal to " + e.Param()

		case "lte":
			errors[field] = label + " must be less than or equal to " + e.Param()

		default:
			errors[field] = "Invalid value for " + label
		}
	}

	return errors
}

func getFieldName(e validator.FieldError, obj interface{}) string {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	field, ok := t.FieldByName(e.StructField())
	if !ok {
		return toSnakeCase(e.Field())
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag != "" && jsonTag != "-" {
		return strings.Split(jsonTag, ",")[0]
	}

	formTag := field.Tag.Get("form")
	if formTag != "" && formTag != "-" {
		return strings.Split(formTag, ",")[0]
	}

	return toSnakeCase(e.Field())
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