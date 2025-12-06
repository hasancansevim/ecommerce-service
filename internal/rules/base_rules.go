package rules

import (
	"go-ecommerce-service/pkg/validation"
)

type BaseRules[T any] struct{}

func (b *BaseRules[T]) ValidateStructure(req T) error {
	return validation.ValidateStruct(req)
}
