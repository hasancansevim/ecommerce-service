package rules

import (
	"errors"
	"go-ecommerce-service/internal/dto"
)

type ProductRules struct {
	BaseRules[dto.CreateProductRequest]
}

func NewProductRules() *ProductRules {
	return &ProductRules{}
}

func (r *ProductRules) ValidateCreate(req dto.CreateProductRequest) error {
	// Validation Struct
	if err := r.ValidateStructure(req); err != nil {
		return err
	}

	// Business Rules
	if req.Price < 0 {
		return errors.New("Ürün fiyatı 0 dan küçük olamaz.")
	}
	if req.Discount < 0 {
		return errors.New("İndirim oranı 0 dan küçük olamaz")
	}

	return nil
}
