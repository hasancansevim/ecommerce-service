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
	return categoryController.Success(c, categories)
}

func (categoryController *CategoryController) GetCategoryById(c echo.Context) error {
	id, convertErr := strconv.Atoi(c.Param("id"))
	if convertErr != nil {
		return categoryController.BadRequest(c, convertErr)
	}
	category, err := categoryController.categoryService.GetCategoryById(id)
	if err != nil {
		return categoryController.BadRequest(c, err)
	}
	return categoryController.Success(c, category)
}

func (categoryController *CategoryController) GetCategoriesByIsActive(c echo.Context) error {
	param := c.Param("true")
	b := parseBool(param)
	categories, err := categoryController.categoryService.GetCategoriesByIsActive(b)
	if err != nil {
		return categoryController.BadRequest(c, err)
	}
	return categoryController.Success(c, categories)
}

func (categoryController *CategoryController) AddCategory(c echo.Context) error {
	var addCategoryRequest request.AddCategoryRequest
	if bindErr := c.Bind(&addCategoryRequest); bindErr != nil {
		return categoryController.BadRequest(c, bindErr)
	}
	// validator
	err := categoryController.categoryService.AddCategory(addCategoryRequest.ToModel())
	if err != nil {
		return categoryController.BadRequest(c, err)
	}
	return categoryController.Success(c, nil)
}

func (categoryController *CategoryController) UpdateCategory(c echo.Context) error {
	id, convertErr := strconv.Atoi(c.Param("id"))
	if convertErr != nil {
		return categoryController.BadRequest(c, convertErr)
	}
	var updatedCategory request.AddCategoryRequest
	if bindErr := c.Bind(&updatedCategory); bindErr != nil {
		return categoryController.BadRequest(c, bindErr)
	}
	err := categoryController.categoryService.UpdateCategory(uint(id), updatedCategory.ToModel())
	if err != nil {
		return categoryController.BadRequest(c, err)
	}
	return categoryController.Success(c, nil)
}

func (categoryController *CategoryController) DeleteCategory(c echo.Context) error {
	id, convertErr := strconv.Atoi(c.Param("id"))
	if convertErr != nil {
		return categoryController.BadRequest(c, convertErr)
	}
	err := categoryController.categoryService.DeleteCategory(uint(id))
	if err != nil {
		return categoryController.BadRequest(c, err)
	}
	return categoryController.Success(c, nil)
}

func parseBool(str string) bool {
	switch strings.ToLower(str) {
	case "true", "1", "yes", "on", "active":
		return true
	default:
		return false
	}
}
