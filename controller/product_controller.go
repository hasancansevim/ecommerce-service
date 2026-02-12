package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/service"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService service.IProductService
	BaseController
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products", productController.GetAllProducts)
	e.GET("/api/v1/products/search", productController.SearchProducts)
	e.GET("/api/v1/products/:id", productController.GetProductById)
	e.POST("/api/v1/products", productController.AddProduct)
	e.PUT("/api/v1/products/:id", productController.UpdateProduct)
	e.DELETE("/api/v1/products/:id", productController.DeleteProduct)
	e.POST("/api/v1/products/sync", productController.SyncElasticsearch)
}

func (productController *ProductController) GetAllProducts(c echo.Context) error {
	products := productController.productService.GetAllProducts()
	return productController.Success(c, products, "All products listed")
}

func (productController *ProductController) GetProductById(c echo.Context) error {
	productId, err := productController.ParseIdParam(c, "id")
	if err != nil {
		return err
	}

	product, productByIdErr := productController.productService.GetProductById(productId)
	if productByIdErr != nil {
		return productByIdErr
	}

	return productController.Success(c, product, "Product retrieved")
}

func (productController *ProductController) AddProduct(c echo.Context) error {
	var addProductRequest request.AddProductRequest
	if bindErr := c.Bind(&addProductRequest); bindErr != nil {
		return bindErr
	}

	addedProduct, serviceErr := productController.productService.AddProduct(addProductRequest.ToModel())
	if serviceErr != nil {
		return serviceErr
	}

	return productController.Success(c, addedProduct, "Product added")
}

func (productController *ProductController) UpdateProduct(c echo.Context) error {
	productId, parseIdErr := productController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}
	var updateProductRequest request.UpdateProductRequest
	if bindErr := c.Bind(&updateProductRequest); bindErr != nil {
		return bindErr
	}

	updatedProduct, serviceErr := productController.productService.UpdateProduct(uint(productId), updateProductRequest.ToModel())
	if serviceErr != nil {
		return serviceErr
	}
	return productController.Success(c, updatedProduct, "Product updated")
}

func (productController *ProductController) DeleteProduct(c echo.Context) error {
	productId, parseIdErr := productController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}

	if serviceErr := productController.productService.DeleteProductById(productId); serviceErr != nil {
		return serviceErr
	}

	return productController.Created(c, nil, "Product deleted")
}

// SearchProducts godoc
// @Summary      Search Products
// @Description  Smart search on products using Elasticsearch (Fuzzy & Wildcard).
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        q    query     string  true  "Aranacak Kelime (Örn: 'laptop')"
// @Success      200  {object}  dto.ProductResponse
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /products/search [get]
func (productController *ProductController) SearchProducts(c echo.Context) error {
	query := c.QueryParam("q")

	products, err := productController.productService.SearchProducts(query)
	if err != nil {
		return productController.BadRequest(c, err)
	}

	return productController.Success(c, products, "Search results")
}

func (productController *ProductController) SyncElasticsearch(ctx echo.Context) error {
	err := productController.productService.SyncElasticsearch()
	if err != nil {
		return productController.BadRequest(ctx, err)
	}
	return productController.Success(ctx, nil, "All products successfully synced to Elasticsearch! 🚀")
}
