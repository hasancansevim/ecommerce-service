package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service/model"
)

type ICategoryService interface {
	GetAllCategories() []domain.Category
	GetCategoryById(id int) (domain.Category, error)
	GetCategoriesByIsActive(isActive bool) ([]domain.Category, error)
	AddCategory(categoryCreate model.CategoryCreate) error
	UpdateCategory(categoryId uint, categoryCreate model.CategoryCreate) error
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

func (categoryService *CategoryService) GetAllCategories() []domain.Category {
	return categoryService.categoryRepository.GetAllCategories()
}
func (categoryService *CategoryService) GetCategoryById(id int) (domain.Category, error) {
	return categoryService.categoryRepository.GetCategoryById(id)
}
func (categoryService *CategoryService) GetCategoriesByIsActive(isActive bool) ([]domain.Category, error) {
	return categoryService.categoryRepository.GetCategoriesByIsActive(isActive)
}
func (categoryService *CategoryService) AddCategory(categoryCreate model.CategoryCreate) error {
	return categoryService.categoryRepository.AddCategory(domain.Category{
		Name:        categoryCreate.Name,
		Description: categoryCreate.Description,
		IsActive:    categoryCreate.IsActive,
	})
}
func (categoryService *CategoryService) UpdateCategory(categoryId uint, categoryCreate model.CategoryCreate) error {
	return categoryService.categoryRepository.UpdateCategory(categoryId, domain.Category{
		Name:        categoryCreate.Name,
		Description: categoryCreate.Description,
		IsActive:    categoryCreate.IsActive,
	})
}
func (categoryService *CategoryService) DeleteCategory(id uint) error {
	return categoryService.categoryRepository.DeleteCategory(id)
}
