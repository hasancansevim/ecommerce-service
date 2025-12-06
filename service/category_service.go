package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/dto"
	"go-ecommerce-service/internal/rules"
	"go-ecommerce-service/persistence"
	_errors "go-ecommerce-service/pkg/errors"
)

type ICategoryService interface {
	GetAllCategories() []dto.CategoryResponse
	GetCategoryById(id int) (dto.CategoryResponse, error)
	GetCategoriesByIsActive(isActive bool) ([]dto.CategoryResponse, error)
	AddCategory(categoryCreate dto.CreateCategoryRequest) (dto.CategoryResponse, error)
	UpdateCategory(categoryId uint, categoryCreate dto.CreateCategoryRequest) (dto.CategoryResponse, error)
	DeleteCategory(id uint) error
}

type CategoryService struct {
	categoryRepository persistence.ICategoryRepository
	validator          *rules.CategoryRules
}

func NewCategoryService(categoryRepository persistence.ICategoryRepository) ICategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
		validator:          rules.NewCategoryRules(),
	}
}

func (categoryService *CategoryService) GetAllCategories() []dto.CategoryResponse {
	categories := categoryService.categoryRepository.GetAllCategories()
	return convertCategoriesResponse(categories)
}
func (categoryService *CategoryService) GetCategoryById(id int) (dto.CategoryResponse, error) {
	category, err := categoryService.categoryRepository.GetCategoryById(id)
	if err != nil {
		return dto.CategoryResponse{}, _errors.NewBadRequest(err.Error())
	}

	return convertToCategoryResponse(category), nil
}
func (categoryService *CategoryService) GetCategoriesByIsActive(isActive bool) ([]dto.CategoryResponse, error) {
	categories, err := categoryService.categoryRepository.GetCategoriesByIsActive(isActive)
	if err != nil {
		return []dto.CategoryResponse{}, _errors.NewBadRequest(err.Error())
	}

	return convertCategoriesResponse(categories), nil
}
func (categoryService *CategoryService) AddCategory(categoryCreate dto.CreateCategoryRequest) (dto.CategoryResponse, error) {
	if validationErr := categoryService.validator.ValidateStructure(categoryCreate); validationErr != nil {
		return dto.CategoryResponse{}, _errors.NewBadRequest(validationErr.Error())
	}
	addedCategory, err := categoryService.categoryRepository.AddCategory(domain.Category{
		Name:        categoryCreate.Name,
		Description: categoryCreate.Description,
		IsActive:    categoryCreate.IsActive,
	})
	if err != nil {
		return dto.CategoryResponse{}, _errors.NewBadRequest(err.Error())
	}

	return convertToCategoryResponse(addedCategory), nil

}
func (categoryService *CategoryService) UpdateCategory(categoryId uint, categoryCreate dto.CreateCategoryRequest) (dto.CategoryResponse, error) {
	if validationErr := categoryService.validator.ValidateStructure(categoryCreate); validationErr != nil {
		return dto.CategoryResponse{}, _errors.NewBadRequest(validationErr.Error())
	}

	updateCategory, err := categoryService.categoryRepository.UpdateCategory(categoryId, domain.Category{
		Name:        categoryCreate.Name,
		Description: categoryCreate.Description,
		IsActive:    categoryCreate.IsActive,
	})
	if err != nil {
		return dto.CategoryResponse{}, _errors.NewBadRequest(err.Error())
	}

	return convertToCategoryResponse(updateCategory), nil
}

func (categoryService *CategoryService) DeleteCategory(id uint) error {
	return categoryService.categoryRepository.DeleteCategory(id)
}

func convertToCategoryResponse(category domain.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
	}
}

func convertCategoriesResponse(categories []domain.Category) []dto.CategoryResponse {
	{
		categoriesDto := make([]dto.CategoryResponse, 0, len(categories))
		for _, category := range categories {
			categoriesDto = append(categoriesDto, convertToCategoryResponse(category))
		}
		return categoriesDto
	}
}
