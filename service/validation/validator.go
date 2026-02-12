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

func (v ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Message)
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

func (v *Validator) RequiredFloat(value float64, field string) *Validator {
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

func (v *Validator) RangeFloat64(value float64, field string, min, max float64) *Validator {
	if value < min || value > max {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must be in range [%f,%f]", field, min, max),
		})
	}
	return v
}

func (v *Validator) GreaterThanFloat32(value float32, threshold float32, field string) *Validator {
	if value <= threshold {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must be greater than %v", field, threshold),
		})
	}
	return v
}

func (v *Validator) GreaterThanFloat64(value float64, threshold float64, field string) *Validator {
	if value <= threshold {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must be greater than %v", field, threshold),
		})
	}
	return v
}

func (v *Validator) GreaterThanInt(value int, threshold int, field string) *Validator {
	if value <= threshold {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must be greater than %d", field, threshold),
		})
	}
	return v
}

func (v *Validator) Error() error {
	if len(v.errors) > 0 {
		return ValidationErrors(v.errors)
	}
	return nil
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	if len(v) == 0 {
		return ""
	}
	return fmt.Sprintf("validation failed with %d errors", len(v))
}
