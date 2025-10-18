package fixture

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"
)

func CreateTestProduct() domain.Product {
	return domain.Product{
		Id:       1,
		Name:     "Test Product",
		Price:    99.99,
		Discount: 10.0,
		Store:    "test-store",
	}
}

func CreateTestProductCreate() model.ProductCreate {
	return model.ProductCreate{
		Name:     "New Product",
		Price:    1,
		Discount: 15.0,
		Store:    "new-store",
	}
}
