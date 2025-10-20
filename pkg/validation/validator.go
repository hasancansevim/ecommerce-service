package validation

import (
	_errors "go-ecommerce-service/pkg/errors"
)

type Validator interface {
	Validate() *_errors.AppError
}
