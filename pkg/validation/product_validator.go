// pkg/validation/product_validator.go
package validation

import (
	_errors "go-ecommerce-service/pkg/errors"
	"go-ecommerce-service/service/model"
	"strings"
)

type ProductCreateValidator struct {
	ProductReq model.ProductCreate
}

func (v ProductCreateValidator) Validate() *_errors.AppError {
	var errors []ValidationError

	if strings.TrimSpace(v.ProductReq.Name) == "" {
		errors = append(errors, ValidationError{Field: "name", Message: "Product name is required"})
	}

	if v.ProductReq.Price <= 0 {
		errors = append(errors, ValidationError{Field: "price", Message: "Price must be greater than 0"})
	}

	if v.ProductReq.Discount < 0 || v.ProductReq.Discount > 100 {
		errors = append(errors, ValidationError{Field: "discount", Message: "Discount must be between 0 and 100"})
	}

	if len(errors) > 0 {
		return _errors.NewValidationError(errors)
	}
	return nil
}
