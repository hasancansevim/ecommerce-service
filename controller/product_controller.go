package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/pkg/validation"
	"go-ecommerce-service/service"
	"go-ecommerce-service/service/model"

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
		return productController.BadRequest(c, err)
	}

	product, productByIdErr := productController.productService.GetProductById(productId)
	if productByIdErr != nil {
		return productController.BadRequest(c, productByIdErr)
	}
	return productController.Success(c, product, "Ürün Getirildi")
}

func (productController *ProductController) AddProduct(c echo.Context) error {
	var addProductRequest request.AddProductRequest
	if bindErr := c.Bind(&addProductRequest); bindErr != nil {
		return productController.BadRequest(c, bindErr)
	}

	validator := validation.ProductCreateValidator{ProductReq: model.ProductCreate{
		Name:            addProductRequest.Name,
		Description:     addProductRequest.Description,
		BasePrice:       addProductRequest.BasePrice,
		StockQuantity:   addProductRequest.StockQuantity,
		ImageUrl:        addProductRequest.ImageUrl,
		CategoryId:      addProductRequest.CategoryId,
		IsActive:        addProductRequest.IsActive,
		IsFeatured:      addProductRequest.IsFeatured,
		MetaDescription: addProductRequest.MetaDescription,
		Slug:            addProductRequest.Slug,
		Price:           addProductRequest.Price,
		Discount:        addProductRequest.Discount,
		StoreId:         addProductRequest.StoreId,
	}}

	if validationErr := validator.Validate(); validationErr != nil {
		return productController.BadRequest(c, validationErr)
	}

	addedProduct, serviceErr := productController.productService.AddProduct(addProductRequest.ToModel())
	if serviceErr != nil {
		return productController.BadRequest(c, serviceErr)
	}

	return productController.Success(c, addedProduct, "Ürün Eklendi")
}

func (productController *ProductController) UpdateProduct(c echo.Context) error {
	productId, parseIdErr := productController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return productController.BadRequest(c, parseIdErr)
	}
	var updateProductRequest request.UpdateProductRequest
	if bindErr := c.Bind(&updateProductRequest); bindErr != nil {
		return productController.BadRequest(c, bindErr)
	}

	validator := validation.ProductCreateValidator{ProductReq: model.ProductCreate{
		Name:            updateProductRequest.Name,
		Description:     updateProductRequest.Description,
		BasePrice:       updateProductRequest.BasePrice,
		StockQuantity:   updateProductRequest.StockQuantity,
		ImageUrl:        updateProductRequest.ImageUrl,
		CategoryId:      updateProductRequest.CategoryId,
		IsActive:        updateProductRequest.IsActive,
		IsFeatured:      updateProductRequest.IsFeatured,
		MetaDescription: updateProductRequest.MetaDescription,
		Slug:            updateProductRequest.Slug,
		Price:           updateProductRequest.Price,
		Discount:        updateProductRequest.Discount,
		StoreId:         updateProductRequest.StoreId,
	}}

	if validationErr := validator.Validate(); validationErr != nil {
		return productController.BadRequest(c, validationErr)
	}

	updatedProduct, serviceErr := productController.productService.UpdateProduct(uint(productId), updateProductRequest.ToModel())
	if serviceErr != nil {
		return productController.BadRequest(c, serviceErr)
	}
	return productController.Success(c, updatedProduct, "Ürün Güncellendi")
}

func (productController *ProductController) DeleteProduct(c echo.Context) error {
	productId, err := productController.ParseIdParam(c, "id")
	if err != nil {
		return productController.BadRequest(c, err)
	}

	if err := productController.productService.DeleteProductById(productId); err != nil {
		return productController.BadRequest(c, err)
	}

	return productController.Created(c, nil, "Ürün Silindi")
}
