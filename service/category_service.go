package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/dto"
	"go-ecommerce-service/persistence"
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
}

func NewCategoryService(categoryRepository persistence.ICategoryRepository) ICategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (categoryService *CategoryService) GetAllCategories() []dto.CategoryResponse {
	categories := categoryService.categoryRepository.GetAllCategories()
	categoriesDto := make([]dto.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		categoriesDto = append(categoriesDto, dto.CategoryResponse{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return categoriesDto
}
func (categoryService *CategoryService) GetCategoryById(id int) (dto.CategoryResponse, error) {
	category, err := categoryService.categoryRepository.GetCategoryById(id)
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	categoryDto := dto.CategoryResponse{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
	}
	return categoryDto, nil
}
func (categoryService *CategoryService) GetCategoriesByIsActive(isActive bool) ([]dto.CategoryResponse, error) {
	categories, err := categoryService.categoryRepository.GetCategoriesByIsActive(isActive)
	if err != nil {
		return []dto.CategoryResponse{}, err
	}
	categoriesDto := make([]dto.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		categoriesDto = append(categoriesDto, dto.CategoryResponse{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return categoriesDto, nil
}
func (categoryService *CategoryService) AddCategory(categoryCreate dto.CreateCategoryRequest) (dto.CategoryResponse, error) {
	addedCategory, err := categoryService.categoryRepository.AddCategory(domain.Category{
		Name:        categoryCreate.Name,
		Description: categoryCreate.Description,
		IsActive:    categoryCreate.IsActive,
	})
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	addedCategoryDto := dto.CategoryResponse{
		Id:          addedCategory.Id,
		Name:        addedCategory.Name,
		Description: addedCategory.Description,
	}
	return addedCategoryDto, nil

}
func (categoryService *CategoryService) UpdateCategory(categoryId uint, categoryCreate dto.CreateCategoryRequest) (dto.CategoryResponse, error) {
	updateCategory, err := categoryService.categoryRepository.UpdateCategory(categoryId, domain.Category{
		Name:        categoryCreate.Name,
		Description: categoryCreate.Description,
		IsActive:    categoryCreate.IsActive,
	})
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	updateCategoryDto := dto.CategoryResponse{
		Id:          updateCategory.Id,
		Name:        updateCategory.Name,
		Description: updateCategory.Description,
	}
	return updateCategoryDto, nil
}
func (categoryService *CategoryService) DeleteCategory(id uint) error {
	return categoryService.categoryRepository.DeleteCategory(id)
}
