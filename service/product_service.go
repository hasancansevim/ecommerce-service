package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service/model"
	"go-ecommerce-service/service/validation"
	"time"
)

type IProductService interface {
	GetAllProducts() []domain.Product
	GetProductById(productId int64) (domain.Product, error)
	AddProduct(productCreate model.ProductCreate) error
	DeleteProductById(productId int64) error
	UpdatePrice(productId int64, newPrice float32) error
	UpdateProduct(productId uint, product model.ProductCreate) error
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

func (productService *ProductService) GetProductById(productId int64) (domain.Product, error) {
	return productService.productRepository.GetProductById(productId)
}

func (productService *ProductService) AddProduct(productCreate model.ProductCreate) error {
	if validationError := validation.ValidateProductCreate(productCreate); validationError != nil {
		return validationError
	}
	return productService.productRepository.AddProduct(domain.Product{
		Name:            productCreate.Name,
		Slug:            productCreate.Slug,
		Description:     productCreate.Description,
		Price:           productCreate.Price,
		BasePrice:       productCreate.BasePrice,
		Discount:        productCreate.Discount,
		ImageUrl:        productCreate.ImageUrl,
		MetaDescription: productCreate.MetaDescription,
		StockQuantity:   productCreate.StockQuantity,
		IsActive:        productCreate.IsActive,
		IsFeatured:      productCreate.IsFeatured,
		CategoryId:      productCreate.CategoryId,
		StoreId:         productCreate.StoreId,
	})
}

func (productService *ProductService) DeleteProductById(productId int64) error {
	return productService.productRepository.DeleteProductById(productId)
}

func (productService *ProductService) UpdatePrice(productId int64, newPrice float32) error {
	return productService.productRepository.UpdatePrice(productId, newPrice)
}

func (productService *ProductService) UpdateProduct(productId uint, product model.ProductCreate) error {
	return productService.productRepository.UpdateProduct(productId, domain.Product{
		Name:            product.Name,
		Slug:            product.Slug,
		Description:     product.Description,
		Price:           product.Price,
		BasePrice:       product.BasePrice,
		Discount:        product.Discount,
		ImageUrl:        product.ImageUrl,
		MetaDescription: product.MetaDescription,
		StockQuantity:   product.StockQuantity,
		IsActive:        product.IsActive,
		IsFeatured:      product.IsFeatured,
		CategoryId:      product.CategoryId,
		StoreId:         product.StoreId,
		UpdatedAt:       time.Now(),
	})
}
