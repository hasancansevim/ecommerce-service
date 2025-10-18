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
	GetAllProductsByStoreName(storeName string) []domain.Product
	GetProductById(productId int64) (domain.Product, error)
	AddProduct(product domain.Product) error
	DeleteProductById(productId int64) error
	UpdatePrice(productId int64, newPrice float32) error
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

func (productRepository *ProductRepository) GetAllProductsByStoreName(storeName string) []domain.Product {
	ctx := context.Background()
	products, err := productRepository.scannner.QueryAndScan(ctx, "SELECT * FROM products WHERE store = $1", storeName)
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
	err := productRepository.scannner.ExecuteExec(ctx, "insert into products (name,price,discount,store) values ($1,$2,$3,$4)", product.Name, product.Price, product.Discount, product.Store)
	if err != nil {
		return err
	}
	return nil
}

func (productRepository *ProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	ctx := context.Background()
	err := productRepository.scannner.ExecuteExec(ctx, "update products set price=$1 where id=$2", newPrice, productId)
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
