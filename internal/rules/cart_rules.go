package rules

import "go-ecommerce-service/internal/dto"

type CartRules struct {
	BaseRules[dto.CreateCartRequest]
}

func NewCartRules() *CartRules {
	return &CartRules{}
}
