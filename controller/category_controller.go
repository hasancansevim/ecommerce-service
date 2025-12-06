package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/service"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryService service.ICategoryService
	BaseController
}

func NewCategoryController(categoryService service.ICategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

func (categoryController *CategoryController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/categories", categoryController.GetAllCategories)
	e.GET("/api/v1/categories/:id", categoryController.GetCategoryById)
	e.GET("/api/v1/categories/:true", categoryController.GetCategoriesByIsActive)
	e.POST("/api/v1/categories", categoryController.AddCategory)
	e.PUT("/api/v1/categories/:id", categoryController.UpdateCategory)
	e.DELETE("/api/v1/categories/:id", categoryController.DeleteCategory)
}

func (categoryController *CategoryController) GetAllCategories(c echo.Context) error {
	categories := categoryController.categoryService.GetAllCategories()
	return categoryController.Success(c, categories, "")
}

func (categoryController *CategoryController) GetCategoryById(c echo.Context) error {
	id, parseIdErr := strconv.Atoi(c.Param("id"))
	if parseIdErr != nil {
		return parseIdErr
	}
	category, serviceErr := categoryController.categoryService.GetCategoryById(id)
	if serviceErr != nil {
		return serviceErr
	}
	return categoryController.Success(c, category, "")
}

func (categoryController *CategoryController) GetCategoriesByIsActive(c echo.Context) error {
	param := c.Param("true")
	b := parseBool(param)
	categories, serviceErr := categoryController.categoryService.GetCategoriesByIsActive(b)
	if serviceErr != nil {
		return serviceErr
	}
	return categoryController.Success(c, categories, "")
}

func (categoryController *CategoryController) AddCategory(c echo.Context) error {
	var addCategoryRequest request.AddCategoryRequest
	if bindErr := c.Bind(&addCategoryRequest); bindErr != nil {
		return bindErr
	}

	addedCategory, serviceErr := categoryController.categoryService.AddCategory(addCategoryRequest.ToModel())
	if serviceErr != nil {
		return serviceErr
	}
	return categoryController.Success(c, addedCategory, "")
}

func (categoryController *CategoryController) UpdateCategory(c echo.Context) error {
	id, parseIdErr := strconv.Atoi(c.Param("id"))
	if parseIdErr != nil {
		return parseIdErr
	}
	var updatedCategory request.AddCategoryRequest
	if bindErr := c.Bind(&updatedCategory); bindErr != nil {
		return bindErr
	}
	category, serviceErr := categoryController.categoryService.UpdateCategory(uint(id), updatedCategory.ToModel())
	if serviceErr != nil {
		return serviceErr
	}
	return categoryController.Success(c, category, "Kategori Güncellendi")
}

func (categoryController *CategoryController) DeleteCategory(c echo.Context) error {
	id, parseIdErr := strconv.Atoi(c.Param("id"))
	if parseIdErr != nil {
		return parseIdErr
	}
	serviceErr := categoryController.categoryService.DeleteCategory(uint(id))
	if serviceErr != nil {
		return serviceErr
	}
	return categoryController.Success(c, nil, "Kategori Silindi")
}

func parseBool(str string) bool {
	switch strings.ToLower(str) {
	case "true", "1", "yes", "on", "active":
		return true
	default:
		return false
	}
}
