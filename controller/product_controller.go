package controller

import (
	"fmt"
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/controller/response"
	"go-ecommerce-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products", productController.GetAllProducts)
	e.GET("/api/v1/products/:productId", productController.GetProductById)
	e.POST("/api/v1/products", productController.AddProduct)
	e.PUT("/api/v1/products/:id", productController.UpdatePrice)
	e.DELETE("/api/v1/products/:id", productController.DeleteProduct)
}

func (productController *ProductController) GetAllProducts(c echo.Context) error {
	queryParam := c.QueryParam("storeName")
	if queryParam != "" {
		allProductsByStoreName := productController.productService.GetAllProductsByStoreName(queryParam)
		return c.JSON(http.StatusOK, response.ToResponseList(allProductsByStoreName))
	} else {
		allProducts := productController.productService.GetAllProducts()
		return c.JSON(http.StatusOK, allProducts)
	}
}

func (productController *ProductController) GetProductById(c echo.Context) error {
	param := c.Param("productId")
	id, convertErr := strconv.Atoi(param)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}

	productById, productByIdErr := productController.productService.GetProductById(int64(id))
	if productByIdErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: productByIdErr.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ToResponse(productById))
}

func (productController *ProductController) AddProduct(c echo.Context) error {
	var addProductRequest request.AddProductRequest
	bindErr := c.Bind(&addProductRequest)

	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}

	addProductErr := productController.productService.AddProduct(addProductRequest.ToModel())

	if addProductErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: addProductErr.Error(),
		})
	}

	return c.NoContent(http.StatusCreated)
}

func (productController *ProductController) UpdatePrice(c echo.Context) error {
	param := c.Param("id")
	productId, convertErr := strconv.Atoi(param)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}

	queryParam := c.QueryParam("newPrice")
	fmt.Printf("Gelen newPrice: %s\n", queryParam)

	if queryParam == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: "newPrice parameter is required",
		})
	}
	newPrice, _ := strconv.ParseFloat(queryParam, 32)
	updatePriceErr := productController.productService.UpdatePrice(int64(productId), float32(newPrice))
	if updatePriceErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: updatePriceErr.Error(),
		})
	}
	return c.NoContent(http.StatusOK)
}

func (productController *ProductController) DeleteProduct(c echo.Context) error {
	param := c.Param("id")
	id, convertErr := strconv.Atoi(param)

	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}

	productDeleteErr := productController.productService.DeleteProductById(int64(id))
	if productDeleteErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: productDeleteErr.Error(),
		})
	}
	return c.NoContent(http.StatusOK)
}
