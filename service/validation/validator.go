package validation

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Validator struct {
	errors []ValidationError
}

func NewValidator() *Validator {
	return &Validator{
		errors: make([]ValidationError, 0),
	}
}

func (v *Validator) RequiredString(value, field string) *Validator {
	if strings.TrimSpace(value) == "" {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s is required", field),
		})
	}
	return v
}

func (v *Validator) RequiredInt(value int, field string) *Validator {
	if value == 0 {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%v is required", field),
		})
	}
	return v
}

func (v *Validator) RequiredFloat(value float32, field string) *Validator {
	if value == 0 {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%v is required", field),
		})
	}
	return v
}

func (v *Validator) MinLength(value, field string, min int) *Validator {
	if utf8.RuneCountInString(strings.TrimSpace(value)) < min {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must be at least %d characters", field, min),
		})
	}
	return v
}

func (v *Validator) MaxLength(value, field string, max int) *Validator {
	if utf8.RuneCountInString(strings.TrimSpace(value)) > max {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must be at most %d characters", field, max),
		})
	}
	return v
}

func (v *Validator) Range(value int, field string, min, max int) *Validator {
	if value < min || value > max {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must be in range [%d,%d]", field, min, max),
		})
	}
	return v
}

func (v *Validator) Error() error {
	if len(v.errors) > 0 {
		return fmt.Errorf("validation failed: %v", v.errors)
	}
	return nil
}
