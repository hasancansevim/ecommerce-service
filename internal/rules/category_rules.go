package rules

import (
	"go-ecommerce-service/internal/dto"
)

type CategoryRules struct {
	BaseRules[dto.CreateCategoryRequest]
}

func NewCategoryRules() *CategoryRules {
	return &CategoryRules{}
}

// Business Rules
