package pkgerror

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	customCodeExample                    = "CODE_{EntityRequest}_{Issue}"
	CustomCodeDefaultBadRequest          = "BAD_REQUEST"
	CustomCodeDefaultUnauthorized        = "UNAUTHORIZED"
	CustomCodeDefaultForbidden           = "FORBIDDEN"
	CustomCodeDefaultNotFound            = "NOT_FOUND"
	CustomCodeDefaultInternalServerError = "INTERNAL_SERVER_ERROR"
)

// GetErrValidationMessage re-format the error message of validator pkg
// Lib `github.com/go-playground/validator/v10`.
func GetErrValidationMessage(err error) string {
	var (
		validationErrors validator.ValidationErrors
		message          = err.Error()
	)

	if errors.As(err, &validationErrors) {
		message = "Invalid field"

		if len(validationErrors) > 1 {
			message += "s"
		}

		for i := range validationErrors {
			message += fmt.Sprintf(
				" '%v',",
				strings.ToLower(regexp.MustCompile("([a-z0-9])([A-Z])").
					ReplaceAllString(validationErrors[i].Field(), "${1}_${2}")),
			)
		}

		message = strings.TrimSuffix(message, ",")
	}

	return message
}

func ErrValidation(err error, info ...string) MyError {
	if err == nil {
		return ErrBadRequest()
	}

	customCode := CustomCodeDefaultBadRequest
	if len(info) > 0 && strings.TrimSpace(info[0]) != "" {
		customCode = info[0]
	}

	return MyError{
		Raw:       nil,
		ErrorCode: customCode,
		Message:   GetErrValidationMessage(err),
	}
}

func ErrBadRequest(info ...string) MyError {
	var (
		message    = "Bad Request"
		customCode = CustomCodeDefaultBadRequest
	)

	if len(info) > 0 && strings.TrimSpace(info[0]) != "" {
		message = info[0]
	}

	if len(info) > 1 && strings.TrimSpace(info[1]) != "" {
		customCode = info[1]
	}

	return MyError{
		Raw:       nil,
		ErrorCode: customCode,
		Message:   message,
	}
}

func ErrInternalServerError(err error, info ...string) MyError {
	var (
		message    = "Internal Server Error"
		customCode = CustomCodeDefaultInternalServerError
	)

	if len(info) > 0 && strings.TrimSpace(info[0]) != "" {
		message = info[0]
	}

	if len(info) > 1 && strings.TrimSpace(info[1]) != "" {
		customCode = info[1]
	}

	return MyError{
		Raw:       err,
		ErrorCode: customCode,
		Message:   message,
	}
}
