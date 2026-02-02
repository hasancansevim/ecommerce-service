package persistence

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/helper"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetProductById(productId int64) (domain.Product, error)
	AddProduct(product domain.Product) (domain.Product, error)
	DeleteProductById(productId int64) error
	UpdateProduct(productId uint, product domain.Product) (domain.Product, error)
	// GetProductsBy : Store,Slug,Featured,Category
	SearchProducts(query string) ([]domain.Product, error)
	IndexProduct(product domain.Product) error
}

type ProductRepository struct {
	dbPool              *pgxpool.Pool
	scannner            *helper.GenericScanner[domain.Product]
	elasticSearchClient *elasticsearch.Client
}

func NewProductRepository(dbPool *pgxpool.Pool, elasticSearchClient *elasticsearch.Client) IProductRepository {
	return &ProductRepository{
		dbPool:              dbPool,
		scannner:            helper.NewGenericScanner(dbPool, helper.ScanProduct),
		elasticSearchClient: elasticSearchClient,
	}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	products, err := productRepository.scannner.QueryAndScan(ctx, "SELECT * FROM products")
	if err != nil {
		fmt.Println("SCANNING ERROR:", err)
		return nil
	}
	return products
}

func (productRepository *ProductRepository) GetProductById(productId int64) (domain.Product, error) {
	ctx := context.Background()
	product, err := productRepository.scannner.QueryRowAndScan(ctx, "SELECT * FROM products WHERE id = $1", productId)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (productRepository *ProductRepository) AddProduct(product domain.Product) (domain.Product, error) {
	ctx := context.Background()
	query := `
		INSERT INTO products 
		(name, slug, description, price, base_price, discount, image_url, meta_description, stock_quantity, is_active, is_featured, category_id, store_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING *
	`

	addedProduct, err := productRepository.scannner.QueryRowAndScan(ctx, query,
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
		product.StoreId)
	if err != nil {
		return domain.Product{}, err
	}

	productJSON, err := json.Marshal(addedProduct)
	if err != nil {
		return addedProduct, err
	}

	req := esapi.IndexRequest{
		Index:      "products", // Table Name
		DocumentID: strconv.Itoa(int(addedProduct.Id)),
		Body:       bytes.NewReader(productJSON),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, productRepository.elasticSearchClient)
	if err != nil {
		return domain.Product{}, err
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Printf("Elasticsearch Kayıt Hatası :%s\n", res.String())
	} else {
		fmt.Printf("Ürün Elasticsearch'e eklendi! ID : %d\n", addedProduct.Id)
	}

	return addedProduct, err
}

func (productRepository *ProductRepository) UpdateProduct(productId uint, product domain.Product) (domain.Product, error) {
	ctx := context.Background()
	query := `UPDATE products set name=$1, slug=$2, description=$3, price=$4, base_price=$5, discount = $6, image_url=$7, meta_description=$8, stock_quantity=$9, is_active=$10, is_featured=$11, category_id=$12, store_id=$13 WHERE id = $14 RETURNING *`
	updatedProduct, err := productRepository.scannner.QueryRowAndScan(ctx, query,
		product.Name, product.Slug, product.Description, product.Price, product.BasePrice, product.Discount, product.ImageUrl, product.MetaDescription, product.StockQuantity, product.IsActive, product.IsFeatured, product.CategoryId, product.StoreId, productId)

	if err != nil {
		return domain.Product{}, err
	}
	return updatedProduct, nil
}

func (productRepository *ProductRepository) DeleteProductById(productId int64) error {
	ctx := context.Background()
	err := productRepository.scannner.ExecuteExec(ctx, "delete from products where id = $1", productId)
	if err != nil {
		return err
	}
	return nil
}

func (productRepository *ProductRepository) SearchProducts(query string) ([]domain.Product, error) {
	ctx := context.Background()
	var buf bytes.Buffer

	queryString := strings.ToLower(query)

	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				// Yazım Hataları
				"should": []interface{}{
					map[string]interface{}{
						"multi_match": map[string]interface{}{
							"query":     query,
							"fields":    []string{"Name", "Description", "Slug"},
							"fuzziness": "AUTO",
						},
					},
					// Name içinde **query** ara
					map[string]interface{}{
						"wildcard": map[string]interface{}{
							"name": map[string]interface{}{
								"value": fmt.Sprintf("*%s*", queryString),
							},
						},
					},
					// Slug içinde **query** ara
					map[string]interface{}{
						"wildcard": map[string]interface{}{
							"slug": map[string]interface{}{
								"value": fmt.Sprintf("*%s*", queryString),
							},
						},
					},
				},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		return nil, err
	}

	res, err := productRepository.elasticSearchClient.Search(
		productRepository.elasticSearchClient.Search.WithContext(ctx),
		productRepository.elasticSearchClient.Search.WithIndex("products"),
		productRepository.elasticSearchClient.Search.WithBody(&buf),
		productRepository.elasticSearchClient.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("Elasticsearch cevap hatası :%s", res.String())
	}

	type ESResponse struct {
		Hits struct {
			Hits []struct {
				Source domain.Product `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	var response ESResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	fmt.Print(response.Hits.Hits)
	products := make([]domain.Product, 0)
	for _, hit := range response.Hits.Hits {
		products = append(products, hit.Source)
	}
	return products, nil
}

func (productRepository *ProductRepository) IndexProduct(product domain.Product) error {
	ctx := context.Background()

	productJSON, err := json.Marshal(product)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      "products",
		DocumentID: strconv.Itoa(int(product.Id)),
		Body:       bytes.NewReader(productJSON),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, productRepository.elasticSearchClient)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("Elasticsearch hatası: %s", res.String())
	}
	return nil
}
