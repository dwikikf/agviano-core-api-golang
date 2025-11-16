package web

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationErrorDetail struct {
	Field    string `json:"field"`
	Rule     string `json:"rule"`
	Expected string `json:"expected,omitempty"`
	Value    any    `json:"value,omitempty"`
	Message  string `json:"message"`
}

func FormatValidationError(err error) []ValidationErrorDetail {
	var valErrs validator.ValidationErrors

	// jika ini validation error
	if errors.As(err, &valErrs) {
		var errResult []ValidationErrorDetail

		for _, e := range valErrs {
			detail := ValidationErrorDetail{
				Field: e.Field(),
				Rule:  e.Tag(),
				Value: e.Value(),
			}

			// jika ada expected value nya
			if e.Param() != "" {
				detail.Expected = e.Param()
			}

			// switch untuk setiap pesan tag error nya
			switch e.Tag() {
			case "required":
				detail.Message = fmt.Sprintf("%s is required", e.Field())
			case "min":
				detail.Message = fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
			case "max":
				detail.Message = fmt.Sprintf("%s must be at most %s characters", e.Field(), e.Param())
			case "email":
				detail.Message = fmt.Sprintf("%s must be a valid email address", e.Field())
			default:
				detail.Message = fmt.Sprintf("%s is not valid", e.Field())
			}

			errResult = append(errResult, detail)
		}
		return errResult
	}

	// jika bukan validation err
	return []ValidationErrorDetail{
		{
			Message: err.Error(),
		},
	}

}

// func FormatValidationError(err error) []string {
// 	var errs validator.ValidationErrors

// 	if errors.As(err, &errs) {
// 		var messages []string

// 		for _, e := range errs {
// 			msg := fmt.Sprintf("Field %s : %s", e.Field(), e.Tag())
// 			messages = append(messages, msg)
// 		}
// 		return messages
// 	}

// 	return []string{err.Error()}
// }

// func ValidationErrorResponse(err error) map[string]string {
// 	errors := map[string]string{}

// 	if errs, ok := err.(validator.ValidationErrors); ok {
// 		for _, e := range errs {
// 			field := strings.ToLower(e.Field())
// 			errors[field] = fmt.Sprintf("%s is %s", field, e.Tag())
// 		}
// 	}

// 	return errors
// }
