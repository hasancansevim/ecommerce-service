package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/dto"
	"go-ecommerce-service/internal/rules"
	"go-ecommerce-service/persistence"
	_errors "go-ecommerce-service/pkg/errors"
	"go-ecommerce-service/pkg/util"
	"time"

	"github.com/redis/go-redis/v9"
)

type IProductService interface {
	GetAllProducts() []dto.ProductResponse
	GetProductById(productId int64) (dto.ProductResponse, error)
	AddProduct(productCreate dto.CreateProductRequest) (dto.ProductResponse, error)
	DeleteProductById(productId int64) error
	UpdateProduct(productId uint, product dto.CreateProductRequest) (dto.ProductResponse, error)
	SearchProducts(query string) ([]dto.ProductResponse, error)
	SyncElasticsearch() error
}

type ProductService struct {
	productRepository persistence.IProductRepository
	validator         *rules.ProductRules
	redisClient       *redis.Client
}

func NewProductService(productRepository persistence.IProductRepository, rdb *redis.Client) IProductService {
	return &ProductService{
		productRepository: productRepository,
		validator:         rules.NewProductRules(),
		redisClient:       rdb,
	}
}

func (productService *ProductService) GetAllProducts() []dto.ProductResponse {
	products := productService.productRepository.GetAllProducts()
	return convertToProductsResponse(products)
}

func (productService *ProductService) GetProductById(productId int64) (dto.ProductResponse, error) {
	ctx := context.Background()
	key := fmt.Sprintf("product:%d", productId)

	result, redisErr := productService.redisClient.Get(ctx, key).Result()
	if redisErr == nil {
		var cachedProduct dto.ProductResponse
		json.Unmarshal([]byte(result), &cachedProduct)
		return cachedProduct, nil
	}
	product, repositoryErr := productService.productRepository.GetProductById(productId)
	if repositoryErr != nil {
		return dto.ProductResponse{}, repositoryErr
	}
	response := convertToProductResponse(product)
	data, _ := json.Marshal(response)
	productService.redisClient.Set(ctx, key, data, 10*time.Minute)
	return response, nil
}

func (productService *ProductService) AddProduct(productCreate dto.CreateProductRequest) (dto.ProductResponse, error) {
	if validationErr := productService.validator.ValidateCreate(productCreate); validationErr != nil {
		return dto.ProductResponse{}, _errors.NewBadRequest(validationErr.Error())
	}

	addedProduct, repositoryErr := productService.productRepository.AddProduct(domain.Product{
		Name:            productCreate.Name,
		Slug:            util.GenerateUniqueSlug(productCreate.Name),
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
	if repositoryErr != nil {
		return dto.ProductResponse{}, _errors.NewInternalServerError(repositoryErr)
	}

	return convertToProductResponse(addedProduct), nil
}

func (productService *ProductService) DeleteProductById(productId int64) error {
	err := productService.productRepository.DeleteProductById(productId)
	if err != nil {
		return err
	}

	ctx := context.Background()

	key := fmt.Sprintf("product:%d", productId)
	productService.redisClient.Del(ctx, key)
	return nil
}

func (productService *ProductService) UpdateProduct(productId uint, product dto.CreateProductRequest) (dto.ProductResponse, error) {
	if validationErr := productService.validator.ValidateCreate(product); validationErr != nil {
		return dto.ProductResponse{}, _errors.NewBadRequest(validationErr.Error())
	}

	updatedProduct, repositoryErr := productService.productRepository.UpdateProduct(productId, domain.Product{
		Name:            product.Name,
		Slug:            util.GenerateUniqueSlug(product.Name),
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

	if repositoryErr != nil {
		return dto.ProductResponse{}, _errors.NewBadRequest(repositoryErr.Error())
	}

	ctx := context.Background()
	key := fmt.Sprintf("product:%d", productId)
	productService.redisClient.Del(ctx, key)
	return convertToProductResponse(updatedProduct), nil
}

func (productService *ProductService) SearchProducts(query string) ([]dto.ProductResponse, error) {
	products, err := productService.productRepository.SearchProducts(query)
	if err != nil {
		return nil, err
	}
	return convertToProductsResponse(products), nil
}

func (productService *ProductService) SyncElasticsearch() error {
	allProducts := productService.GetAllProducts()

	fmt.Printf("🔄 Senkronizasyon Başlıyor... Toplam Ürün: %d\n", len(allProducts))
	for _, p := range allProducts {
		err := productService.productRepository.IndexProduct(domain.Product{
			Id:              p.Id,
			Name:            p.Name,
			Slug:            p.Slug,
			Description:     p.Description,
			Price:           p.Price,
			BasePrice:       p.BasePrice,
			Discount:        p.Discount,
			ImageUrl:        p.ImageUrl,
			MetaDescription: p.MetaDescription,
			StockQuantity:   p.StockQuantity,
			IsActive:        p.IsActive,
			IsFeatured:      p.IsFeatured,
			CategoryId:      p.CategoryId,
			StoreId:         p.StoreId,
			UpdatedAt:       time.Now(),
		})
		if err != nil {
			fmt.Printf("❌ Hata (ID: %d): %v\n", p.Id, err)
			continue
		}
		fmt.Printf("✅ Indekslendi: %s\n", p.Name)
	}

	return nil
}

func convertToProductResponse(product domain.Product) dto.ProductResponse {
	return dto.ProductResponse{
		Id:              product.Id,
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
		CreatedAt:       product.CreatedAt,
		UpdatedAt:       product.UpdatedAt,
	}
}

func convertToProductsResponse(products []domain.Product) []dto.ProductResponse {
	productsDto := make([]dto.ProductResponse, 0, len(products))
	for _, product := range products {
		productsDto = append(productsDto, convertToProductResponse(product))
	}
	return productsDto
}
