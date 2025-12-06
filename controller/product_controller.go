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
	e.GET("/api/v1/products/:id", productController.GetProductById)
	e.POST("/api/v1/products", productController.AddProduct)
	e.PUT("/api/v1/products/:id", productController.UpdateProduct)
	e.DELETE("/api/v1/products/:id", productController.DeleteProduct)
}

func (productController *ProductController) GetAllProducts(c echo.Context) error {
	products := productController.productService.GetAllProducts()
	return productController.Success(c, products, "Tüm Ürünler Listelendi")
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

	return productController.Success(c, product, "Ürün Getirildi")
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

	return productController.Success(c, addedProduct, "Ürün Eklendi")
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
	return productController.Success(c, updatedProduct, "Ürün Güncellendi")
}

func (productController *ProductController) DeleteProduct(c echo.Context) error {
	productId, parseIdErr := productController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}

	if serviceErr := productController.productService.DeleteProductById(productId); serviceErr != nil {
		return serviceErr
	}

	return productController.Created(c, nil, "Ürün Silindi")
}
