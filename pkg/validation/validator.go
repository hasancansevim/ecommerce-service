package validation

import (
	_errors "go-ecommerce-service/pkg/errors"
	"go-ecommerce-service/service/model"
	"regexp"
	"strings"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateRegisterRequest(registerReq model.RegisterCreate) *_errors.AppError {
	var errors []ValidationError

	if strings.TrimSpace(registerReq.FirstName) == "" {
		errors = append(errors, ValidationError{Field: "first_name", Message: "First name is required"})
	}

	if strings.TrimSpace(registerReq.LastName) == "" {
		errors = append(errors, ValidationError{Field: "last_name", Message: "Last name is required"})
	}

	if !isValidEmail(registerReq.Email) {
		errors = append(errors, ValidationError{Field: "email", Message: "Invalid email format"})
	}

	if len(registerReq.Password) < 6 {
		errors = append(errors, ValidationError{Field: "password", Message: "Password must be at least 6 characters"})
	}

	if len(errors) > 0 {
		return _errors.NewValidationError(errors)
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(email, emailRegex)
	return matched
}
