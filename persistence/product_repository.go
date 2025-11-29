package persistence

import (
	"context"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/helper"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetProductById(productId int64) (domain.Product, error)
	AddProduct(product domain.Product) error
	DeleteProductById(productId int64) error
	UpdatePrice(productId int64, newPrice float32) error
	UpdateProduct(productId uint, product domain.Product) error
	// GetProductsBy : Store,Slug,Featured,Category
}

type ProductRepository struct {
	dbPool   *pgxpool.Pool
	scannner *helper.GenericScanner[domain.Product]
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{
		dbPool:   dbPool,
		scannner: helper.NewGenericScanner(dbPool, helper.ScanProduct),
	}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	products, err := productRepository.scannner.QueryAndScan(ctx, "SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return products
}

func (productRepository *ProductRepository) GetProductById(productId int64) (domain.Product, error) {
	ctx := context.Background()
	product, err := productRepository.scannner.QueryRowAndScan(ctx, "SELECT * FROM products WHERE id = $1", productId)
	if err != nil {
		log.Fatal(err)
		return domain.Product{}, err
	}
	return product, nil
}

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()
	query := `
		INSERT INTO products 
		(name, slug, description, price, base_price, discount, image_url, meta_description, stock_quantity, is_active, is_featured, category_id, store_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	err := productRepository.scannner.ExecuteExec(ctx, query,
		product.Name,
		product.Slug,
		product.Description,
		product.Price,
		product.BasePrice,
		product.Discount,
		product.ImageUrl,
		product.MetaDescription,
		product.StockQuantity,
		product.IsActive,
		product.IsFeatured,
		product.CategoryId,
		product.StoreId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (productRepository *ProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	ctx := context.Background()
	query := "UPDATE products SET price = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2"
	err := productRepository.scannner.ExecuteExec(ctx, query, newPrice, productId)
	if err != nil {
		return err
	}
	return nil
}

func (productRepository *ProductRepository) UpdateProduct(productId uint, product domain.Product) error {
	ctx := context.Background()
	query := "UPDATE products set name=$1, slug=$2, description=$3, price=$4, base_price=$5, discount = $6, image_url=$7, meta_description=$8, stock_quantity=$9, is_active=$10, is_featured=$11, category_id=$12, store_id=$13 WHERE id = $14"
	err := productRepository.scannner.ExecuteExec(ctx, query,
		product.Name, product.Slug, product.Description, product.Price, product.BasePrice, product.Discount, product.ImageUrl, product.MetaDescription, product.StockQuantity, product.IsActive, product.IsFeatured, product.CategoryId, product.StoreId, product.Id)

	if err != nil {
		return err
	}
	return nil
}

func (productRepository *ProductRepository) DeleteProductById(productId int64) error {
	ctx := context.Background()
	err := productRepository.scannner.ExecuteExec(ctx, "delete from products where id = $1", productId)
	if err != nil {
		return err
	}
	return nil
}
