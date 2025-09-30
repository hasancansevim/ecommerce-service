package service

import (
	"errors"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service/model"
)

type IProductService interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStoreName(storeName string) []domain.Product
	GetProductById(productId int64) (domain.Product, error)
	AddProduct(productCreate model.ProductCreate) error
	DeleteProductById(productId int64) error
	UpdatePrice(productId int64, newPrice float32) error
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (productService *ProductService) GetAllProducts() []domain.Product {
	return productService.productRepository.GetAllProducts()
}

func (productService *ProductService) GetAllProductsByStoreName(storeName string) []domain.Product {
	return productService.productRepository.GetAllByStoreName(storeName)
}

func (productService *ProductService) GetProductById(productId int64) (domain.Product, error) {
	return productService.productRepository.GetProductById(productId)
}

func (productService *ProductService) AddProduct(productCreate model.ProductCreate) error {
	validateError := validateProductCreate(productCreate)
	if validateError != nil {
		return validateError
	}
	return productService.productRepository.AddProduct(domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	})
}

func (productService *ProductService) DeleteProductById(productId int64) error {
	return productService.productRepository.DeleteProductById(productId)
}

func (productService *ProductService) UpdatePrice(productId int64, newPrice float32) error {
	return productService.productRepository.UpdatePrice(productId, newPrice)
}

func validateProductCreate(productCreate model.ProductCreate) error {
	if productCreate.Name == "" {
		return errors.New("product name is required")
	}
	if productCreate.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}
	if productCreate.Discount < 0 {
		return errors.New("the discount rate must be at least 0")
	}
	return nil
}
