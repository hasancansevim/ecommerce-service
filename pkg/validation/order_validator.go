package validation

import (
	_errors "go-ecommerce-service/pkg/errors"
	"go-ecommerce-service/service/model"
)

type OrderCreateValidator struct {
	OrderReq model.OrderCreate
}

func (v OrderCreateValidator) Validate() *_errors.AppError {
	var errors []ValidationError

	if v.OrderReq.UserId == 0 {
		errors = append(errors, ValidationError{Field: "user_id", Message: "UserId is required"})
	}

	if v.OrderReq.TotalPrice < 0 {
		errors = append(errors, ValidationError{Field: "total_price", Message: "TotalPrice must be greater than zero"})
	}

	if len(errors) > 0 {
		return _errors.NewValidationError(errors)
	}
	return nil
}
