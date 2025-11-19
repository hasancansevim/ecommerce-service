package validation

import "go-ecommerce-service/service/model"

func ValidateProductCreate(productCreate model.ProductCreate) error {
	return NewValidator().
		RequiredString(productCreate.Name, "name").
		MinLength(productCreate.Name, "name", 2).
		MaxLength(productCreate.Name, "name", 255).
		RequiredString(productCreate.Slug, "slug").
		MinLength(productCreate.Slug, "slug", 2).
		MaxLength(productCreate.Slug, "slug", 255).
		RequiredFloat(productCreate.Price, "price").
		RangeFloat64(productCreate.Price, "price", 0.01, 1_000_000).
		RangeFloat64(productCreate.Discount, "discount", 0, 100).
		Error()
}
