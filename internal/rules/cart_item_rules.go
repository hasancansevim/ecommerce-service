package rules

import "go-ecommerce-service/internal/dto"

type CartItemRules struct {
	BaseRules[dto.CreateCartItemRequest]
}

func NewCartItemRules() *CartItemRules {
	return &CartItemRules{}
}
