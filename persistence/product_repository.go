package persistence

import (
	"context"
	"errors"
	"fmt"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/common"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllByStoreName(storeName string) []domain.Product
	GetProductById(productId int64) (domain.Product, error)
	AddProduct(product domain.Product) error
	DeleteProductById(productId int64) error
	UpdateProduct(productId int64, newPrice float32) error
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{
		dbPool: dbPool,
	}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "select * from products")
	if err != nil {
		log.Error("Error while getting all products: %v", err)
		return []domain.Product{}
	}
	return extractProductsFromRows(productRows)
}

func (productRepository *ProductRepository) GetAllByStoreName(storeName string) []domain.Product {
	ctx := context.Background()
	getProductsByStoreNameSql := `select * from products where store = $1`
	productRowsByStoreName, productRowsErr := productRepository.dbPool.Query(ctx, getProductsByStoreNameSql, storeName)

	if productRowsErr != nil {
		log.Error("Error while getting all products: %v", productRowsErr)
		return []domain.Product{}
	}

	return extractProductsFromRows(productRowsByStoreName)
}

func (productRepository *ProductRepository) GetProductById(productId int64) (domain.Product, error) {
	ctx := context.Background()
	getByIdSql := `select * from products where id = $1`
	queryRow := productRepository.dbPool.QueryRow(ctx, getByIdSql, productId)

	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	scanError := queryRow.Scan(&id, &name, &price, &discount, &store)

	if scanError != nil && scanError.Error() == common.NOT_FOUND {
		return domain.Product{},
			errors.New(fmt.Sprintf("Product not found with id : %v", productId))
	}
	if scanError != nil {
		return domain.Product{}, errors.New(fmt.Sprintf("Error while getting product : %v", productId))
	}

	product := domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}

	return product, nil
}

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()

	insertProductSql := `insert into products (name,price,discount,store) values ($1,$2,$3,$4)`

	addNewProduct, addNewProductErr := productRepository.dbPool.Exec(ctx, insertProductSql,
		product.Name, product.Price, product.Discount, product.Store)

	if addNewProductErr != nil {
		log.Error("Failed to add new product %v :", addNewProductErr)
		return addNewProductErr
	}
	log.Info("Added new product %v : ", addNewProduct)
	return nil
}

func (productRepository *ProductRepository) UpdateProduct(productId int64, newPrice float32) error {
	ctx := context.Background()
	updatePriceSql := "update products set price=$1 where id=$2"
	_, execErr := productRepository.dbPool.Exec(ctx, updatePriceSql, newPrice, productId)

	if execErr != nil {
		return errors.New(fmt.Sprintf("Error while updateing product with id %v", productId))
	}
	log.Info("Updated product. id : %v , newPrice : %v", productId, newPrice)
	return nil
}

func (productRepository *ProductRepository) DeleteProductById(productId int64) error {
	ctx := context.Background()
	deleteProductSql := "delete from products where id = $1"
	_, execErr := productRepository.dbPool.Exec(ctx, deleteProductSql, productId)

	if execErr != nil {
		return errors.New(fmt.Sprintf("Error while deleting product with id %v", productId))
	}
	log.Info("Deleted product. id : %v ", productId)
	return nil
}

func extractProductsFromRows(productRows pgx.Rows) []domain.Product {
	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}
	return products
}
