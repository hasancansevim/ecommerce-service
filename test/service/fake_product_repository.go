package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

func (fakeRepository *FakeProductRepository) DeleteProductById(productId int64) error {
	//TODO implement me
	panic("implement me")
}

func (fakeRepository *FakeProductRepository) UpdateProduct(productId int64, newPrice float32) error {
	//TODO implement me
	panic("implement me")
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeRepository *FakeProductRepository) GetAllProducts() []domain.Product {
	return fakeRepository.products
}

func (fakeRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	fakeRepository.products = append(fakeRepository.products, domain.Product{
		Id:       int64(len(fakeRepository.products) + 1),
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
	return nil
}

func (f FakeProductRepository) GetAllByStoreName(storeName string) []domain.Product {
	//TODO implement me
	panic("implement me")
}

func (f FakeProductRepository) GetProductById(productId int64) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}
