package validation

import "go-ecommerce-service/service/model"

func ValidateProductCreate(productCreate model.ProductCreate) error {
	return NewValidator().
		RequiredString(productCreate.Name, "name").
		MinLength(productCreate.Name, "name", 2).
		RequiredFloat(productCreate.Price, "price").
		Range(int(productCreate.Price), "price", 1, 1.000_000).
		Range(int(productCreate.Discount), "discount", 1, 90).
		RequiredString(productCreate.Store, "store").
		MinLength(productCreate.Store, "store", 2).
		Error()
}
