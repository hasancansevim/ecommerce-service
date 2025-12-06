package rules

import "go-ecommerce-service/internal/dto"

type StoreRules struct {
	BaseRules[dto.CreateStoreRequest]
}

func NewStoreRules() *StoreRules {
	return &StoreRules{}
}
