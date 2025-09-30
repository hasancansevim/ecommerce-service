package service

import (
	"errors"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
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

func (fakeRepository *FakeProductRepository) GetAllByStoreName(storeName string) []domain.Product {
	return []domain.Product{
		{
			Id:       1,
			Name:     "Laptop",
			Price:    20000.0,
			Discount: 10.0,
			Store:    storeName,
		},
		{
			Id:       2,
			Name:     "Klavye",
			Price:    800.0,
			Discount: 0.0,
			Store:    storeName,
		},
	}
}

func (fakeRepository *FakeProductRepository) GetProductById(productId int64) (domain.Product, error) {
	for i, product := range fakeRepository.products {
		if product.Id == productId {
			return fakeRepository.products[i], nil
		}
	}
	return domain.Product{}, errors.New("Product not found")
}

func (fakeRepository *FakeProductRepository) DeleteProductById(productId int64) error {
	for i, product := range fakeRepository.products {
		if product.Id == productId {
			fakeRepository.products = append(fakeRepository.products[:i], fakeRepository.products[i+1:]...)
		}
	}
	return nil
}

func (fakeRepository *FakeProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	for i, product := range fakeRepository.products {
		if product.Id == productId {
			fakeRepository.products[i].Price = newPrice
		}
	}
	return nil
}
