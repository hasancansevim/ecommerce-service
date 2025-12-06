package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Message string
	Errors  map[string]string
}

func (v *ValidationError) Error() string {
	return v.Message
}

func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			validationErrors[field] = msgForTag(err.Tag(), err.Param())
		}

		return &ValidationError{
			Message: "Validation Error",
			Errors:  validationErrors,
		}
	}
	return nil
}

func msgForTag(tag string, param string) string {
	switch tag {
	case "required":
		return "Bu alan zorunludur"
	case "email":
		return "Geçerli bir email adresi giriniz."
	case "min":
		return fmt.Sprintf("En az %s karakter olmalıdır", param)
	case "max":
		return fmt.Sprintf("En çok %s karakter olmalı.", param)
	case "gt":
		return fmt.Sprintf("%s değerinden büyük olmalı.", param)
	default:
		return "Geçersiz değer."
	}
}
