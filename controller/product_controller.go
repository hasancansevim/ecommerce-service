package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/pkg/validation"
	"go-ecommerce-service/service"
	"go-ecommerce-service/service/model"
	"strconv"

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
	e.PUT("/api/v1/products/:id", productController.UpdatePrice)
	e.DELETE("/api/v1/products/:id", productController.DeleteProduct)
}

func (productController *ProductController) GetAllProducts(c echo.Context) error {
	products := productController.productService.GetAllProducts()
	return productController.Success(c, products)
}

func (productController *ProductController) GetProductById(c echo.Context) error {
	productId, err := productController.ParseIdParam(c, "id")
	if err != nil {
		return productController.BadRequest(c, err)
	}

	productById, productByIdErr := productController.productService.GetProductById(productId)
	if productByIdErr != nil {
		return productController.BadRequest(c, productByIdErr)
	}
	return productController.Success(c, productById)
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

	if addProductErr := productController.productService.AddProduct(addProductRequest.ToModel()); addProductErr != nil {
		return productController.BadRequest(c, addProductErr)
	}

	return productController.Created(c)
}

func (productController *ProductController) UpdatePrice(c echo.Context) error {
	productID, err := productController.ParseIdParam(c, "id")
	if err != nil {
		return productController.BadRequest(c, err)
	}

	newPrice, err := strconv.ParseFloat(c.QueryParam("newPrice"), 32)
	if err != nil {
		return productController.BadRequest(c, err)
	}

	if err := productController.productService.UpdatePrice(productID, float32(newPrice)); err != nil {
		return productController.BadRequest(c, err)
	}

	return productController.Created(c)
}

func (productController *ProductController) DeleteProduct(c echo.Context) error {
	productID, err := productController.ParseIdParam(c, "id")
	if err != nil {
		return productController.BadRequest(c, err)
	}

	if err := productController.productService.DeleteProductById(productID); err != nil {
		return productController.BadRequest(c, err)
	}

	return productController.Created(c)
}
