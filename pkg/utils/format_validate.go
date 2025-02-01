package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	fields := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		switch err.Tag() {
		case "required":
			fields[field] = "This field is required"
		case "email":
			fields[field] = "Invalid email format"
		case "min":
			fields[field] = "Value is too short"
		case "max":
			fields[field] = "Value is too long"
		case "oneof":
			fields[field] = "Invalid value selected"
		default:
			fields[field] = "Invalid value"
		}
	}

	return fields
}
